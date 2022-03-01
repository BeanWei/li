package lihcl

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/BeanWei/li/liadl/lihcl/spec"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

var (
	// Marshal returns the Atlas HCL encoding of v.
	Marshal = spec.MarshalerFunc(New().MarshalSpec)
)

var (
	// Unmarshal parses the Atlas HCL-encoded data and stores the result in the target.
	Unmarshal = spec.UnmarshalerFunc(New().UnmarshalSpec)
)

type (
	container struct {
		Body hcl.Body `hcl:",remain"`
	}

	// state implements spec.Unmarshaler and spec.Marshaler for Atlas HCL syntax
	// and stores a configuration for these operations.
	state struct {
		config *Config
	}
)

// MarshalSpec implements spec.Marshaler for Atlas HCL documents.
func (s *state) MarshalSpec(v interface{}) ([]byte, error) {
	r := &spec.Resource{}
	if err := r.Scan(v); err != nil {
		return nil, fmt.Errorf("lihcl: failed scanning %T to resource: %w", v, err)
	}
	return s.encode(r)
}

// UnmarshalSpec implements spec.Unmarshaler.
func (s *state) UnmarshalSpec(data []byte, v interface{}) error {
	ctx := s.config.newCtx()
	spec, err := s.decode(ctx, data)
	if err != nil {
		return fmt.Errorf("lihcl: failed decoding: %w", err)
	}
	if err := spec.As(v); err != nil {
		return fmt.Errorf("lihcl: failed reading spec as %T: %w", v, err)
	}
	return nil
}

// decode decodes the input Atlas HCL document and returns a *spec.Resource representing it.
func (s *state) decode(ctx *hcl.EvalContext, body []byte) (*spec.Resource, error) {
	parser := hclparse.NewParser()
	srcHCL, diag := parser.ParseHCL(body, "")
	if diag.HasErrors() {
		return nil, diag
	}
	if srcHCL == nil {
		return nil, fmt.Errorf("lihcl: no HCL syntax found in body")
	}
	c := &container{}
	ctx, err := evalCtx(ctx, srcHCL)
	if err != nil {
		return nil, err
	}
	if diag := gohcl.DecodeBody(srcHCL.Body, ctx, c); diag.HasErrors() {
		return nil, diag
	}
	return s.extract(ctx, c.Body)
}

func (s *state) extract(ctx *hcl.EvalContext, remain hcl.Body) (*spec.Resource, error) {
	body, ok := remain.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("lihcl: expected remainder to be of type *hclsyntax.Body")
	}
	attrs, err := s.toAttrs(ctx, body.Attributes, nil)
	if err != nil {
		return nil, err
	}
	res := &spec.Resource{
		Attrs: attrs,
	}
	for _, blk := range body.Blocks {
		ctx, err := setBlockVars(ctx.NewChild(), blk.Body)
		if err != nil {
			return nil, err
		}
		resource, err := s.toResource(ctx, blk, []string{blk.Type})
		if err != nil {
			return nil, err
		}
		res.Children = append(res.Children, resource)
	}
	return res, nil
}

// mayExtendVars gets the current scope context, and extend it with additional
// variables if it was configured this way using WithScopedEnums.
func (s *state) mayExtendVars(ctx *hcl.EvalContext, scope []string) *hcl.EvalContext {
	vars, ok := s.config.pathVars[strings.Join(scope, ".")]
	if !ok {
		return ctx
	}
	ctx = ctx.NewChild()
	ctx.Variables = vars
	return ctx
}

func (s *state) toAttrs(ctx *hcl.EvalContext, hclAttrs hclsyntax.Attributes, scope []string) ([]*spec.Attr, error) {
	var attrs []*spec.Attr
	for _, hclAttr := range hclAttrs {
		ctx := s.mayExtendVars(ctx, append(scope, hclAttr.Name))
		at := &spec.Attr{K: hclAttr.Name}
		value, diag := hclAttr.Expr.Value(ctx)
		if diag.HasErrors() {
			return nil, diag
		}
		var err error
		switch {
		case isRef(value):
			at.V = &spec.Ref{V: value.GetAttr("__ref").AsString()}
		case value.Type() == ctyRawExpr:
			at.V = value.EncapsulatedValue().(*spec.RawExpr)
		case value.Type() == ctyTypeSpec:
			at.V = value.EncapsulatedValue().(*spec.Type)
		case value.Type().IsTupleType():
			at.V, err = extractListValue(value)
		default:
			at.V, err = extractLiteralValue(value)
		}
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, at)
	}
	// hclsyntax.Attrs is an alias for map[string]*Attribute
	sort.Slice(attrs, func(i, j int) bool {
		return attrs[i].K < attrs[j].K
	})
	return attrs, nil
}

