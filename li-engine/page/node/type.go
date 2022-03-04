package node

type SchemaType uint8

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
