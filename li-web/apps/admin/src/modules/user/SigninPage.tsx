import { ISchema, useForm } from "@formily/react";
import { uid } from "@formily/shared";
import { useHistory } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useAPIClient } from "../api-client";
import { useCurrentDocumentTitle } from "../document-title";

const schema: ISchema = {
  type: "object",
  name: uid(),
  "x-component": "Form",
  properties: {
    email: {
      type: "string",
      required: true,
      "x-component": "Input",
      "x-validator": "email",
      "x-decorator": "FormItem",
      "x-component-props": { placeholder: '{{t("Email")}}', style: {} },
    },
    password: {
      type: "string",
      "x-component": "Password",
      required: true,
      "x-decorator": "FormItem",
      "x-component-props": { placeholder: '{{t("Password")}}', style: {} },
    },
    actions: {
      type: "void",
      "x-component": "div",
      properties: {
        submit: {
          title: '{{t("Sign in")}}',
          type: "void",
          "x-component": "Action",
          "x-component-props": {
            block: true,
            type: "primary",
            useAction: "{{ useSignin }}",
            style: { width: "100%" },
          },
        },
      },
    },
    link: {
      type: "void",
      "x-component": "div",
      properties: {
        link: {
          title: '{{t("Create an account")}}',
          type: "void",
          "x-component": "Link",
          "x-content": '{{t("Create an account")}}',
          "x-component-props": { to: "/signup" },
        },
      },
    },
  },
};

const useSignin = () => {
  const history = useHistory();
  const form = useForm();
  const api = useAPIClient();
  return {
    async run() {
      await form.submit();
      const response = await api.resource("users").signin({
        values: form.values,
      });
      if (response?.data?.data?.token) {
        api.setBearerToken(response?.data?.data?.token);
        history.push("/admin");
      }
    },
  };
};

export const SigninPage = () => {
  useCurrentDocumentTitle("Signin");
  return (
    <div>
      <SchemaComponent scope={{ useSignin }} schema={schema} />
    </div>
  );
};
