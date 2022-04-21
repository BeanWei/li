package node

import (
	"strings"

	"github.com/BeanWei/li/li-engine/ac"
	"github.com/BeanWei/li/li-engine/view"
	"github.com/BeanWei/li/li-engine/view/ui"
	"github.com/gogf/gf/v2/container/gmap"
)

func ListTableColumn(name string) *listtablecolumnBuilder {
	return &listtablecolumnBuilder{&NodeBuilder{
		schema: &ui.Schema{
			Name:            name,
			Type:            ui.SchemaTypeVoid,
			XComponent:      ui.ComponentListTableColumn,
			XComponentProps: make(map[string]interface{}),
			Properties:      gmap.NewListMap(),
		},
	}}
}

type listtablecolumnBuilder struct {
	*NodeBuilder
}

func (b *listtablecolumnBuilder) AC(f ac.AC) *listtablecolumnBuilder {
	b.schema.AC = f
	return b
}

// Title 列标题
func (b *listtablecolumnBuilder) Title(title string) *listtablecolumnBuilder {
	if !strings.HasPrefix(title, "{{t('") {
		title = "{{t('" + title + "')}}"
	}
	b.schema.XComponentProps["title"] = title
	return b
}

// DataIndex .
func (b *listtablecolumnBuilder) DataIndex(dataIndex string) *listtablecolumnBuilder {
	b.schema.XComponentProps["dataIndex"] = dataIndex
	return b
}

func (b *listtablecolumnBuilder) Order(order int) *listtablecolumnBuilder {
	b.schema.XComponentProps["order"] = order
	return b
}

// Align 设置列的对齐方式 'left' | 'center' | 'right'
func (b *listtablecolumnBuilder) Align(align string) *listtablecolumnBuilder {
	b.schema.XComponentProps["align"] = align
	return b
}

// Ellipsis 单元格内容超出长度后，是否自动省略
func (b *listtablecolumnBuilder) Ellipsis(ellipsis bool) *listtablecolumnBuilder {
	b.schema.XComponentProps["ellipsis"] = ellipsis
	return b
}

// Width 列宽度
func (b *listtablecolumnBuilder) Width(width int) *listtablecolumnBuilder {
	b.schema.XComponentProps["width"] = width
	return b
}

func (b *listtablecolumnBuilder) Render(element view.Node) *listtablecolumnBuilder {
	b.schema.Properties.Set(element.Schema().Name, element.Schema())
	return b
}

func (b *listtablecolumnBuilder) Filterable(filterable bool) *listtablecolumnBuilder {
	b.schema.XComponentProps["filterable"] = filterable
	return b
}

func (b *listtablecolumnBuilder) Sortable(sortable bool) *listtablecolumnBuilder {
	b.schema.XComponentProps["sortable"] = sortable
	return b
}

func (b *listtablecolumnBuilder) HideInTable(hideInTable bool) *listtablecolumnBuilder {
	b.schema.XComponentProps["hideInTable"] = hideInTable
	return b
}
