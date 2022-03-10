package ui

// https://react.formilyjs.org/zh-CN/api/shared/schema
type Schema struct {
	Name             string                 `json:"name,omitempty"`
	Type             string                 `json:"type,omitempty"`
	Title            string                 `json:"title,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Default          interface{}            `json:"default,omitempty"`
	ReadOnly         bool                   `json:"readOnly,omitempty"`
	WriteOnly        bool                   `json:"writeOnly,omitempty"`
	Enum             interface{}            `json:"enum,omitempty"`
	Const            interface{}            `json:"const,omitempty"`
	MultipleOf       int                    `json:"multipleOf,omitempty"`
	Maximum          int                    `json:"maximum,omitempty"`
	ExclusiveMaximum int                    `json:"exclusiveMaximum,omitempty"`
	Minimum          int                    `json:"minimum,omitempty"`
	ExclusiveMinimum int                    `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                    `json:"maxLength,omitempty"`
	MinLength        int                    `json:"minLength,omitempty"`
	Pattern          string                 `json:"pattern,omitempty"`
	MaxItems         int                    `json:"maxItems,omitempty"`
	MinItems         int                    `json:"minItems,omitempty"`
	UniqueItems      bool                   `json:"uniqueItems,omitempty"`
	MaxProperties    int                    `json:"maxProperties,omitempty"`
	MinProperties    int                    `json:"minProperties,omitempty"`
	Required         bool                   `json:"required,omitempty"`
	Format           string                 `json:"format,omitempty"`
	XIndex           int                    `json:"x-index,omitempty"`
	XPattern         string                 `json:"x-pattern,omitempty"`
	XDisplay         string                 `json:"x-display,omitempty"`
	XValidator       string                 `json:"x-validator,omitempty"`
	XDecorator       string                 `json:"x-decorator,omitempty"`
	XDecoratorProps  map[string]interface{} `json:"x-decorator-props,omitempty"`
	XComponent       string                 `json:"x-component,omitempty"`
	XComponentProps  map[string]interface{} `json:"x-component-props,omitempty"`
	XReactions       map[string]interface{} `json:"x-reactions,omitempty"`
	XContent         string                 `json:"x-content,omitempty"`
	XVisible         bool                   `json:"x-visible,omitempty"`
	XHidden          bool                   `json:"x-hidden,omitempty"`
	XDisabled        bool                   `json:"x-disabled,omitempty"`
	XEditable        bool                   `json:"x-editable,omitempty"`
	XReadOnly        bool                   `json:"x-read-only,omitempty"`
	XReadPretty      bool                   `json:"x-read-pretty,omitempty"`
	XData            map[string]interface{} `json:"x-data,omitempty"`
	XOperation       string                 `json:"x-operation,omitempty"` // 扩展属性, 定义数据处理的 handler name
	Properties       map[string]*Schema     `json:"properties,omitempty"`
}

const (
	SchemaTypeAny      = "(string & {})"
	SchemaTypeString   = "string"
	SchemaTypeObject   = "object"
	SchemaTypeArray    = "array"
	SchemaTypeNumber   = "number"
	SchemaTypeBool     = "boolean"
	SchemaTypeVoid     = "void"
	SchemaTypeDate     = "date"
	SchemaTypeDatetime = "datetime"
)

const (
	ComponentListTable = "List.Table"
	ComponentListCard  = "List.Card"
	ComponentForm      = "Form"
	ComponentCheckbox  = "Checkbox"
	ComponentGridRow   = "Grid.Row"
	ComponentGridCol   = "Grid.Col"
)
