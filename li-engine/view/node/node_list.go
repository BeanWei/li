package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func List(name string) *listBuilder {
	return &listBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentList,
			XComponentProps: map[string]interface{}{
				"selection": make(map[string]interface{}),
			},
			Properties:   gmap.NewListMap(),
			HandlerNames: make([]string, 0),
		},
	}}
}

type listBuilder struct {
	*NodeBuilder
}

func (b *listBuilder) AC(f ac.AC) *listBuilder {
	b.schema.AC = f
	return b
}

func (b *listBuilder) Title(title string) *listBuilder {
	b.schema.Title = title
	return b
}

func (b *listBuilder) Description(description string) *listBuilder {
	b.schema.Description = description
	return b
}

func (b *listBuilder) DecoratorCard() *listBuilder {
	b.schema.XDecorator = ui.DecoratorCardItem
	return b
}

func (b *listBuilder) ForInit(operation string, handler interface{}) *listBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listBuilder) EnableFilter() *listBuilder {
	b.schema.XComponentProps["filter"] = true
	return b
}

func (b *listBuilder) EnableLightFilter() *listBuilder {
	b.schema.XComponentProps["filter"] = "light"
	return b
}

func (b *listBuilder) SelectionMultiple(multiple bool) *listBuilder {
	sel, ok := b.schema.XComponentProps["selection"].(map[string]interface{})
	if ok {
		sel["multiple"] = multiple
		b.schema.XComponentProps["selection"] = sel
	}
	return b
}

func (b *listBuilder) SelectionColumnTitle(title string) *listBuilder {
	sel, ok := b.schema.XComponentProps["selection"].(map[string]interface{})
	if ok {
		sel["columnTitle"] = title
		b.schema.XComponentProps["selection"] = sel
	}
	return b
}

func (b *listBuilder) SelectionColumnWidth(width int) *listBuilder {
	sel, ok := b.schema.XComponentProps["selection"].(map[string]interface{})
	if ok {
		sel["columnWidth"] = width
		b.schema.XComponentProps["selection"] = sel
	}
	return b
}

func (b *listBuilder) SelectionFixed(fixed string) *listBuilder {
	sel, ok := b.schema.XComponentProps["selection"].(map[string]interface{})
	if ok {
		sel["fixed"] = fixed
		b.schema.XComponentProps["selection"] = sel
	}
	return b
}
