import {
  ButtonProps,
  DrawerProps,
  SelectProps,
  TableProps,
} from "@arco-design/web-react";
import { InputSearchProps } from "@arco-design/web-react/es/Input";
import { ColumnProps } from "@arco-design/web-react/es/Table";
import React from "react";
import {
  ActionProps,
  ComposedActionDrawer,
  ComposedActionModal,
} from "../action/types";
import { FormProps } from "../form";

export type ComposedList = React.FC & {
  Filter?: React.FC<FormProps>;
  Action?: ComposedListAction;
  Table?: ComposedListTable;
};

export type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

export type ComposedListAction = React.FC & {
  FilterGroup?: React.FC<ActionProps>;
  FilterSelect?: React.FC<SelectProps>;
  FilterItem?: React.FC;
  RowSelection?: React.FC<ActionProps>;
  Search?: React.FC<InputSearchProps>;
  Refresh?: React.FC<ButtonProps>;
  BulkDelete?: React.FC<ActionProps>;
};
