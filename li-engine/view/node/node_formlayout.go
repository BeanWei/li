package node

import (
	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func FormLayout(name string) *formlayoutBuilder {
	return &formlayoutBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentFormLayout,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type formlayoutBuilder struct {
	*NodeBuilder
}

func (b *formlayoutBuilder) AC(f ac.AC) *formlayoutBuilder {
	b.schema.AC = f
	return b
}

func (b *formlayoutBuilder) Colon(colon bool) *formlayoutBuilder {
	b.schema.XComponentProps["colon"] = colon
	return b
}

func (b *formlayoutBuilder) LabelAlign(labelAlign string) *formlayoutBuilder {
	b.schema.XComponentProps["labelAlign"] = labelAlign
	return b
}

func (b *formlayoutBuilder) WrapperAlign(wrapperAlign string) *formlayoutBuilder {
	b.schema.XComponentProps["wrapperAlign"] = wrapperAlign
	return b
}

func (b *formlayoutBuilder) LabelWrap(labelWrap bool) *formlayoutBuilder {
	b.schema.XComponentProps["labelWrap"] = labelWrap
	return b
}

func (b *formlayoutBuilder) LabelWidth(labelWidth int) *formlayoutBuilder {
	b.schema.XComponentProps["labelWidth"] = labelWidth
	return b
}

func (b *formlayoutBuilder) WrapperWrap(wrapperWrap bool) *formlayoutBuilder {
	b.schema.XComponentProps["wrapperWrap"] = wrapperWrap
	return b
}

func (b *formlayoutBuilder) WrapperWidth(wrapperWidth int) *formlayoutBuilder {
	b.schema.XComponentProps["wrapperWidth"] = wrapperWidth
	return b
}

func (b *formlayoutBuilder) LabelCol(labelCol int) *formlayoutBuilder {
	b.schema.XComponentProps["labelCol"] = labelCol
	return b
}

func (b *formlayoutBuilder) WrapperCol(wrapperCol int) *formlayoutBuilder {
	b.schema.XComponentProps["wrapperCol"] = wrapperCol
	return b
}

func (b *formlayoutBuilder) Fullness(fullness bool) *formlayoutBuilder {
	b.schema.XComponentProps["fullness"] = fullness
	return b
}

func (b *formlayoutBuilder) Size(size string) *formlayoutBuilder {
	b.schema.XComponentProps["size"] = size
	return b
}

func (b *formlayoutBuilder) Layout(layout string) *formlayoutBuilder {
	b.schema.XComponentProps["layout"] = layout
	return b
}

func (b *formlayoutBuilder) Direction(direction string) *formlayoutBuilder {
	b.schema.XComponentProps["direction"] = direction
	return b
}

func (b *formlayoutBuilder) Inset(inset bool) *formlayoutBuilder {
	b.schema.XComponentProps["inset"] = inset
	return b
}

func (b *formlayoutBuilder) Shallow(shallow bool) *formlayoutBuilder {
	b.schema.XComponentProps["shallow"] = shallow
	return b
}

func (b *formlayoutBuilder) TooltipLayout(tooltipLayout string) *formlayoutBuilder {
	b.schema.XComponentProps["tooltipLayout"] = tooltipLayout
	return b
}

func (b *formlayoutBuilder) TooltipIcon(tooltipIcon string) *formlayoutBuilder {
	b.schema.XComponentProps["tooltipIcon"] = tooltipIcon
	return b
}

func (b *formlayoutBuilder) FeedbackLayout(feedbackLayout string) *formlayoutBuilder {
	b.schema.XComponentProps["feedbackLayout"] = feedbackLayout
	return b
}

func (b *formlayoutBuilder) Bordered(bordered bool) *formlayoutBuilder {
	b.schema.XComponentProps["bordered"] = bordered
	return b
}

func (b *formlayoutBuilder) Breakpoints(breakpoints ...int) *formlayoutBuilder {
	b.schema.XComponentProps["breakpoints"] = breakpoints
	return b
}

func (b *formlayoutBuilder) SpaceGap(spaceGap int) *formlayoutBuilder {
	b.schema.XComponentProps["spaceGap"] = spaceGap
	return b
}

func (b *formlayoutBuilder) GridColumnGap(gridColumnGap int) *formlayoutBuilder {
	b.schema.XComponentProps["gridColumnGap"] = gridColumnGap
	return b
}

func (b *formlayoutBuilder) GridRowGap(gridRowGap int) *formlayoutBuilder {
	b.schema.XComponentProps["gridRowGap"] = gridRowGap
	return b
}
