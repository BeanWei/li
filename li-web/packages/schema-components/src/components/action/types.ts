import React from "react";
import { ButtonProps } from "@arco-design/web-react";
import { IFormLayoutProps } from "../form-layout";
import { IModalProps } from "../form-modal";
import { IDrawerProps } from "../form-drawer";

export type ComposedAction = React.FC & {
  FormDrawer?: ActionFormDrawerProps;
  FormModal?: ActionFormModalProps;
};

export type ActionFormDrawerProps = React.FC<{
  initialValues?: Record<string, any>;
  forOpen?: string;
  forOpenVariables?: Record<string, any>;
  forSubmit?: string;
  buttonProps?: ButtonProps;
  isMenuItem?: boolean;
  drawerProps?: IDrawerProps;
  layoutProps?: IFormLayoutProps;
}>;

export type ActionFormModalProps = React.FC<{
  initialValues?: Record<string, any>;
  forOpen?: string;
  forOpenVariables?: Record<string, any>;
  forSubmit?: string;
  buttonProps?: ButtonProps;
  isMenuItem?: boolean;
  modalProps?: IModalProps;
  layoutProps?: IFormLayoutProps;
}>;
