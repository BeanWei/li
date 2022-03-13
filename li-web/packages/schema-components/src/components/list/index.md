# List

## ListTable

```tsx
import React from "react";
import {
  Input,
  ActionBar,
  SchemaComponent,
  SchemaComponentProvider,
  List,
  Checkbox,
  Form,
  FormItem,
} from "schema-components";
import { Button } from "@arco-design/web-react";
import "@arco-design/web-react/dist/css/arco.css";

const schema: ISchema = {
  type: "object",
  properties: {
    list: {
      type: "void",
      "x-component": "List",
      properties: {
        actions: {
          type: "void",
          "x-component": "List.Action",
          properties: {
            filtergroup: {
              type: "void",
              "x-component": "List.Action.FilterGroup",
              "x-component-props": {
                position: "left",
              },
              properties: {
                input1: {
                  type: "string",
                  title: "Name",
                  "x-decorator": "FormItem",
                  "x-component": "Input",
                },
              },
            },
            search: {
              type: "void",
              "x-component": "List.Action.Search",
              "x-component-props": {
                position: "left",
              },
            },
            filter2: {
              type: "void",
              "x-component": "List.Action.FilterSelect",
            },
            delete: {
              type: "void",
              title: "Delete",
              "x-component": "List.Action.BulkDelete",
            },
            refresh: {
              type: "void",
              "x-component": "List.Action.Refresh",
            },
          },
        },
        table: {
          type: "array",
          "x-component": "List.Table",
          "x-component-props": {
            rowKey: "id",
            rowSelection: {
              type: "checkbox",
            },
          },
          properties: {
            column1: {
              type: "void",
              "x-component": "List.Table.Column",
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
              "x-component": "List.Table.Column",
              "x-component-props": {
                title: "OK",
                dataIndex: "ok",
                sorter: true,
                filters: [
                  {
                    text: "True",
                    value: true,
                  },
                  {
                    text: "Flase",
                    value: false,
                  },
                ],
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
    },
  },
};

export default () => (
  <SchemaComponentProvider
    components={{ ActionBar, List, Input, Checkbox, Form, FormItem }}
  >
    <SchemaComponent schema={schema} />
  </SchemaComponentProvider>
);
```
