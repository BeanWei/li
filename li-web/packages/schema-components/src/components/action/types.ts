import React from "react";
import { ButtonProps } from "@arco-design/web-react";
import { IFormLayoutProps } from "../form-layout";
import { IModalProps } from "../form-modal";
import { IDrawerProps } from "../form-drawer";

export type ComposedAction = React.FC & {
  FormDrawer?: React.FC<ActionFormDrawerProps>;
  FormModal?: React.FC<ActionFormModalProps>;
};

export type ActionFormDrawerProps = ButtonProps & {
  initialValues?: Record<string, any>;
  forInit?: string;
  forInitVariables?: Record<string, any>;
  forSubmitSuccess?: (payload: any) => void;
  isMenuItem?: boolean;
  drawerProps?: IDrawerProps;
  layoutProps?: IFormLayoutProps;
};

export type ActionFormModalProps = ButtonProps & {
  initialValues?: Record<string, any>;
  forInit?: string;
  forInitVariables?: Record<string, any>;
  forSubmitSuccess?: (payload: any) => void;
  isMenuItem?: boolean;
  modalProps?: IModalProps;
  layoutProps?: IFormLayoutProps;
};
