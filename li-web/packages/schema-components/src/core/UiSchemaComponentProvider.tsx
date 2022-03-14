import {
  Action,
  ActionBar,
  ArrayBase,
  ArrayItems,
  CardItem,
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
  List,
  Menu,
  Password,
  Radio,
  Reset,
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
        CardItem,
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
        List,
        Menu,
        Password,
        Radio,
        Reset,
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