func isRef(v cty.Value) bool {
	return v.Type().IsObjectType() && v.Type().HasAttribute("__ref")
}

func extractListValue(value cty.Value) (*spec.ListValue, error) {
	lst := &spec.ListValue{}
	it := value.ElementIterator()
	for it.Next() {
		_, v := it.Element()
		if isRef(v) {
			lst.V = append(lst.V, &spec.Ref{V: v.GetAttr("__ref").AsString()})
			continue
		}
		litv, err := extractLiteralValue(v)
		if err != nil {
			return nil, err
		}
		lst.V = append(lst.V, litv)
	}
	return lst, nil
}

func extractLiteralValue(value cty.Value) (*spec.LiteralValue, error) {
	switch value.Type() {
	case ctySchemaLit:
		return value.EncapsulatedValue().(*spec.LiteralValue), nil
	case cty.String:
		return &spec.LiteralValue{V: strconv.Quote(value.AsString())}, nil
	case cty.Number:
		bf := value.AsBigFloat()
		num, _ := bf.Float64()
		return &spec.LiteralValue{V: strconv.FormatFloat(num, 'f', -1, 64)}, nil
	case cty.Bool:
		return &spec.LiteralValue{V: strconv.FormatBool(value.True())}, nil
	default:
		return nil, fmt.Errorf("lihcl: unsupported type %q", value.Type().GoString())
	}
}

func (s *state) toResource(ctx *hcl.EvalContext, block *hclsyntax.Block, scope []string) (*spec.Resource, error) {
	spec := &spec.Resource{
		Type: block.Type,
	}
	if len(block.Labels) > 0 {
		spec.Name = block.Labels[0]
	}
	ctx = s.mayExtendVars(ctx, scope)
	attrs, err := s.toAttrs(ctx, block.Body.Attributes, scope)
	if err != nil {
		return nil, err
	}
	spec.Attrs = attrs
	for _, blk := range block.Body.Blocks {
		res, err := s.toResource(ctx, blk, append(scope, blk.Type))
		if err != nil {
			return nil, err
		}
		spec.Children = append(spec.Children, res)
	}
	return spec, nil
}

// encode encodes the give *spec.Resource into a byte slice containing an Atlas HCL
// document representing it.
func (s *state) encode(r *spec.Resource) ([]byte, error) {
	f := hclwrite.NewFile()
	body := f.Body()
	// If the resource has a Type then it is rendered as an HCL block.
	if r.Type != "" {
		blk := body.AppendNewBlock(r.Type, labels(r))
		body = blk.Body()
	}
	for _, attr := range r.Attrs {
		if err := s.writeAttr(attr, body); err != nil {
			return nil, err
		}
	}
	for _, res := range r.Children {
		if err := s.writeResource(res, body); err != nil {
			return nil, err
		}
	}
	var buf bytes.Buffer
	_, err := f.WriteTo(&buf)
	return buf.Bytes(), err
}

func (s *state) writeResource(b *spec.Resource, body *hclwrite.Body) error {
	blk := body.AppendNewBlock(b.Type, labels(b))
	nb := blk.Body()
	for _, attr := range b.Attrs {
		if err := s.writeAttr(attr, nb); err != nil {
			return err
		}
	}
	for _, b := range b.Children {
		if err := s.writeResource(b, nb); err != nil {
			return err
		}
	}
	return nil
}

func labels(r *spec.Resource) []string {
	var l []string
	if r.Name != "" {
		l = append(l, r.Name)
	}
	return l
}

