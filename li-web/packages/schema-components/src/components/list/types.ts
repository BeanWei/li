import React from "react";
import {
  ButtonProps,
  CardProps,
  PopconfirmProps,
  SelectProps,
  TableProps,
} from "@arco-design/web-react";
import { ColumnProps } from "@arco-design/web-react/es/Table";
import { FormProps } from "../form";
import { ActionFormDrawerProps, ActionFormModalProps } from "../action/types";

export type ComposedList = React.FC<{
  forInit: string;
  forInitVariables?: Record<string, any>;
  cardProps?: CardProps;
}> & {
  Filter?: React.FC<FormProps>;
  Action?: ComposedListAction;
  Table?: ComposedListTable;
};

export type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

export type ComposedListAction = React.FC & {
  FilterGroup?: React.FC<ButtonProps>;
  FilterSelect?: React.FC<SelectProps>;
  RowSelection?: React.FC<
    ButtonProps & {
      confirmProps: PopconfirmProps;
      forSubmit?: string;
      afterReload?: boolean;
    }
  >;
  Refresh?: React.FC<ButtonProps>;
  RecordEditDrawer?: React.FC<ActionFormDrawerProps>;
  RecordEditModal?: React.FC<ActionFormModalProps>;
  RecordDelete?: React.FC<
    ButtonProps & {
      confirmProps?: PopconfirmProps;
      forSubmit?: string;
    }
  >;
};
