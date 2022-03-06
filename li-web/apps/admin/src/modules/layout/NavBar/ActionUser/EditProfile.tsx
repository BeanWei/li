import { ISchema } from "@formily/react";
import {
  Form,
  FormDrawer,
  FormItem,
  FormLayout,
  Input,
  SchemaComponent,
  SchemaComponentProvider,
} from "schema-components";
import { useAdminLayoutContext } from "../../AdminLayoutProvider";

const schema: ISchema = {
  type: "object",
  properties: {
    nickname: {
      type: "string",
      title: "{{t('Nickname')}}",
      "x-decorator": "FormItem",
      "x-component": "Input",
      required: true,
    },
  },
};

export const openEditProfileDrawer = (initialValues: Record<string, any>) => {
  FormDrawer("个人资料", () => {
    return (
      <FormLayout>
        <SchemaComponentProvider components={{ Form, Input, FormItem }}>
          <SchemaComponent schema={schema} />
        </SchemaComponentProvider>
      </FormLayout>
    );
  }).open({
    initialValues,
  });
};
