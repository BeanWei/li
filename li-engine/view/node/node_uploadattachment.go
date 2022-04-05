package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func UploadAttachment(name string) *uploadattachmentBuilder {
	return &uploadattachmentBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentUploadAttachment,
			XComponentProps: map[string]interface{}{
				"action": "/api/liql",
				"data": map[string]string{
					"operation": "@uploadFile",
				},
			},
			XDecorator: ui.DecoratorFormItem,
		},
	}}
}

type uploadattachmentBuilder struct {
	*NodeBuilder
}

func (b *uploadattachmentBuilder) AC(f ac.AC) *uploadattachmentBuilder {
	b.schema.AC = f
	return b
}

func (b *uploadattachmentBuilder) Title(title string) *uploadattachmentBuilder {
	b.schema.Title = title
	return b
}

func (b *uploadattachmentBuilder) Description(description string) *uploadattachmentBuilder {
	b.schema.Description = description
	return b
}

func (b *uploadattachmentBuilder) Default(value interface{}) *uploadattachmentBuilder {
	b.schema.Default = value
	return b
}

func (b *uploadattachmentBuilder) Directory(directory bool) *uploadattachmentBuilder {
	b.schema.XComponentProps["directory"] = directory
	return b
}

func (b *uploadattachmentBuilder) Accept(accept string) *uploadattachmentBuilder {
	b.schema.XComponentProps["accept"] = accept
	return b
}

func (b *uploadattachmentBuilder) ListType(listType string) *uploadattachmentBuilder {
	b.schema.XComponentProps["listType"] = listType
	return b
}

func (b *uploadattachmentBuilder) AutoUpload(autoUpload bool) *uploadattachmentBuilder {
	b.schema.XComponentProps["autoUpload"] = autoUpload
	return b
}

func (b *uploadattachmentBuilder) Limit(limit int) *uploadattachmentBuilder {
	b.schema.XComponentProps["limit"] = limit
	return b
}

func (b *uploadattachmentBuilder) Drag(drag bool) *uploadattachmentBuilder {
	b.schema.XComponentProps["drag"] = drag
	return b
}

func (b *uploadattachmentBuilder) Multiple(multiple bool) *uploadattachmentBuilder {
	b.schema.XComponentProps["multiple"] = multiple
	return b
}

func (b *uploadattachmentBuilder) Tip(tip string) *uploadattachmentBuilder {
	b.schema.XComponentProps["tip"] = tip
	return b
}
