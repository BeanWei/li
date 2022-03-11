# ListTable

## Example 1.

```tsx
import React from "react";
import {
  Input,
  SchemaComponent,
  SchemaComponentProvider,
  ListTable,
} from "schema-components";
import { Button } from "@arco-design/web-react";
import "@arco-design/web-react/dist/css/arco.css";

const schema: ISchema = {
  type: "object",
  properties: {
    input: {
      type: "array",
      title: `编辑模式`,
      default: [
        { id: 1, name: "Name1" },
        { id: 2, name: "Name2" },
        { id: 3, name: "Name3" },
      ],
      "x-component": "ListTable",
      "x-component-props": {
        rowKey: "id",
        rowSelection: {
          type: "checkbox",
        },
        data: [{ id: 1, name: "阿璃" }],
      },
      properties: {
        column1: {
          type: "void",
          title: "Name",
          "x-component": "ListTable.Column",
          properties: {
            name: {
              type: "string",
              "x-component": "Input",
              "x-read-pretty": true,
            },
          },
        },
      },
    },
  },
};

export default () => (
  <SchemaComponentProvider components={{ ListTable, Input }}>
    <SchemaComponent schema={schema} />
  </SchemaComponentProvider>
);
```
