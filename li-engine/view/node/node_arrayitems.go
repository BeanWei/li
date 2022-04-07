package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ArrayItems(name string) *arrayitemsBuilder {
	return &arrayitemsBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentArrayItems,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
			Properties:      gmap.NewListMap(),
		},
	}}
}

type arrayitemsBuilder struct {
	*NodeBuilder
}

func (b *arrayitemsBuilder) AC(f ac.AC) *arrayitemsBuilder {
	b.schema.AC = f
	return b
}

// Rows 一维数组
func (b *arrayitemsBuilder) Rows(elements ...view.Node) *arrayitemsBuilder {
	if b.schema.Items == nil {
		b.schema.Items = &ui.Schema{
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentSpace,
			Properties: gmap.NewListMap(),
		}
	}
	for _, element := range elements {
		es := element.Schema()
		b.schema.Items.Properties.Set(es.Name, es)
	}
	return b
}

// Columns 对象数组
func (b *arrayitemsBuilder) Columns(elements ...view.Node) *arrayitemsBuilder {
	b.Items(elements...)
	return b
}
