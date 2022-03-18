package com

import "github.com/BeanWei/li/li-engine/view/ui"

func ThemeSwitch(name string) *themeswitchBuilder {
	return &themeswitchBuilder{schema: &ui.Schema{
		Name:       name,
		Type:       ui.SchemaTypeVoid,
		XComponent: "ThemeSwitch",
	}}
}

type themeswitchBuilder struct {
	schema *ui.Schema
}

func (b *themeswitchBuilder) Schema() *ui.Schema {
	return b.schema
}
