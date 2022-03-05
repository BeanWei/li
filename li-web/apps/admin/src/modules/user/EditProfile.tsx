import { useContext, useState } from "react";
import { ISchema, useForm } from "@formily/react";
import { uid } from "@formily/shared";
import { Menu } from "@arco-design/web-react";
import { useTranslation } from "react-i18next";
import {
  ActionContext,
  SchemaComponent,
  useActionContext,
} from "schema-components";
import { useCurrentUserContext } from "./CurrentUserProvider";
import { DropdownVisibleContext } from "./CurrentUser";
import { useRequest } from "pro-utils";

const useCloseAction = () => {
  const { setVisible } = useActionContext();
  const form = useForm();
  return {
    async run() {
      setVisible(false);
      form.submit((values) => {
        console.log(values);
      });
    },
  };
};

const useCurrentUserValues = (options: any) => {
  const ctx = useCurrentUserContext();
  return useRequest(() => Promise.resolve(ctx.data), options);
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
      "x-decorator": "Form",
      "x-decorator-props": {
        useValues: "{{ useCurrentUserValues }}",
      },
      "x-component": "Action.Drawer",
      type: "void",
      title: "个人资料",
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
  const { t } = useTranslation();
  const ctx = useContext(DropdownVisibleContext);
  return (
    <ActionContext.Provider value={{ visible, setVisible }}>
      <Menu.Item
        key={"ProfileAction"}
        onClick={() => {
          setVisible(true);
          ctx.setVisible(false);
        }}
      >
        {t("Edit profile")}
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
