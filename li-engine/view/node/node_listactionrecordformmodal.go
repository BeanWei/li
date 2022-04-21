package node

import (
	"strings"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListActionRecordFormModal(name string) *listactionrecordformmodalBuilder {
	return &listactionrecordformmodalBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeVoid,
			XComponent: ui.ComponentListActionRecordFormModal,
			XComponentProps: map[string]interface{}{
				"modalProps":  make(map[string]interface{}),
				"layoutProps": make(map[string]interface{}),
			},
			HandlerNames: make([]string, 0),
		},
	}}
}

type listactionrecordformmodalBuilder struct {
	*NodeBuilder
}

func (b *listactionrecordformmodalBuilder) AC(f ac.AC) *listactionrecordformmodalBuilder {
	b.schema.AC = f
	return b
}

func (b *listactionrecordformmodalBuilder) Title(title string) *listactionrecordformmodalBuilder {
	b.SetTitle(title)
	return b
}

func (b *listactionrecordformmodalBuilder) Description(description string) *listactionrecordformmodalBuilder {
	b.SetDescription(description)
	return b
}

func (b *listactionrecordformmodalBuilder) ForInit(operation string, handler interface{}) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonStyle(style map[string]interface{}) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["style"] = style
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonType(typ string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["type"] = typ
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonStatus(status string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["status"] = status
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonSize(size string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonShape(shape string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["shape"] = shape
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonIcon(icon string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["icon"] = icon
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonIconOnly(iconOnly bool) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["iconOnly"] = iconOnly
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonLong(long bool) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["long"] = long
	return b
}

func (b *listactionrecordformmodalBuilder) ButtonPosition(position string) *listactionrecordformmodalBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *listactionrecordformmodalBuilder) ModalTitle(title string) *listactionrecordformmodalBuilder {
	modalProps, ok := b.schema.XComponentProps["modalProps"].(map[string]interface{})
	if ok {
		if !strings.HasPrefix(title, "{{t('") {
			title = "{{t('" + title + "')}}"
		}
		modalProps["title"] = title
		b.schema.XComponentProps["modalProps"] = modalProps
	}
	return b
}

func (b *listactionrecordformmodalBuilder) Body(elements ...view.Node) *listactionrecordformmodalBuilder {
	b.Items(elements...)
	return b
}

func (b *listactionrecordformmodalBuilder) Footer(elements ...view.Node) *listactionrecordformmodalBuilder {
	b.Children(elements...)
	return b
}
