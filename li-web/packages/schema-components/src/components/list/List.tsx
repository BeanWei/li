import { TableProps } from "@arco-design/web-react";
import { observer } from "@formily/reactive-react";
import { isValid } from "@formily/shared";
import { useRequest } from "pro-utils";
import { useCallback, useState } from "react";
import { ListContext, ReloadData } from "./context";
import { ListAction } from "./List.Action";
import ListTable from "./List.Table";
import { ComposedList } from "./types";

export const List: ComposedList = observer((props_) => {
  const useProps = props_.useProps?.() || {};
  const props = { ...props_, ...useProps };

  const result = useRequest(props.forInit, {
    ...props.forInitVariables,
    page: 1,
    limit: 10,
  });

  const total = result.data?.total || 0;
  const { page: current = 1, limit: pageSize = 10 } = result.params[0] || {};

  const reload = (values?: ReloadData) => {
    if (!values) {
      result.refresh();
    } else {
      result.run({
        page: current,
        limit: pageSize,
        ...result.params[0],
        ...values,
      });
    }
  };

  const onPageChange = (p: number, c: number) => {
    let toCurrent = c <= 0 ? 1 : c;
    const toPageSize = p <= 0 ? 1 : p;
    const tempTotalPage = Math.ceil(total / toPageSize);
    if (toCurrent > tempTotalPage) {
      toCurrent = Math.max(1, tempTotalPage);
    }
    reload({
      page: toCurrent,
      limit: toPageSize,
    });
  };

  const changePageSize = (p: number) => {
    onPageChange(current, p);
  };

  const onTableChange: TableProps["onChange"] = (
    pagination,
    sorter,
    filters
  ) => {
    reload({
      page: pagination.current,
      limit: pagination.pageSize,
      filter: filters,
      sorter: {
        [sorter.field as string]:
          sorter.direction === "ascend"
            ? 1
            : sorter.direction === "descend"
            ? -1
            : 0,
      },
    });
  };

  const [selectedKeys, setSelectedKeys] = useState<(string | number)[]>(
    props.selection?.defaultSelectedKeys || []
  );

  return (
    <ListContext.Provider
      value={{
        result,
        reload: useCallback(reload, []),
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
          rowSelection: props.selection
            ? {
                checkAll: isValid(props.selection.enableCheckAll)
                  ? props.selection.enableCheckAll
                  : true,
                type: props.selection.multiple ? "checkbox" : "radio",
                selectedRowKeys: selectedKeys,
                preserveSelectedRowKeys: props.selection.preserveSelectedKeys,
                onChange: (keys, records) => {
                  props.selection?.onChange?.(keys, records);
                  setSelectedKeys(keys);
                },
              }
            : undefined,
          filter: props.filter,
        },
        selectedKeys,
      }}
    >
      {props.children}
    </ListContext.Provider>
  );
});

List.Action = ListAction;
List.Table = ListTable;

export default List;
