package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Node(name string) *NodeBuilder {
	return &NodeBuilder{schema: &ui.Schema{
		Name: name,
	}}
}

type NodeBuilder struct {
	schema *ui.Schema
}

func (b *NodeBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *NodeBuilder) SetSchema(schema *ui.Schema) *NodeBuilder {
	name := b.schema.Name
	b.schema = schema
	b.schema.Name = name
	return b
}

func (b *NodeBuilder) SetName(name string) *NodeBuilder {
	b.schema.Name = name
	return b
}

func (b *NodeBuilder) SetType(typ string) *NodeBuilder {
	b.schema.Type = typ
	return b
}

func (b *NodeBuilder) SetTitle(title string) *NodeBuilder {
	b.schema.Title = title
	return b
}

func (b *NodeBuilder) SetDescription(description string) *NodeBuilder {
	b.schema.Description = description
	return b
}

func (b *NodeBuilder) SetDefault(value interface{}) *NodeBuilder {
	b.schema.Default = value
	return b
}

func (b *NodeBuilder) SetReadOnly(readOnly bool) *NodeBuilder {
	b.schema.ReadOnly = readOnly
	return b
}

func (b *NodeBuilder) SetWriteOnly(writeOnly bool) *NodeBuilder {
	b.schema.WriteOnly = writeOnly
	return b
}

func (b *NodeBuilder) SetEnum(enum []map[string]interface{}) *NodeBuilder {
	b.schema.Enum = enum
	return b
}

func (b *NodeBuilder) SetConst(cst interface{}) *NodeBuilder {
	b.schema.Const = cst
	return b
}

func (b *NodeBuilder) SetMultipleOf(multipleOf int) *NodeBuilder {
	b.schema.MultipleOf = multipleOf
	return b
}

func (b *NodeBuilder) SetMaximum(maximum int) *NodeBuilder {
	b.schema.Maximum = maximum
	return b
}

func (b *NodeBuilder) SetExclusiveMaximum(exclusiveMaximum int) *NodeBuilder {
	b.schema.ExclusiveMaximum = exclusiveMaximum
	return b
}

func (b *NodeBuilder) SetMinimum(minimum int) *NodeBuilder {
	b.schema.Minimum = minimum
	return b
}

func (b *NodeBuilder) SetExclusiveMinimum(exclusiveMinimum int) *NodeBuilder {
	b.schema.ExclusiveMinimum = exclusiveMinimum
	return b
}

func (b *NodeBuilder) SetMaxLength(maxLength int) *NodeBuilder {
	b.schema.MaxLength = maxLength
	return b
}

func (b *NodeBuilder) SetMinLength(minLength int) *NodeBuilder {
	b.schema.MinLength = minLength
	return b
}

func (b *NodeBuilder) SetPattern(pattern string) *NodeBuilder {
	b.schema.Pattern = pattern
	return b
}

func (b *NodeBuilder) SetMaxItems(maxItems int) *NodeBuilder {
	b.schema.MaxItems = maxItems
	return b
}

func (b *NodeBuilder) SetMinItems(minItems int) *NodeBuilder {
	b.schema.MinItems = minItems
	return b
}

func (b *NodeBuilder) SetUniqueItems(uniqueItems bool) *NodeBuilder {
	b.schema.UniqueItems = uniqueItems
	return b
}

func (b *NodeBuilder) SetMaxProperties(maxProperties int) *NodeBuilder {
	b.schema.MaxProperties = maxProperties
	return b
}

func (b *NodeBuilder) SetMinProperties(minProperties int) *NodeBuilder {
	b.schema.MinProperties = minProperties
	return b
}

func (b *NodeBuilder) SetRequired(required bool) *NodeBuilder {
	b.schema.Required = required
	return b
}

func (b *NodeBuilder) SetFormat(format string) *NodeBuilder {
	b.schema.Format = format
	return b
}

func (b *NodeBuilder) SetXIndex(xIndex int) *NodeBuilder {
	b.schema.XIndex = xIndex
	return b
}

func (b *NodeBuilder) SetXPattern(xPattern string) *NodeBuilder {
	b.schema.XPattern = xPattern
	return b
}

func (b *NodeBuilder) SetXDisplay(xDisplay string) *NodeBuilder {
	b.schema.XDisplay = xDisplay
	return b
}

func (b *NodeBuilder) SetXValidator(xValidator string) *NodeBuilder {
	b.schema.XValidator = xValidator
	return b
}

func (b *NodeBuilder) SetXDecorator(xDecorator string) *NodeBuilder {
	b.schema.XDecorator = xDecorator
	return b
}

func (b *NodeBuilder) SetXDecoratorProps(xDecoratorProps map[string]interface{}) *NodeBuilder {
	b.schema.XDecoratorProps = xDecoratorProps
	return b
}

func (b *NodeBuilder) SetXComponent(xComponent string) *NodeBuilder {
	b.schema.XComponent = xComponent
	return b
}

func (b *NodeBuilder) SetXComponentProps(xComponentProps map[string]interface{}) *NodeBuilder {
	b.schema.XComponentProps = xComponentProps
	return b
}

func (b *NodeBuilder) SetXReactions(xReactions map[string]interface{}) *NodeBuilder {
	b.schema.XReactions = xReactions
	return b
}

func (b *NodeBuilder) SetXContent(xContent string) *NodeBuilder {
	b.schema.XContent = xContent
	return b
}

func (b *NodeBuilder) SetXVisible(xVisible bool) *NodeBuilder {
	b.schema.XVisible = xVisible
	return b
}

func (b *NodeBuilder) SetXHidden(xHidden bool) *NodeBuilder {
	b.schema.XHidden = xHidden
	return b
}

func (b *NodeBuilder) SetXDisabled(xDisabled bool) *NodeBuilder {
	b.schema.XDisabled = xDisabled
	return b
}

func (b *NodeBuilder) SetXEditable(xEditable bool) *NodeBuilder {
	b.schema.XEditable = xEditable
	return b
}

func (b *NodeBuilder) SetXReadOnly(xReadOnly bool) *NodeBuilder {
	b.schema.XReadOnly = xReadOnly
	return b
}

func (b *NodeBuilder) SetXReadPretty(xreadpretty bool) *NodeBuilder {
	b.schema.XReadPretty = xreadpretty
	return b
}

func (b *NodeBuilder) SetXData(xData map[string]interface{}) *NodeBuilder {
	b.schema.XData = xData
	return b
}

func (b *NodeBuilder) SetProperties(properties map[string]*ui.Schema) *NodeBuilder {
	b.schema.Properties = properties
	return b
}
