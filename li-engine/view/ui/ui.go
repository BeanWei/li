package ui

// https://react.formilyjs.org/zh-CN/api/shared/schema
type Schema struct {
	Name             string                   `json:"name,omitempty"`
	Type             string                   `json:"type,omitempty"`
	Title            string                   `json:"title,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Default          interface{}              `json:"default,omitempty"`
	ReadOnly         bool                     `json:"readOnly,omitempty"`
	WriteOnly        bool                     `json:"writeOnly,omitempty"`
	Enum             []map[string]interface{} `json:"enum,omitempty"`
	Const            interface{}              `json:"const,omitempty"`
	MultipleOf       int                      `json:"multipleOf,omitempty"`
	Maximum          int                      `json:"maximum,omitempty"`
	ExclusiveMaximum int                      `json:"exclusiveMaximum,omitempty"`
	Minimum          int                      `json:"minimum,omitempty"`
	ExclusiveMinimum int                      `json:"exclusiveMinimum,omitempty"`
	MaxLength        int                      `json:"maxLength,omitempty"`
	MinLength        int                      `json:"minLength,omitempty"`
	Pattern          string                   `json:"pattern,omitempty"`
	MaxItems         int                      `json:"maxItems,omitempty"`
	MinItems         int                      `json:"minItems,omitempty"`
	UniqueItems      bool                     `json:"uniqueItems,omitempty"`
	MaxProperties    int                      `json:"maxProperties,omitempty"`
	MinProperties    int                      `json:"minProperties,omitempty"`
	Required         bool                     `json:"required,omitempty"`
	Format           string                   `json:"format,omitempty"`
	XIndex           int                      `json:"x-index,omitempty"`
	XPattern         string                   `json:"x-pattern,omitempty"`
	XDisplay         string                   `json:"x-display,omitempty"`
	XValidator       string                   `json:"x-validator,omitempty"`
	XDecorator       string                   `json:"x-decorator,omitempty"`
	XDecoratorProps  map[string]interface{}   `json:"x-decorator-props,omitempty"`
	XComponent       string                   `json:"x-component,omitempty"`
	XComponentProps  map[string]interface{}   `json:"x-component-props,omitempty"`
	XReactions       map[string]interface{}   `json:"x-reactions,omitempty"`
	XContent         string                   `json:"x-content,omitempty"`
	XVisible         bool                     `json:"x-visible,omitempty"`
	XHidden          bool                     `json:"x-hidden,omitempty"`
	XDisabled        bool                     `json:"x-disabled,omitempty"`
	XEditable        bool                     `json:"x-editable,omitempty"`
	XReadOnly        bool                     `json:"x-read-only,omitempty"`
	XReadPretty      bool                     `json:"x-read-pretty,omitempty"`
	XData            map[string]interface{}   `json:"x-data,omitempty"`
	Properties       map[string]*Schema       `json:"properties,omitempty"`
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
	DecoratorFormItem = "FormItem"
	DecoratorCardItem = "CardItem"
)

const (
	ComponentActionFormDrawer           = "Action.FormDrawer"
	ComponentActionFormModal            = "Action.FormModal"
	ComponentArrayItems                 = "ArrayItems"
	ComponentArrayTable                 = "ArrayTable"
	ComponentCascader                   = "Cascader"
	ComponentCheckbox                   = "Checkbox"
	ComponentCheckboxGroup              = "Checkbox.Group"
	ComponentColorSelect                = "ColorSelect"
	ComponentDatePicker                 = "DatePicker"
	ComponentDatePickerRangePicker      = "DatePicker.RangePicker"
	ComponentDivider                    = "Divider"
	ComponentDropdownMenu               = "DropdownMenu"
	ComponentDropdownMenuItem           = "DropdownMenu.Item"
	ComponentDropdownMenuSubMenu        = "DropdownMenu.SubMenu"
	ComponentDropdownMenuURL            = "DropdownMenu.URL"
	ComponentForm                       = "Form"
	ComponentFormButtonGroup            = "FormButtonGroup"
	ComponentFormDrawer                 = "FormDrawer"
	ComponentFormGrid                   = "FormGrid"
	ComponentFormLayout                 = "FormLayout"
	ComponentFormModal                  = "FormModal"
	ComponentGridRow                    = "Grid.Row"
	ComponentGridCol                    = "Grid.Col"
	ComponentInput                      = "Input"
	ComponentInputTextArea              = "Input.TextArea"
	ComponentInputNumber                = "InputNumber"
	ComponentInputTag                   = "InputTag"
	ComponentList                       = "List"
	ComponentListAction                 = "List.Action"
	ComponentListActionRowSelection     = "List.Action.RowSelection"
	ComponentListActionRefresh          = "List.Action.Refresh"
	ComponentListActionRecordEditDrawer = "List.Action.RecordEditDrawer"
	ComponentListActionRecordEditModal  = "List.Action.RecordEditModal"
	ComponentListActionRecordDelete     = "List.Action.RecordDelete"
	ComponentListTable                  = "List.Table"
	ComponentListTableColumn            = "List.Table.Column"
	ComponentListCard                   = "List.Card"
	ComponentPassword                   = "Password"
	ComponentRadio                      = "Radio"
	ComponentRadioGroup                 = "Radio.Group"
	ComponentRate                       = "Rate"
	ComponentSelect                     = "Select"
	ComponentSpace                      = "Space"
	ComponentSubmit                     = "Submit"
	ComponentSwitch                     = "Switch"
	ComponentTimePicker                 = "TimePicker"
	ComponentTimePickerRangPicker       = "TimePicker.RangPicker"
)

func NewSchema(xcomponent string) *Schema {
	return &Schema{
		XComponent: xcomponent,
	}
}

func (s *Schema) SetName(name string) *Schema {
	s.Name = name
	return s
}

func (s *Schema) SetTitle(title string) *Schema {
	s.Title = title
	return s
}

func (s *Schema) SetXComponent(xcomponent string) *Schema {
	s.XComponent = xcomponent
	return s
}
