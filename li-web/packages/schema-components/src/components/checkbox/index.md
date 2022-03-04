# Checkbox

## Examples

### 勾选

```tsx
import React from "react";
import {
  Checkbox,
  FormItem,
  SchemaComponent,
  SchemaComponentProvider,
} from "@li/schema-components";

const schema = {
  type: "object",
  properties: {
    input: {
      type: "boolean",
      title: `编辑模式`,
      "x-decorator": "FormItem",
      "x-component": "Checkbox",
      "x-content": "编辑模式",
      "x-reactions": {
        target: "read",
        fulfill: {
          state: {
            value: "{{$self.value}}",
          },
        },
      },
    },
    read: {
      type: "boolean",
      title: `阅读模式`,
      "x-read-pretty": true,
      "x-decorator": "FormItem",
      "x-component": "Checkbox",
    },
  },
};

export default () => {
  return (
    <SchemaComponentProvider components={{ Checkbox, FormItem }}>
      <SchemaComponent schema={schema} />
    </SchemaComponentProvider>
  );
};
```
