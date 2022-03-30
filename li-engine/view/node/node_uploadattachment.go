package node

import "github.com/BeanWei/li/li-engine/view/ui"

func UploadAttachment(name string) *uploadattachmentBuilder {
	return &uploadattachmentBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:       name,
			XPath:      name,
			Type:       ui.SchemaTypeArray,
			XComponent: ui.ComponentUploadAttachment,
			XDecorator: ui.DecoratorFormItem,
			XComponentProps: map[string]interface{}{
				"action": "/api/liql",
				"data": map[string]string{
					"operation": "@uploadFile",
				},
			},
		},
	}}
}

type uploadattachmentBuilder struct {
	*NodeBuilder
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

func (b *uploadattachmentBuilder) Directory() *uploadattachmentBuilder {
	b.schema.XComponentProps["directory"] = true
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

func (b *uploadattachmentBuilder) AutoUpload() *uploadattachmentBuilder {
	b.schema.XComponentProps["autoUpload"] = true
	return b
}

func (b *uploadattachmentBuilder) Limit(limit int) *uploadattachmentBuilder {
	b.schema.XComponentProps["limit"] = limit
	return b
}

func (b *uploadattachmentBuilder) Drag() *uploadattachmentBuilder {
	b.schema.XComponentProps["drag"] = true
	return b
}

func (b *uploadattachmentBuilder) Multiple() *uploadattachmentBuilder {
	b.schema.XComponentProps["multiple"] = true
	return b
}

func (b *uploadattachmentBuilder) Tip(tip string) *uploadattachmentBuilder {
	b.schema.XComponentProps["tip"] = tip
	return b
}
