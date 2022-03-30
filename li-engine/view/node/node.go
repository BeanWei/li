package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/text/gstr"
)

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

func (b *NodeBuilder) SetProperties(properties *gmap.ListMap) *NodeBuilder {
	b.schema.Properties = properties
	return b
}

func (b *NodeBuilder) Items(elements ...view.Node) *NodeBuilder {
	b.children(true, elements...)
	return b
}

func (b *NodeBuilder) Children(elements ...view.Node) *NodeBuilder {
	b.children(false, elements...)
	return b
}

func (b *NodeBuilder) children(isItems bool, elements ...view.Node) *NodeBuilder {
	if isItems {
		if b.schema.Items == nil {
			b.schema.Items = &ui.Schema{
				Type:       ui.SchemaTypeObject,
				Properties: gmap.NewListMap(),
			}
		}
	} else {
		if b.schema.Properties == nil {
			b.schema.Properties = gmap.NewListMap()
		}
	}
	for _, element := range elements {
		es := element.Schema()
		if es.XPath == "" {
			es.XPath = es.Name
		}
		if isItems {
			es.XPath = b.schema.XPath + ".items.properties." + es.XPath
		} else {
			es.XPath = b.schema.XPath + ".properties." + es.XPath
		}
		if es.AC != nil {
			ac.Bind(es.XPath, es.AC)
			for _, hn := range es.HandlerNames {
				controller.UseWithSchemaPath(hn, es.XPath)
			}
		} else if len(es.HandlerNames) > 0 {
			// 如果当前节点上定义了 controller 但是没有定义 AC.
			// 则向上查找离的最近的父节点上的 AC.
			pathItems := gstr.Split(b.schema.XPath, ".")
			pathItemsLen := len(pathItems)
			parentPaths := make([]string, pathItemsLen)
			for i := 0; i < pathItemsLen; i++ {
				// 把路径最短(离的最远)的放在最后面
				parentPaths[pathItemsLen-1-i] = gstr.Join(pathItems[0:i+1], ".")
			}
			for _, path := range parentPaths {
				f := ac.Get(path)
				if f != nil {
					for _, hn := range es.HandlerNames {
						controller.UseWithSchemaPath(hn, path)
					}
					break
				}
			}
		}
		if isItems {
			b.schema.Items.Properties.Set(es.Name, es)
		} else {
			b.schema.Properties.Set(es.Name, es)
		}
	}
	return b
}
