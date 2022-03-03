package field

type Type uint8

const (
	TypeInvalid Type = iota
	TypeString
	TypeBool
	TypeInt16
	TypeInt32
	TypeInt64
	TypeFloat32
	TypeFloat64
	TypeBigInt
	TypeDecimal
	TypeMap
	TypeStrings
	TypeObjects
	TypeUUID
	TypeBytes
	TypeDatetime
	TypeDuration
	TypeSequences
	TypeEnum
	TypeLink
	endTypes
)

func (t Type) String() string {
	if t < endTypes {
		return typeNames[t]
	}
	return typeNames[TypeInvalid]
}

var (
	typeNames = [...]string{
		TypeInvalid:   "invalid",
		TypeString:    "str",
		TypeBool:      "bool",
		TypeInt16:     "int16",
		TypeInt32:     "int32",
		TypeInt64:     "int64",
		TypeFloat32:   "float32",
		TypeFloat64:   "float64",
		TypeBigInt:    "bigint",
		TypeDecimal:   "decimal",
		TypeMap:       "json",
		TypeStrings:   "array<str>",
		TypeObjects:   "array<json>",
		TypeUUID:      "uuid",
		TypeBytes:     "bytes",
		TypeDatetime:  "datetime",
		TypeDuration:  "duration",
		TypeSequences: "sequences",
		TypeEnum:      "enum",
		TypeLink:      "link",
	}
)
