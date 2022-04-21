package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func Date(name string) *dateBuilder {
	return &dateBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeString,
			XComponent:      ui.ComponentDatePicker,
			XComponentProps: make(map[string]interface{}),
			XDecorator:      ui.DecoratorFormItem,
		},
	}}
}

type dateBuilder struct {
	*NodeBuilder
}

func (b *dateBuilder) AC(f ac.AC) *dateBuilder {
	b.schema.AC = f
	return b
}

func (b *dateBuilder) Title(title string) *dateBuilder {
	b.SetTitle(title)
	return b
}

func (b *dateBuilder) Description(description string) *dateBuilder {
	b.SetDescription(description)
	return b
}

func (b *dateBuilder) Default(value interface{}) *dateBuilder {
	b.schema.Default = value
	return b
}

// Mode Time, Week, Month, Quarter, Year
func (b *dateBuilder) Mode(mode string) *dateBuilder {
	b.schema.XComponentProps["mode"] = mode
	return b
}

func (b *dateBuilder) Format(format string) *dateBuilder {
	b.schema.XComponentProps["format"] = format
	return b
}

func (b *dateBuilder) AllowClear(allowClear bool) *dateBuilder {
	b.schema.XComponentProps["allowClear"] = allowClear
	return b
}

func (b *dateBuilder) DayStartOfWeek(dayStartOfWeek int8) *dateBuilder {
	b.schema.XComponentProps["dayStartOfWeek"] = dayStartOfWeek
	return b
}

func (b *dateBuilder) Position(position string) *dateBuilder {
	b.schema.XComponentProps["position"] = position
	return b
}

func (b *dateBuilder) ShortcutsPlacementLeft(shortcutsPlacementLeft bool) *dateBuilder {
	b.schema.XComponentProps["shortcutsPlacementLeft"] = shortcutsPlacementLeft
	return b
}

func (b *dateBuilder) Size(size string) *dateBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *dateBuilder) HideNotInViewDates(hideNotInViewDates bool) *dateBuilder {
	b.schema.XComponentProps["hideNotInViewDates"] = hideNotInViewDates
	return b
}

func (b *dateBuilder) UtcOffset(utcOffset int) *dateBuilder {
	b.schema.XComponentProps["utcOffset"] = utcOffset
	return b
}

func (b *dateBuilder) Timezone(timezone string) *dateBuilder {
	b.schema.XComponentProps["timezone"] = timezone
	return b
}
