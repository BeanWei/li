# ListTable

## Example 1.

```tsx
import React from "react";
import {
  Input,
  SchemaComponent,
  SchemaComponentProvider,
  ListTable,
  Checkbox,
} from "schema-components";
import { Button } from "@arco-design/web-react";
import "@arco-design/web-react/dist/css/arco.css";

const schema: ISchema = {
  type: "object",
  properties: {
    table1: {
      type: "array",
      title: `ListTable-1`,
      "x-component": "ListTable",
      "x-component-props": {
        rowKey: "id",
        rowSelection: {
          type: "checkbox",
        },
      },
      properties: {
        column1: {
          type: "void",
          "x-component": "ListTable.Column",
          "x-component-props": {
            title: "Name",
            dataIndex: "name",
          },
          properties: {
            name: {
              type: "string",
              "x-component": "Input",
              "x-read-pretty": true,
            },
          },
        },
        column2: {
          type: "void",
          "x-component": "ListTable.Column",
          "x-component-props": {
            title: "OK",
            dataIndex: "ok",
          },
          properties: {
            ok: {
              type: "boolean",
              "x-component": "Checkbox",
              "x-read-pretty": true,
            },
          },
        },
      },
    },
  },
};

export default () => (
  <SchemaComponentProvider components={{ ListTable, Input, Checkbox }}>
    <SchemaComponent schema={schema} />
  </SchemaComponentProvider>
);
```
