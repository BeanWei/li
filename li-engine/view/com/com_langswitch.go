package com

import "github.com/BeanWei/li/li-engine/view/ui"

func LangSwitch(name string) *langswitchBuilder {
	return &langswitchBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: "LangSwitch",
	}}
}

type langswitchBuilder struct {
	schema *ui.Schema
}

func (b *langswitchBuilder) Schema() *ui.Schema {
	return b.schema
}
