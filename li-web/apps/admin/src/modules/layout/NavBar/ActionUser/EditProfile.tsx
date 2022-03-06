import { useState } from "react";
import { ISchema, useForm } from "@formily/react";
import { IconSettings } from "@arco-design/web-react/icon";
import { uid } from "@formily/shared";
import { Menu } from "@arco-design/web-react";
import {
  ActionContext,
  SchemaComponent,
  useActionContext,
  useCloseAction,
} from "schema-components";
import { useRequest } from "pro-utils";
import { useAdminLayoutContext } from "../../AdminLayoutProvider";
import styles from "../style/index.module.less";

const useCurrentUserValues = (options: any) => {
  const ctx = useAdminLayoutContext();
  return useRequest(() => Promise.resolve(ctx.profile), options);
};

const useSaveCurrentUserValues = () => {
  const { setVisible } = useActionContext();
  const form = useForm();
  return {
    async run() {
      form.submit((values) => {
        setVisible(false);
        console.log(values);
      });
    },
  };
};

const schema: ISchema = {
  type: "object",
  properties: {
    [uid()]: {
      type: "void",
      title: "个人资料",
      "x-decorator": "Form",
      "x-decorator-props": {
        useValues: "{{ useCurrentUserValues }}",
      },
      "x-component": "Action.Drawer",
      properties: {
        nickname: {
          type: "string",
          title: "{{t('Nickname')}}",
          "x-decorator": "FormItem",
          "x-component": "Input",
          required: true,
        },
        footer: {
          "x-component": "Action.Drawer.Footer",
          type: "void",
          properties: {
            cancel: {
              title: "Cancel",
              "x-component": "Action",
              "x-component-props": {
                useAction: "{{ useCloseAction }}",
              },
            },
            submit: {
              title: "Submit",
              "x-component": "Action",
              "x-component-props": {
                type: "primary",
                useAction: "{{ useSaveCurrentUserValues }}",
              },
            },
          },
        },
      },
    },
  },
};

export const EditProfile = () => {
  const [visible, setVisible] = useState(false);
  return (
    <ActionContext.Provider value={{ visible, setVisible }}>
      <Menu.Item
        key="editprofile"
        onClick={() => {
          setVisible(true);
        }}
      >
        <IconSettings className={styles["dropdown-icon"]} />
        用户设置
      </Menu.Item>
      <SchemaComponent
        scope={{
          useCurrentUserValues,
          useCloseAction,
          useSaveCurrentUserValues,
        }}
        schema={schema}
      />
    </ActionContext.Provider>
  );
};
