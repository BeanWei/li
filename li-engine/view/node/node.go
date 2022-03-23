package node

import "github.com/BeanWei/li/li-engine/view/ui"

func Node(name string) *NodeBuilder {
	return &NodeBuilder{schema: &ui.Schema{
		Name: name,
	}}
}

type NodeBuilder struct {
	schema *ui.Schema
}

func (b *NodeBuilder) Schema() *ui.Schema {
	return b.schema
}

func (b *NodeBuilder) SetSchema(schema *ui.Schema) *NodeBuilder {
	name := b.schema.Name
	b.schema = schema
	b.schema.Name = name
	return b
}

func (b *NodeBuilder) SetXReadPretty(xreadpretty bool) *NodeBuilder {
	b.schema.XReadPretty = xreadpretty
	return b
}

// func (d *Descriptor) ToSchema() map[string]interface{} {
// 	schema := map[string]interface{}{
// 		"type":              d.SchemaType.String(),
// 		"x-component":       d.ComponentName,
// 		"x-component-props": d.ComponentProps,
// 	}
// 	if d.SchemaProps != nil {
// 		schema["title"] = d.SchemaProps.Title
// 	}
// 	pl := len(d.Properties)
// 	if pl > 0 {
// 		properties := make(map[string]interface{}, pl)
// 		for _, dd := range d.Properties {
// 			properties[dd.Name] = dd.ToSchema()
// 		}
// 		schema["properties"] = properties
// 	}
// 	return schema
// }
