package node

import "github.com/BeanWei/li/li-engine/view/ui"

func TimeRange(name string) *timerangeBuilder {
	return &timerangeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentDatePickerRangePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type timerangeBuilder struct {
	*NodeBuilder
}

func (b *timerangeBuilder) Title(title string) *timerangeBuilder {
	b.schema.Title = title
	return b
}

func (b *timerangeBuilder) Description(description string) *timerangeBuilder {
	b.schema.Description = description
	return b
}

func (b *timerangeBuilder) Default(value interface{}) *timerangeBuilder {
	b.schema.Default = value
	return b
}

func (b *timerangeBuilder) Format(format string) *timerangeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}

func (b *timerangeBuilder) AllowClear() *timerangeBuilder {
	b.schema.XComponentProps["allowClear"] = true
	return b
}

func (b *timerangeBuilder) DisableConfirm() *timerangeBuilder {
	b.schema.XComponentProps["disableConfirm"] = true
	return b
}

func (b *timerangeBuilder) Position(position string) *timerangeBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *timerangeBuilder) Placeholder(placeholder ...string) *timerangeBuilder {
	b.schema.XComponentProps["placeholder"] = placeholder
	return b
}

func (b *timerangeBuilder) Use12Hours() *timerangeBuilder {
	b.schema.XComponentProps["use12Hours"] = true
	return b
}

func (b *timerangeBuilder) DideDisabledOptions() *timerangeBuilder {
	b.schema.XComponentProps["hideDisabledOptions"] = true
	return b
}

func (b *timerangeBuilder) Size(size string) *timerangeBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *timerangeBuilder) Extra(extra string) *timerangeBuilder {
	b.schema.XComponentProps["extra"] = extra
	return b
}

func (b *timerangeBuilder) UtcOffset(utcOffset int) *timerangeBuilder {
	b.schema.XComponentProps["utcOffset"] = utcOffset
	return b
}

func (b *timerangeBuilder) Timezone(timezone string) *timerangeBuilder {
	b.schema.XComponentProps["timezone"] = timezone
	return b
}

func (b *timerangeBuilder) Order() *timerangeBuilder {
	b.schema.XComponentProps["order"] = true
	return b
}
