import { TableProps } from "@arco-design/web-react";
import { Result } from "pro-utils";
import { createContext } from "react";

export type ReloadData = {
  page?: number;
  limit?: number;
  sorter?: Record<string, number>;
  filter?: Record<string, any>;
};

export type ListContextProps = {
  result?: Result<any, any>;
  reload?: (values?: ReloadData) => void;
  tableProps?: TableProps & {
    onSearch?: (values: Record<string, any>) => void;
  };
  selectedKeys?: (string | number)[];
};

export const ListContext = createContext<ListContextProps>({});
