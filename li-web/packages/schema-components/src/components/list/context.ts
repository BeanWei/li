import { TableProps } from "@arco-design/web-react";
import { Result } from "pro-utils";
import { createContext } from "react";

export const ListContext = createContext<{
  result?: Result<any, any>;
  tableProps?: TableProps & {
    onSearch?: (values: Record<string, any>) => void;
  };
  selectedRowKeys?: (string | number)[];
  setSelectedRowKeys?: (keys: (string | number)[]) => void;
}>({});
