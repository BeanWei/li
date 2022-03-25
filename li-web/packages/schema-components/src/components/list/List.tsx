import { observer } from "@formily/reactive-react";
import { useRequest } from "pro-utils";
import { useCallback, useState } from "react";
import { ListContext } from "./context";
import { ListAction } from "./List.Action";
import ListTable from "./List.Table";
import { ComposedList } from "./types";

export const List: ComposedList = observer((props) => {
  const result = useRequest(props.forInit, {
    ...props.forInitVariables,
    page: 1,
    limit: 10,
  });

  const total = result.data?.total || 0;
  const { page: current = 1, limit: pageSize = 10 } = result.params[0] || {};

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
      page: toCurrent,
      limit: toPageSize,
    });
  };

  const changePageSize = (p: number) => {
    onPageChange(current, p);
  };

  const onTableChange = (pagination: any, sorter: any, filters: any) => {
    result.run({
      page: pagination.current,
      limit: pagination.pageSize,
      // filter: filters,
      // sort: sorter,
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
          data: result.data?.list || [],
          loading: result.loading,
          onChange: useCallback(onTableChange, []),
          pagination: total
            ? {
                current,
                pageSize,
                total,
                onChange: onPageChange,
                onPageSizeChange: changePageSize,
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
