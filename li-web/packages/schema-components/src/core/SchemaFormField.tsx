import { useMemo } from "react";
import { createForm, onFieldValueChange } from "@formily/core";
import { FormProvider, Schema, Stringify } from "@formily/react";
import { UiSchemaComponentProvider } from "./UiSchemaComponentProvider";
import { SchemaComponent } from "./SchemaComponent";

export const SchemaFormField: React.FC<{
  name: string;
  schema: Stringify<any>;
  value?: any;
  onChange?: (value: any) => void;
}> = (props) => {
  const form = useMemo(
    () =>
      createForm({
        initialValues: {
          [props.name]: props.value,
        },
        effects() {
          onFieldValueChange(props.name, (field) => {
            props.onChange?.(field.value);
          });
        },
      }),
    []
  );
  const schema = new Schema({
    type: "object",
    properties: {
      [props.name]: props.schema,
    },
  });
  schema.reduceProperties((b, s) => {
    s["x-decorator"] = undefined;
    s["x-decorator-props"] = undefined;
    s["x-read-pretty"] = false;
  });

  return (
    <FormProvider form={form}>
      <UiSchemaComponentProvider>
        <SchemaComponent schema={schema} />
      </UiSchemaComponentProvider>
    </FormProvider>
  );
};
