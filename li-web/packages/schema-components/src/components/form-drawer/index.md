# FormDrawer

## Example 1.

```tsx
import React from "react";
import {
  Form,
  Input,
  FormItem,
  FormDrawer,
  FormLayout,
  SchemaComponentProvider,
  SchemaComponent,
} from "schema-components";
import { Button } from "@arco-design/web-react";

const schema = {
  type: "object",
  properties: {
    aaa: {
      type: "string",
      title: "输入框1",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
    },
    bbb: {
      type: "string",
      title: "输入框2",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
    },
    ccc: {
      type: "string",
      title: "输入框3",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
    },
    ddd: {
      type: "string",
      title: "输入框4",
      required: true,
      "x-decorator": "FormItem",
      "x-component": "Input",
    },
  },
};

export default () => (
  <Button
    onClick={() => {
      FormDrawer("弹窗表单", (resolve) => (
        <FormLayout labelCol={6} wrapperCol={10}>
          <SchemaComponentProvider components={{ Form, Input, FormItem }}>
            <SchemaComponent schema={schema} />
          </SchemaComponentProvider>
        </FormLayout>
      ))
        .open({
          initialValues: {
            aaa: "123",
          },
        })
        .then(console.log);
    }}
  >
    点我打开表单
  </Button>
);
```
