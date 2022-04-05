package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Time(name string) *timeBuilder {
	return &timeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
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

func (b *timeBuilder) AC(f ac.AC) *timeBuilder {
	b.schema.AC = f
	return b
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

func (b *timeBuilder) AllowClear(allowClear bool) *timeBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *timeBuilder) DisableConfirm(disableConfirm bool) *timeBuilder {
	b.schema.XComponentProps["disableConfirm"] = disableConfirm
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

func (b *timeBuilder) Use12Hours(use12Hours bool) *timeBuilder {
	b.schema.XComponentProps["use12Hours"] = use12Hours
	return b
}

func (b *timeBuilder) DideDisabledOptions(hideDisabledOptions bool) *timeBuilder {
	b.schema.XComponentProps["hideDisabledOptions"] = hideDisabledOptions
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

func (b *timeBuilder) ShowNowBtn(showNowBtn bool) *timeBuilder {
	b.schema.XComponentProps["showNowBtn"] = showNowBtn
	return b
}
