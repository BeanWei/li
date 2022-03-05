import { ISchema, useForm } from "@formily/react";
import { uid } from "@formily/shared";
import { Message } from "@arco-design/web-react";
import { useHistory } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useCurrentDocumentTitle } from "../document-title";
import { request } from "pro-utils";

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
      required: true,
      "x-component": "Password",
      "x-decorator": "FormItem",
      "x-component-props": {
        placeholder: '{{t("Password")}}',
        checkStrength: true,
        style: {},
      },
      "x-reactions": [
        {
          dependencies: [".confirm_password"],
          fulfill: {
            state: {
              selfErrors:
                '{{$deps[0] && $self.value && $self.value !== $deps[0] ? t("Password mismatch") : ""}}',
            },
          },
        },
      ],
    },
    confirm_password: {
      type: "string",
      required: true,
      "x-component": "Password",
      "x-decorator": "FormItem",
      "x-component-props": {
        placeholder: '{{t("Confirm password")}}',
        checkStrength: true,
        style: {},
      },
      "x-reactions": [
        {
          dependencies: [".password"],
          fulfill: {
            state: {
              selfErrors:
                '{{$deps[0] && $self.value && $self.value !== $deps[0] ? t("Password mismatch") : ""}}',
            },
          },
        },
      ],
    },
    actions: {
      type: "void",
      "x-component": "div",
      properties: {
        submit: {
          title: '{{t("Sign up")}}',
          type: "void",
          "x-component": "Action",
          "x-component-props": {
            block: true,
            type: "primary",
            useAction: "{{ useSignup }}",
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
          type: "void",
          "x-component": "Link",
          "x-component-props": { to: "/signin" },
          "x-content": '{{t("Log in with an existing account")}}',
        },
      },
    },
  },
};

const useSignup = () => {
  const history = useHistory();
  const form = useForm();
  return {
    async run() {
      await form.submit();
      await request("userSignup", form.values);
      Message.success("注册成功，即将跳转登录页");
      setTimeout(() => {
        history.push("/signin");
      }, 2000);
    },
  };
};

export const SignupPage = () => {
  useCurrentDocumentTitle("Signup");
  return <SchemaComponent schema={schema} scope={{ useSignup }} />;
};
