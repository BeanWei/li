package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func DateRange(name string) *daterangeBuilder {
	return &daterangeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeArray,
			XComponent:      ui.ComponentDatePickerRangePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type daterangeBuilder struct {
	*NodeBuilder
}

func (b *daterangeBuilder) AC(f ac.AC) *daterangeBuilder {
	b.schema.AC = f
	return b
}

func (b *daterangeBuilder) Title(title string) *daterangeBuilder {
	b.schema.Title = title
	return b
}

func (b *daterangeBuilder) Description(description string) *daterangeBuilder {
	b.schema.Description = description
	return b
}

func (b *daterangeBuilder) Default(value interface{}) *daterangeBuilder {
	b.schema.Default = value
	return b
}

// Mode Time, Week, Month, Quarter, Year
func (b *daterangeBuilder) Mode(mode string) *daterangeBuilder {
	b.schema.XComponentProps["mode"] = mode
	return b
}

func (b *daterangeBuilder) Format(format string) *daterangeBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}

func (b *daterangeBuilder) AllowClear(allowClear bool) *daterangeBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *daterangeBuilder) DayStartOfWeek(dayStartOfWeek int8) *daterangeBuilder {
	b.schema.XComponentProps["dayStartOfWeek"] = dayStartOfWeek
	return b
}

func (b *daterangeBuilder) Position(position string) *daterangeBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *daterangeBuilder) ShortcutsPlacementLeft(shortcutsPlacementLeft bool) *daterangeBuilder {
	b.schema.XComponentProps["shortcutsPlacementLeft"] = shortcutsPlacementLeft
	return b
}

func (b *daterangeBuilder) Size(size string) *daterangeBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *daterangeBuilder) HideNotInViewDates(hideNotInViewDates bool) *daterangeBuilder {
	b.schema.XComponentProps["hideNotInViewDates"] = hideNotInViewDates
	return b
}

func (b *daterangeBuilder) UtcOffset(utcOffset string) *daterangeBuilder {
	b.schema.XComponentProps["utcOffset"] = utcOffset
	return b
}

func (b *daterangeBuilder) Timezone(timezone string) *daterangeBuilder {
	b.schema.XComponentProps["timezone"] = timezone
	return b
}
