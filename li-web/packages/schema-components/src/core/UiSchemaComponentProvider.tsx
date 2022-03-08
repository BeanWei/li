import {
  Action,
  ActionBar,
  ArrayBase,
  ArrayItems,
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
  Grid,
  Input,
  InputNumber,
  Menu,
  Password,
  Radio,
  Select,
  Space,
  Submit,
  Switch,
  TimePicker,
} from "../components";
import { SchemaComponentOptions } from "./SchemaComponentOptions";

export const UiSchemaComponentProvider = (props: any) => {
  const { children } = props;
  return (
    <SchemaComponentOptions
      components={{
        Action,
        ActionBar,
        ArrayBase,
        ArrayItems,
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
        Grid,
        Input,
        InputNumber,
        Menu,
        Password,
        Radio,
        Select,
        Space,
        Submit,
        Switch,
        TimePicker,
      }}
    >
      {children}
    </SchemaComponentOptions>
  );
};
