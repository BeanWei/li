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
  ArrayTable,
  ArrayItems,
  Space,
  Submit,
  Form,
  FormItem,
  FormButtonGroup,
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
    array: {
      title: "数组1",
      type: "array",
      "x-decorator": "FormItem",
      "x-component": "ArrayTable",
      "x-component-props": {
        pagination: { pageSize: 10 },
        scroll: { x: "100%" },
      },
      items: {
        type: "object",
        properties: {
          column1: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": { width: 50, title: "Sort", align: "center" },
            properties: {
              sort: {
                type: "void",
                "x-component": "ArrayTable.SortHandle",
              },
            },
          },
          column2: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": { width: 80, title: "Index", align: "center" },
            properties: {
              index: {
                type: "void",
                "x-component": "ArrayTable.Index",
              },
            },
          },
          column3: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": { width: 200, title: "A1" },
            properties: {
              a1: {
                type: "string",
                "x-decorator": "FormItem",
                "x-component": "Input",
              },
            },
          },
          column4: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": { width: 200, title: "A2" },
            properties: {
              a2: {
                type: "string",
                "x-decorator": "FormItem",
                "x-component": "Input",
              },
            },
          },
          column5: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": { width: 200, title: "A3" },
            properties: {
              a3: {
                type: "string",
                "x-decorator": "FormItem",
                "x-component": "Input",
              },
            },
          },
          column6: {
            type: "void",
            "x-component": "ArrayTable.Column",
            "x-component-props": {
              title: "Operations",
              dataIndex: "operations",
              width: 200,
              fixed: "right",
            },
            properties: {
              item: {
                type: "void",
                "x-component": "FormItem",
                properties: {
                  remove: {
                    type: "void",
                    "x-component": "ArrayTable.Remove",
                  },
                  moveDown: {
                    type: "void",
                    "x-component": "ArrayTable.MoveDown",
                  },
                  moveUp: {
                    type: "void",
                    "x-component": "ArrayTable.MoveUp",
                  },
                },
              },
            },
          },
        },
      },
      properties: {
        add: {
          type: "void",
          "x-component": "ArrayTable.Addition",
          title: "添加条目",
        },
      },
    },
    array2: {
      title: "数组2",
      type: "array",
      "x-component": "ArrayItems",
      "x-decorator": "FormItem",
      title: "对象数组",
      items: {
        type: "object",
        properties: {
          space: {
            type: "void",
            "x-component": "Space",
            properties: {
              sort: {
                type: "void",
                "x-decorator": "FormItem",
                "x-component": "ArrayItems.SortHandle",
              },
              date: {
                type: "string",
                title: "日期",
                "x-decorator": "FormItem",
                "x-component": "DatePicker.RangePicker",
                "x-component-props": {
                  style: {
                    width: 160,
                  },
                },
              },
              input: {
                type: "string",
                title: "输入框",
                "x-decorator": "FormItem",
                "x-component": "Input",
              },
              select: {
                type: "string",
                title: "下拉框",
                enum: [
                  { label: "选项1", value: 1 },
                  { label: "选项2", value: 2 },
                ],
                "x-decorator": "FormItem",
                "x-component": "Select",
                "x-component-props": {
                  style: {
                    width: 160,
                  },
                },
              },
              remove: {
                type: "void",
                "x-decorator": "FormItem",
                "x-component": "ArrayItems.Remove",
              },
            },
          },
        },
      },
      properties: {
        add: {
          type: "void",
          title: "添加条目",
          "x-component": "ArrayItems.Addition",
        },
      },
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
        ArrayTable,
        ArrayItems,
        Space,
      }}
    >
      <SchemaComponent schema={schema} />
      <FormButtonGroup.FormItem>
        <Submit long size="large" onSubmit={(values) => console.log(values)}>
          提交
        </Submit>
      </FormButtonGroup.FormItem>
    </SchemaComponentProvider>
  </FormLayout>
);
```
