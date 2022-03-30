package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Time(name string) *timeBuilder {
	return &timeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			XPath:           name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentTimePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type timeBuilder struct {
	*NodeBuilder
}

func (b *timeBuilder) Title(title string) *timeBuilder {
	b.schema.Title = title
	return b
}

func (b *timeBuilder) Description(description string) *timeBuilder {
	b.schema.Description = description
	return b
}

func (b *timeBuilder) Default(value interface{}) *timeBuilder {
	b.schema.Default = value
	return b
}

func (b *timeBuilder) Format(format string) *timeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}

func (b *timeBuilder) AllowClear() *timeBuilder {
	b.schema.XComponentProps["allowClear"] = true
	return b
}

func (b *timeBuilder) DisableConfirm() *timeBuilder {
	b.schema.XComponentProps["disableConfirm"] = true
	return b
}

func (b *timeBuilder) Position(position string) *timeBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *timeBuilder) Placeholder(placeholder string) *timeBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *timeBuilder) Use12Hours() *timeBuilder {
	b.schema.XComponentProps["use12Hours"] = true
	return b
}

func (b *timeBuilder) DideDisabledOptions() *timeBuilder {
	b.schema.XComponentProps["hideDisabledOptions"] = true
	return b
}

func (b *timeBuilder) Size(size string) *timeBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *timeBuilder) Extra(extra string) *timeBuilder {
	b.schema.XComponentProps["extra"] = extra
	return b
}

func (b *timeBuilder) UtcOffset(utcOffset int) *timeBuilder {
	b.schema.XComponentProps["utcOffset"] = utcOffset
	return b
}

func (b *timeBuilder) Timezone(timezone string) *timeBuilder {
	b.schema.XComponentProps["timezone"] = timezone
	return b
}

func (b *timeBuilder) ShowNowBtn() *timeBuilder {
	b.schema.XComponentProps["showNowBtn"] = true
	return b
}
