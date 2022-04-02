import React from "react";
import {
  ButtonProps,
  PopconfirmProps,
  TableProps,
} from "@arco-design/web-react";
import { ColumnProps } from "@arco-design/web-react/es/Table";
import { FormProps } from "../form";
import { ActionFormDrawerProps, ActionFormModalProps } from "../action/types";
import { ReloadData } from "./context";

export type ComposedList = React.FC<{
  forInit: string;
  forInitVariables?: Record<string, any>;
}> & {
  Filter?: React.FC<FormProps>;
  Action?: ComposedListAction;
  Table?: ComposedListTable;
};

export type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

export type ComposedListAction = React.FC & {
  RowSelection?: React.FC<
    ButtonProps & {
      confirmProps: PopconfirmProps;
      forSubmit?: string;
      afterReload?: boolean;
    }
  >;
  Reload?: React.FC<
    ButtonProps & {
      data?: ReloadData;
    }
  >;
  RecordFormDrawer?: React.FC<ActionFormDrawerProps>;
  RecordFormModal?: React.FC<ActionFormModalProps>;
  RecordDelete?: React.FC<
    ButtonProps & {
      confirmProps?: PopconfirmProps;
      forSubmit?: string;
    }
  >;
};
