package spec

import "strconv"

// LitAttr is a helper method for constructing *Attr instances that contain literal values.
func LitAttr(k, v string) *Attr {
	return &Attr{
		K: k,
		V: &LiteralValue{V: v},
	}
}

// StrLitAttr is a helper method for constructing *Attr instances that contain literal values
// representing string literals.
func StrLitAttr(k, v string) *Attr {
	return LitAttr(k, strconv.Quote(v))
}

// ListAttr is a helper method for constructing *Attr instances that contain list values.
func ListAttr(k string, values ...string) *Attr {
	lst := &ListValue{}
	for _, v := range values {
		lst.V = append(lst.V, &LiteralValue{V: strconv.Quote(v)})
	}
	return &Attr{
		K: k,
		V: lst,
	}
}
