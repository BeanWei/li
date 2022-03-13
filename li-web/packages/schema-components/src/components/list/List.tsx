import { useFieldSchema } from "@formily/react";
import { observer } from "@formily/reactive-react";
import { useRequest } from "pro-utils";
import { useCallback, useState } from "react";
import { ListContext } from "./context";
import { ListAction } from "./List.Action";
import ListTable from "./List.Table";
import { ComposedList } from "./types";

export const List: ComposedList = observer((props) => {
  const schema = useFieldSchema();
  const operation = schema["x-operation"];

  const result = useRequest(operation, {
    current: 1,
    pageSize: 10,
  });

  const total = result.data?.total || 0;
  const { current = 1, pageSize = 10 } = result.params[0] || {};

  const onPageChange = (p: number, c: number) => {
    let toCurrent = c <= 0 ? 1 : c;
    const toPageSize = p <= 0 ? 1 : p;
    const tempTotalPage = Math.ceil(total / toPageSize);
    if (toCurrent > tempTotalPage) {
      toCurrent = Math.max(1, tempTotalPage);
    }
    const [oldPaginationParams = {}, ...restParams] = result.params || [];
    result.run({
      ...oldPaginationParams,
      ...restParams,
      current: toCurrent,
      pageSize: toPageSize,
    });
  };

  const changePageSize = (p: number) => {
    onPageChange(current, p);
  };

  const onTableChange = (pagination: any, sorter: any, filters: any) => {
    const [oldPaginationParams, ...restParams] = result.params || [];
    result.run({
      ...oldPaginationParams,
      ...restParams,
      current: pagination.current,
      pageSize: pagination.pageSize,
      filters,
      sorter,
    });
  };

  const [selectedRowKeys, setSelectedRowKeys] = useState<(string | number)[]>(
    []
  );

  return (
    <ListContext.Provider
      value={{
        result,
        tableProps: {
          data: result.data?.data || [{ id: 1, name: "Bean", ok: true }],
          loading: result.loading,
          onChange: useCallback(onTableChange, []),
          pagination: total
            ? {
                current,
                pageSize,
                total,
                onChange: useCallback(onPageChange, []),
                onPageSizeChange: useCallback(changePageSize, []),
              }
            : false,
        },
        selectedRowKeys,
        setSelectedRowKeys: useCallback(setSelectedRowKeys, []),
      }}
    >
      {props.children}
    </ListContext.Provider>
  );
});

List.Action = ListAction;
List.Table = ListTable;

export default List;