func (s *state) writeAttr(attr *spec.Attr, body *hclwrite.Body) error {
	attr = normalizeLiterals(attr)
	switch v := attr.V.(type) {
	case *spec.Ref:
		expr := strings.ReplaceAll(v.V, "$", "")
		body.SetAttributeRaw(attr.K, hclRawTokens(expr))
	case *spec.Type:
		if v.IsRef {
			expr := strings.ReplaceAll(v.T, "$", "")
			body.SetAttributeRaw(attr.K, hclRawTokens(expr))
			break
		}
		spec, ok := s.findTypeSpec(v.T)
		if !ok {
			v := fmt.Sprintf("sql(%q)", v.T)
			body.SetAttributeRaw(attr.K, hclRawTokens(v))
			break
		}
		st, err := hclType(spec, v)
		if err != nil {
			return err
		}
		body.SetAttributeRaw(attr.K, hclRawTokens(st))
	case *spec.LiteralValue:
		body.SetAttributeRaw(attr.K, hclRawTokens(v.V))
	case *spec.RawExpr:
		// TODO(rotemtam): the func name should be decided on contextual basis.
		fnc := fmt.Sprintf("sql(%q)", v.X)
		body.SetAttributeRaw(attr.K, hclRawTokens(fnc))
	case *spec.ListValue:
		// Skip scanning nil slices ([]T(nil)) by default. Users that
		// want to print empty lists, should use make([]T, 0) instead.
		if v.V == nil {
			return nil
		}
		lst := make([]string, 0, len(v.V))
		for _, item := range v.V {
			switch v := item.(type) {
			case *spec.Ref:
				expr := strings.ReplaceAll(v.V, "$", "")
				lst = append(lst, expr)
			case *spec.LiteralValue:
				lst = append(lst, v.V)
			default:
				return fmt.Errorf("cannot write elem type %T of attr %q to HCL list", v, attr)
			}
		}
		body.SetAttributeRaw(attr.K, hclRawList(lst))
	default:
		return fmt.Errorf("schemacl: unknown literal type %T", v)
	}
	return nil
}

// normalizeLiterals transforms attriburtes with LiteralValue that cannot be
// written as correct HCL into RawExpr.
func normalizeLiterals(attr *spec.Attr) *spec.Attr {
	lv, ok := attr.V.(*spec.LiteralValue)
	if !ok {
		return attr
	}
	exp := "x = " + lv.V
	p := hclparse.NewParser()
	if _, diag := p.ParseHCL([]byte(exp), ""); diag != nil {
		return &spec.Attr{K: attr.K, V: &spec.RawExpr{X: lv.V}}
	}
	return attr
}

func (s *state) findTypeSpec(t string) (*spec.TypeSpec, bool) {
	for _, v := range s.config.types {
		if v.T == t {
			return v, true
		}
	}
	return nil, false
}

func hclType(s *spec.TypeSpec, typ *spec.Type) (string, error) {
	if len(typeFuncArgs(s)) == 0 {
		return s.Name, nil
	}
	args := make([]string, 0, len(s.Attributes))
	for _, param := range typeFuncArgs(s) {
		arg, ok := findAttr(typ.Attrs, param.Name)
		if !ok {
			continue
		}
		switch val := arg.V.(type) {
		case *spec.LiteralValue:
			args = append(args, val.V)
		case *spec.ListValue:
			for _, li := range val.V {
				lit, ok := li.(*spec.LiteralValue)
				if !ok {
					return "", errors.New("expecting literal value")
				}
				args = append(args, lit.V)
			}
		}
	}
	// If no args were chosen and the type can be described without a function.
	if len(args) == 0 && len(typeFuncReqArgs(s)) == 0 {
		return s.Name, nil
	}
	return fmt.Sprintf("%s(%s)", s.Name, strings.Join(args, ",")), nil
}

func findAttr(attrs []*spec.Attr, k string) (*spec.Attr, bool) {
	for _, attr := range attrs {
		if attr.K == k {
			return attr, true
		}
	}
	return nil, false
}

func hclRawTokens(s string) hclwrite.Tokens {
	return hclwrite.Tokens{
		&hclwrite.Token{
			Type:  hclsyntax.TokenIdent,
			Bytes: []byte(s),
		},
	}
}

func hclRawList(items []string) hclwrite.Tokens {
	t := hclwrite.Tokens{&hclwrite.Token{
		Type:  hclsyntax.TokenOBrack,
		Bytes: []byte("["),
	}}
	for i, item := range items {
		if i > 0 {
			t = append(t, &hclwrite.Token{Type: hclsyntax.TokenComma, Bytes: []byte(",")})
		}
		t = append(t, &hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(item)})
	}
	t = append(t, &hclwrite.Token{
		Type:  hclsyntax.TokenCBrack,
		Bytes: []byte("]"),
	})
	return t
}
