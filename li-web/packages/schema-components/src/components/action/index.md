# Action

# Action.Drawer

```tsx
import React from "react";
import { ISchema, observer, useForm } from "@formily/react";
import {
  FormItem,
  Input,
  Action,
  Form,
  SchemaComponent,
  SchemaComponentProvider,
  useActionContext,
  useCloseAction,
} from "schema-components";
import "@arco-design/web-react/dist/css/arco.css";

const schema: ISchema = {
  type: "object",
  properties: {
    action1: {
      "x-component": "Action",
      "x-component-props": {
        type: "primary",
      },
      type: "void",
      title: "Open",
      properties: {
        drawer1: {
          "x-component": "Action.Drawer",
          type: "void",
          title: "Drawer Title",
          properties: {
            hello1: {
              "x-content": "Hello",
              title: "T1",
            },
            input: {
              type: "string",
              title: "Input",
              "x-component": "Input",
            },
            footer1: {
              "x-component": "Action.Drawer.Footer",
              type: "void",
              properties: {
                close1: {
                  title: "Close",
                  "x-component": "Action",
                  "x-component-props": {
                    useAction: "{{ useCloseAction }}",
                  },
                },
              },
            },
          },
        },
      },
    },
  },
};

export default () => {
  return (
    <SchemaComponentProvider
      scope={{ useCloseAction }}
      components={{ Form, Action, Input, FormItem }}
    >
      <SchemaComponent schema={schema} />
    </SchemaComponentProvider>
  );
};
```
