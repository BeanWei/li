package node

// A Descriptor for field configuration.
type Descriptor struct {
	Name           string
	SchemaType     SchemaType
	SchemaProps    *SchemaProps
	ComponentName  string
	ComponentProps map[string]interface{}
	Properties     []*Descriptor
}

func (d *Descriptor) ToSchema() map[string]interface{} {
	schema := map[string]interface{}{
		"type":              d.SchemaType.String(),
		"x-component":       d.ComponentName,
		"x-component-props": d.ComponentProps,
	}
	if d.SchemaProps != nil {
		schema["title"] = d.SchemaProps.Title
	}
	pl := len(d.Properties)
	if pl > 0 {
		properties := make(map[string]interface{}, pl)
		for _, dd := range d.Properties {
			properties[dd.Name] = dd.ToSchema()
		}
		schema["properties"] = properties
	}
	return schema
}
