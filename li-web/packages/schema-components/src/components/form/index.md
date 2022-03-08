# Form

## Example 1.

```tsx
import React from "react";
import {
  FormGrid,
  Input,
  DatePicker,
  Cascader,
  Select,
  FormItem,
  FormLayout,
  SchemaComponentProvider,
  SchemaComponent,
} from "schema-components";
import { Button } from "@arco-design/web-react";
import "@arco-design/web-react/dist/css/arco.css";

const schema = {
  type: "object",
  properties: {
    username: {
      type: "string",
      title: "用户名",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
    },
    name: {
      type: "void",
      title: "姓名",
      "x-decorator": "FormItem",
      "x-decorator-props": {
        asterisk: true,
        feedbackLayout: "none",
      },
      "x-component": "FormGrid",
      properties: {
        firstName: {
          type: "string",
          required: true,
          "x-decorator": "FormItem",
          "x-component": "Input",
          "x-component-props": {
            placeholder: "姓",
          },
        },
        lastName: {
          type: "string",
          required: true,
          "x-decorator": "FormItem",
          "x-component": "Input",
          "x-component-props": {
            placeholder: "名",
          },
        },
      },
    },
    email: {
      type: "string",
      title: "邮箱",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
      "x-validator": "email",
    },
    gender: {
      type: "string",
      title: "性别",
      enum: [
        {
          label: "男",
          value: 1,
        },
        {
          label: "女",
          value: 2,
        },
        {
          label: "第三性别",
          value: 3,
        },
      ],
      "x-decorator": "FormItem",
      "x-component": "Select",
    },
    birthday: {
      type: "string",
      required: true,
      title: "生日",
      "x-decorator": "FormItem",
      "x-component": "DatePicker",
    },
    address: {
      type: "string",
      required: true,
      title: "地址",
      "x-decorator": "FormItem",
      "x-component": "Cascader",
    },
  },
};

export default () => (
  <FormLayout labelCol={5} wrapperCol={16}>
    <SchemaComponentProvider
      components={{
        FormItem,
        FormGrid,
        FormLayout,
        Input,
        DatePicker,
        Cascader,
        Select,
      }}
    >
      <SchemaComponent schema={schema} />
    </SchemaComponentProvider>
  </FormLayout>
);
```
