package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func RecordPickerRecordFormModal(name string) *recordpickerrecordformmodalBuilder {
	return &recordpickerrecordformmodalBuilder{
		&NodeBuilder{
			schema: &ui.Schema{
				Name:            name,
				Type:            ui.SchemaTypeVoid,
				XComponent:      ui.ComponentRecordPickerRecordFormModal,
				XComponentProps: make(map[string]interface{}),
				HandlerNames:    make([]string, 0),
			},
		},
	}
}

type recordpickerrecordformmodalBuilder struct {
	*NodeBuilder
}

func (b *recordpickerrecordformmodalBuilder) AC(f ac.AC) *recordpickerrecordformmodalBuilder {
	b.schema.AC = f
	return b
}

func (b *recordpickerrecordformmodalBuilder) ForInit(operation string, handler interface{}) *recordpickerrecordformmodalBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}

func (b *recordpickerrecordformmodalBuilder) Body(elements ...view.Node) *recordpickerrecordformmodalBuilder {
	b.Items(elements...)
	return b
}

func (b *recordpickerrecordformmodalBuilder) Footer(elements ...view.Node) *recordpickerrecordformmodalBuilder {
	b.Children(elements...)
	return b
}
