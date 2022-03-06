import {
  Action,
  ActionBar,
  Card,
  Cascader,
  Checkbox,
  ColorSelect,
  DatePicker,
  Form,
  FormButtonGroup,
  FormGrid,
  FormItem,
  FormLayout,
  Input,
  InputNumber,
  Menu,
  Password,
  Radio,
  Select,
  Space,
  Switch,
  TimePicker,
} from "../components";
import { SchemaComponentOptions } from "./SchemaComponentOptions";

export const UiSchemaComponentProvider = (props: any) => {
  const { children } = props;
  return (
    <SchemaComponentOptions
      components={{
        Action: Action,
        ActionBar: ActionBar,
        Card: Card,
        Cascader: Cascader,
        Checkbox: Checkbox,
        ColorSelect: ColorSelect,
        DatePicker: DatePicker,
        Form: Form,
        FormButtonGroup: FormButtonGroup,
        FormGrid: FormGrid,
        FormItem: FormItem,
        FormLayout: FormLayout,
        Input: Input,
        InputNumber: InputNumber,
        Menu: Menu,
        Password: Password,
        Radio: Radio,
        Select: Select,
        Space: Space,
        Switch: Switch,
        TimePicker: TimePicker,
      }}
    >
      {children}
    </SchemaComponentOptions>
  );
};
