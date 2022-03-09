import React, { Fragment, useState, useRef, useEffect } from "react";
import { Table, Pagination, PaginationProps } from "@arco-design/web-react";
import { ColumnProps, TableProps } from "@arco-design/web-react/es/Table";
import { SortableContainer, SortableElement } from "react-sortable-hoc";
import { GeneralField, FieldDisplayTypes, ArrayField } from "@formily/core";
import {
  useField,
  observer,
  useFieldSchema,
  RecursionField,
} from "@formily/react";
import { isArr, isBool } from "@formily/shared";
import { Schema } from "@formily/json-schema";
import { usePrefixCls } from "../__builtins__";
import { ArrayBase, ArrayBaseMixins } from "../array-base";
import "./index.less";

interface ObservableColumnSource {
  field: GeneralField;
  columnProps: ColumnProps<any>;
  schema: Schema;
  display: FieldDisplayTypes;
  name: string;
}
interface IArrayTablePaginationProps extends PaginationProps {
  dataSource?: any[];
  children?: (
    dataSource: any[],
    pagination: React.ReactNode
  ) => React.ReactElement;
}

type ComposedArrayTable = React.FC<TableProps<any>> &
  ArrayBaseMixins & {
    Column?: React.FC<ColumnProps<any>>;
  };

const SortableRow = SortableElement((props: any) => <tr {...props} />);
const SortableBody = SortableContainer((props: any) => <tbody {...props} />);

const isColumnComponent = (schema: Schema) => {
  return schema["x-component"]?.indexOf("Column") > -1;
};

const isOperationsComponent = (schema: Schema) => {
  return schema["x-component"]?.indexOf("Operations") > -1;
};

const isAdditionComponent = (schema: Schema) => {
  return schema["x-component"]?.indexOf("Addition") > -1;
};

const useArrayTableSources = () => {
  const arrayField = useField();
  const schema = useFieldSchema();
  const parseSources = (schema: Schema): ObservableColumnSource[] => {
    if (
      isColumnComponent(schema) ||
      isOperationsComponent(schema) ||
      isAdditionComponent(schema)
    ) {
      if (!schema["x-component-props"]?.["dataIndex"] && !schema["name"])
        return [];
      const name = schema["x-component-props"]?.["dataIndex"] || schema["name"];
      const field = arrayField.query(arrayField.address.concat(name)).take();
      const columnProps =
        // @ts-ignore
        field?.component?.[1] || schema["x-component-props"] || {};
      const display = field?.display || schema["x-display"];
      return [
        {
          name,
          display,
          field,
          schema,
          columnProps,
        },
      ];
    } else if (schema.properties) {
      return schema.reduceProperties((buf, schema) => {
        // @ts-ignore
        return buf.concat(parseSources(schema));
      }, []);
    }
    return [];
  };

  const parseArrayItems = (schema: Schema["items"]) => {
    if (!schema) return [];
    const sources: ObservableColumnSource[] = [];
    const items = isArr(schema) ? schema : [schema];
    return items.reduce((columns, schema) => {
      const item = parseSources(schema);
      if (item) {
        return columns.concat(item);
      }
      return columns;
    }, sources);
  };

  if (!schema) throw new Error("can not found schema object");

  return parseArrayItems(schema.items);
};

const useArrayTableColumns = (
  dataSource: any[],
  sources: ObservableColumnSource[]
): TableProps<any>["columns"] => {
  return sources.reduce((buf, { name, columnProps, schema, display }, key) => {
    if (display !== "visible") return buf;
    if (!isColumnComponent(schema)) return buf;
    return buf.concat({
      ...columnProps,
      // @ts-ignore
      key,
      dataIndex: name,
      render: (value: any, record: any) => {
        const index = dataSource.indexOf(record);
        const children = (
          <ArrayBase.Item index={index} record={record}>
            <RecursionField schema={schema} name={index} onlyRenderProperties />
          </ArrayBase.Item>
        );
        return children;
      },
    });
  }, []);
};

const useAddition = () => {
  const schema = useFieldSchema();
  return schema.reduceProperties((addition, schema, key) => {
    if (isAdditionComponent(schema)) {
      return <RecursionField schema={schema} name={key} />;
    }
    return addition;
  }, null);
};

