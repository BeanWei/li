import { TableProps } from "@arco-design/web-react";
import { BaseResult } from "pro-utils";
import { createContext } from "react";

export const ListContext = createContext<{
  result?: BaseResult<any, any>;
  tableProps?: TableProps;
  selectedRowKeys?: (string | number)[];
  setSelectedRowKeys?: (keys: (string | number)[]) => void;
}>({});
