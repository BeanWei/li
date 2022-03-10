package ui

// https://react.formilyjs.org/zh-CN/api/shared/schema
type (
	Schema struct {
		Name             string                 `json:"name"`
		Type             string                 `json:"type"`
		Title            string                 `json:"title"`
		Description      string                 `json:"description"`
		Default          interface{}            `json:"default"`
		ReadOnly         bool                   `json:"readOnly"`
		WriteOnly        bool                   `json:"writeOnly"`
		Enum             interface{}            `json:"enum"`
		Const            interface{}            `json:"const"`
		MultipleOf       int                    `json:"multipleOf"`
		Maximum          int                    `json:"maximum"`
		ExclusiveMaximum int                    `json:"exclusiveMaximum"`
		Minimum          int                    `json:"minimum"`
		ExclusiveMinimum int                    `json:"exclusiveMinimum"`
		MaxLength        int                    `json:"maxLength"`
		MinLength        int                    `json:"minLength"`
		Pattern          string                 `json:"pattern"`
		MaxItems         int                    `json:"maxItems"`
		MinItems         int                    `json:"minItems"`
		UniqueItems      bool                   `json:"uniqueItems"`
		MaxProperties    int                    `json:"maxProperties"`
		MinProperties    int                    `json:"minProperties"`
		Required         bool                   `json:"required"`
		Format           string                 `json:"format"`
		XIndex           int                    `json:"x-index"`
		XPattern         string                 `json:"x-pattern"`
		XDisplay         string                 `json:"x-display"`
		XValidator       string                 `json:"x-validator"`
		XDecorator       string                 `json:"x-decorator"`
		XDecoratorProps  map[string]interface{} `json:"x-decorator-props"`
		XComponent       string                 `json:"x-component"`
		XComponentProps  map[string]interface{} `json:"x-component-props"`
		XReactions       map[string]interface{} `json:"x-reactions"`
		XContent         string                 `json:"x-content"`
		XVisible         bool                   `json:"x-visible"`
		XHidden          bool                   `json:"x-hidden"`
		XDisabled        bool                   `json:"x-disabled"`
		XEditable        bool                   `json:"x-editable"`
		XReadOnly        bool                   `json:"x-read-only"`
		XReadPretty      bool                   `json:"x-read-pretty"`
		XData            map[string]interface{} `json:"x-data"`
		XOperation       string                 `json:"x-operation"` // 扩展属性, 定义数据处理的 handler name
		Properties       map[string]*Schema     `json:"properties"`
	}
)

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
)