const ArrayTablePagination: React.FC<IArrayTablePaginationProps> = (props) => {
  const [current, setCurrent] = useState(1);
  const prefixCls = usePrefixCls("formily-array-table");
  const pageSize = props.pageSize || 10;
  const size = props.size || "default";
  const dataSource = props.dataSource || [];
  const startIndex = (current - 1) * pageSize;
  const endIndex = startIndex + pageSize - 1;
  const total = dataSource?.length || 0;
  const totalPage = Math.ceil(total / pageSize);
  const handleChange = (current: number) => {
    setCurrent(current);
  };

  useEffect(() => {
    if (totalPage > 0 && totalPage < current) {
      handleChange(totalPage);
    }
  }, [totalPage, current]);

  const renderPagination = () => {
    if (totalPage <= 1) return;
    return (
      <div className={`${prefixCls}-pagination`}>
        <Pagination
          {...props}
          simple
          pageSize={pageSize}
          current={current}
          total={dataSource.length}
          size={size}
          sizeCanChange={false}
          onChange={handleChange}
        />
      </div>
    );
  };

  return (
    <Fragment>
      {props.children?.(
        dataSource?.slice(startIndex, endIndex + 1),
        renderPagination()
      )}
    </Fragment>
  );
};

const RowComp = (props: any) => {
  return <SortableRow index={props["index"] || 0} {...props} />;
};

export const ArrayTable: ComposedArrayTable = observer(
  (props: TableProps<any>) => {
    const ref = useRef<HTMLDivElement | null>(null);
    const field = useField<ArrayField>();
    const prefixCls = usePrefixCls("formily-array-table");
    const dataSource = Array.isArray(field.value) ? field.value.slice() : [];
    const sources = useArrayTableSources();
    const columns = useArrayTableColumns(dataSource, sources);
    const pagination = isBool(props.pagination) ? {} : props.pagination;
    const addition = useAddition();
    const defaultRowKey = (record: any): string => {
      return dataSource.indexOf(record).toString();
    };
    const addTdStyles = (node: HTMLElement) => {
      const helper = document.body.querySelector(`.${prefixCls}-sort-helper`);
      if (helper) {
        const tds = node.querySelectorAll("td");
        requestAnimationFrame(() => {
          helper.querySelectorAll("td").forEach((td, index) => {
            if (tds[index]) {
              td.style.width = getComputedStyle(tds[index]).width;
            }
          });
        });
      }
    };

    return (
      <ArrayTablePagination {...pagination} dataSource={dataSource}>
        {(dataSource, pager) => (
          <div ref={ref} className={prefixCls}>
            <ArrayBase>
              <Table
                size="small"
                border
                rowKey={defaultRowKey}
                {...props}
                onChange={() => {}}
                pagination={false}
                columns={columns}
                data={dataSource}
                components={{
                  body: {
                    tbody: (props: any) => (
                      <SortableBody
                        useDragHandle
                        lockAxis="y"
                        helperClass={`${prefixCls}-sort-helper`}
                        helperContainer={() => {
                          return ref.current?.querySelector("tbody");
                        }}
                        onSortStart={({ node }: any) => {
                          addTdStyles(node);
                        }}
                        onSortEnd={({ oldIndex, newIndex }: any) => {
                          field.move(oldIndex, newIndex);
                        }}
                        {...props}
                      />
                    ),
                    row: RowComp,
                  },
                }}
              />
              <div style={{ marginTop: 5, marginBottom: 5 }}>{pager}</div>
              {sources.map((column, key) => {
                //专门用来承接对Column的状态管理
                if (!isColumnComponent(column.schema)) return;
                return React.createElement(RecursionField, {
                  name: column.name,
                  schema: column.schema,
                  onlyRenderSelf: true,
                  key,
                });
              })}
              {addition}
            </ArrayBase>
          </div>
        )}
      </ArrayTablePagination>
    );
  }
);

ArrayTable.displayName = "ArrayTable";

ArrayTable.Column = () => {
  return <Fragment />;
};

ArrayBase.mixin(ArrayTable);

export default ArrayTable;
