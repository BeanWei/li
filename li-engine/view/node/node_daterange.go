package node

import "github.com/BeanWei/li/li-engine/view/ui"

func DateRange(name string) *daterangeBuilder {
	return &daterangeBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			XPath:           name,
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

func (b *daterangeBuilder) AllowClear() *daterangeBuilder {
	b.schema.XComponentProps["allowClear"] = true
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

func (b *daterangeBuilder) ShortcutsPlacementLeft() *daterangeBuilder {
	b.schema.XComponentProps["shortcutsPlacementLeft"] = true
	return b
}

func (b *daterangeBuilder) Size(size string) *daterangeBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *daterangeBuilder) HideNotInViewDates() *daterangeBuilder {
	b.schema.XComponentProps["hideNotInViewDates"] = true
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
