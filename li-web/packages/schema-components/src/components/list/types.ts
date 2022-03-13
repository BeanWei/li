import { ButtonProps, CardProps, TableProps } from "@arco-design/web-react";
import { ColumnProps } from "@arco-design/web-react/es/Table";
import React from "react";

export type ComposedList = React.FC & {
  Table?: ComposedListTable;
  Action?: ComposedListAction;
};

export type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

export type ComposedListAction = React.FC & {
  BulkDelete?: React.FC<
    {
      title?: string;
    } & ButtonProps
  >;
};
