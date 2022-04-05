package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/controller"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func RecordPickerRecordFormDrawer(name string) *recordpickerrecordformdrawerBuilder {
	return &recordpickerrecordformdrawerBuilder{
		&NodeBuilder{
			schema: &ui.Schema{
				Name:            name,
				Type:            ui.SchemaTypeVoid,
				XComponent:      ui.ComponentRecordPickerRecordFormDrawer,
				XComponentProps: make(map[string]interface{}),
				HandlerNames:    make([]string, 0),
			},
		},
	}
}

type recordpickerrecordformdrawerBuilder struct {
	*NodeBuilder
}

func (b *recordpickerrecordformdrawerBuilder) AC(f ac.AC) *recordpickerrecordformdrawerBuilder {
	b.schema.AC = f
	return b
}

func (b *recordpickerrecordformdrawerBuilder) ForInit(operation string, handler interface{}) *recordpickerrecordformdrawerBuilder {
	b.schema.XComponentProps["forInit"] = operation
	b.schema.HandlerNames = append(b.schema.HandlerNames, operation)
	controller.Bind(operation, handler)
	return b
}
