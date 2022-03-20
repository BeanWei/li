package node

import (
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
)

func ListTableColumn(name string) *listtablecolumnBuilder {
	return &listtablecolumnBuilder{schema: &ui.Schema{
		Name:            name,
		Type:            ui.SchemaTypeVoid,
		XComponent:      ui.ComponentListTableColumn,
		XComponentProps: make(map[string]interface{}),
		Properties:      make(map[string]*ui.Schema),
	}}
}

type listtablecolumnBuilder struct {
	schema *ui.Schema
}

func (b *listtablecolumnBuilder) Schema() *ui.Schema {
	return b.schema
}

// Title 列标题
func (b *listtablecolumnBuilder) Title(title string) *listtablecolumnBuilder {
	b.schema.XComponentProps["title"] = title
	return b
}

// DataIndex .
func (b *listtablecolumnBuilder) DataIndex(dataIndex string) *listtablecolumnBuilder {
	b.schema.XComponentProps["dataIndex"] = dataIndex
	return b
}

// Align 设置列的对齐方式 'left' | 'center' | 'right'
func (b *listtablecolumnBuilder) Align(align string) *listtablecolumnBuilder {
	b.schema.XComponentProps["align"] = align
	return b
}

// Ellipsis 单元格内容超出长度后，是否自动省略
func (b *listtablecolumnBuilder) Ellipsis() *listtablecolumnBuilder {
	b.schema.XComponentProps["ellipsis"] = true
	return b
}

// Width 列宽度
func (b *listtablecolumnBuilder) Width(width int) *listtablecolumnBuilder {
	b.schema.XComponentProps["width"] = width
	return b
}

func (b *listtablecolumnBuilder) Render(element view.Node) *listtablecolumnBuilder {
	b.schema.Properties[element.Schema().Name] = element.Schema()
	return b
}
