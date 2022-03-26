package node

import "github.com/BeanWei/li/li-engine/view/ui"

func FormGrid(name string) *formgridBuilder {
	return &formgridBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentFormGrid,
			XComponentProps: make(map[string]interface{}),
		},
	}}
}

type formgridBuilder struct {
	*NodeBuilder
}

func (b *formgridBuilder) MaxRows(maxRows int) *formgridBuilder {
	b.schema.XComponentProps["maxRows"] = maxRows
	return b
}

func (b *formgridBuilder) MaxColumns(maxColumns ...int) *formgridBuilder {
	b.schema.XComponentProps["maxColumns"] = maxColumns
	return b
}

func (b *formgridBuilder) MinColumns(minColumns ...int) *formgridBuilder {
	b.schema.XComponentProps["minColumns"] = minColumns
	return b
}

func (b *formgridBuilder) MaxWidth(maxWidth ...int) *formgridBuilder {
	b.schema.XComponentProps["maxWidth"] = maxWidth
	return b
}

func (b *formgridBuilder) MinWidth(minWidth ...int) *formgridBuilder {
	b.schema.XComponentProps["minWidth"] = minWidth
	return b
}

func (b *formgridBuilder) Breakpoints(breakpoints ...int) *formgridBuilder {
	b.schema.XComponentProps["breakpoints"] = breakpoints
	return b
}

func (b *formgridBuilder) ColumnGap(columnGap ...int) *formgridBuilder {
	b.schema.XComponentProps["columnGap"] = columnGap
	return b
}

func (b *formgridBuilder) RowGap(rowGap int) *formgridBuilder {
	b.schema.XComponentProps["rowGap"] = rowGap
	return b
}

func (b *formgridBuilder) ColWrap() *formgridBuilder {
	b.schema.XComponentProps["colWrap"] = true
	return b
}

func (b *formgridBuilder) StrictAutoFit() *formgridBuilder {
	b.schema.XComponentProps["strictAutoFit"] = true
	return b
}
