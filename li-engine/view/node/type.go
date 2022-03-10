package node

type (
	SchemaType  uint8
	SchemaProps struct {
		Title            string
		Description      string
		Default          interface{}
		ReadOnly         bool
		WriteOnly        bool
		Enum             interface{}
		Const            interface{}
		MultipleOf       int
		Maximum          int
		ExclusiveMaximum int
		Minimum          int
		ExclusiveMinimum int
		MaxLength        int
		MinLength        int
		Pattern          string
		MaxItems         int
		MinItems         int
		UniqueItems      bool
		MaxProperties    int
		MinProperties    int
		Required         bool
		Format           string
		XIndex           int
		XPattern         string
		XDisplay         string
		XValidator       string
		XDecorator       string
		XDecoratorProps  map[string]interface{}
		XReactions       map[string]interface{}
		XContent         string
		XVisible         bool
		XHidden          bool
		XDisabled        bool
		XEditable        bool
		XReadOnly        bool
		XReadPretty      bool
		XData            map[string]interface{}
		Properties       map[string]*Descriptor
	}
)

const (
	SchemaTypeAny = iota
	SchemaTypeString
	SchemaTypeObject
	SchemaTypeArray
	SchemaTypeNumber
	SchemaTypeBool
	SchemaTypeVoid
	SchemaTypeDate
	SchemaTypeDatetime
	schemaEndTypes
)

func (t SchemaType) String() string {
	if t < schemaEndTypes {
		return schematypeNames[t]
	}
	return schematypeNames[SchemaTypeAny]
}

var (
	schematypeNames = [...]string{
		SchemaTypeAny:      "(string & {})",
		SchemaTypeString:   "string",
		SchemaTypeObject:   "object",
		SchemaTypeArray:    "array",
		SchemaTypeNumber:   "number",
		SchemaTypeBool:     "boolean",
		SchemaTypeVoid:     "void",
		SchemaTypeDate:     "date",
		SchemaTypeDatetime: "datetime",
	}
)
