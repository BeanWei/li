import { Fragment, useContext, useEffect, useMemo } from "react";
import { Table } from "@arco-design/web-react";
import { ColumnProps, TableProps } from "@arco-design/web-react/es/Table";
import { ArrayField, createForm } from "@formily/core";
import {
  FieldContext,
  FormContext,
  RecursionField,
  Schema,
  useField,
  useFieldSchema,
} from "@formily/react";
import { observer } from "@formily/reactive-react";
import { RecordIndexProvider, RecordProvider } from "../../core";
import { useAttach } from "../../hooks";
import { ComposedListTable } from "./types";
import { ListContext } from "./context";

const isColumnComponent = (schema: Schema) => {
  return schema["x-component"]?.endsWith(".Column") > -1;
};

const useTableColumns = () => {
  const field = useField<ArrayField>();
  const schema = useFieldSchema();
  // @ts-ignore
  const columns = schema
    .reduceProperties((buf: any, s) => {
      if (isColumnComponent(s)) {
        return buf.concat(
          s.reduceProperties((buf_: any, s_) => {
            return buf_.concat([
              {
                colProps: s["x-component-props"],
                title: s?.["x-component-props"]?.["title"] || s.title,
                key: s_.name,
                schema: s,
              },
            ]);
          }, [])
        );
      }
    }, [])
    .map((col: any) => {
      return {
        ...col.colProps,
        title: col.title,
        dataIndex: col.key,
        key: col.key,
        render: (v, record) => {
          const index = field.value?.indexOf(record);
          return (
            <RecordIndexProvider index={index}>
              <RecordProvider record={record}>
                <RecursionField
                  schema={col.schema}
                  name={index}
                  onlyRenderProperties
                />
              </RecordProvider>
            </RecordIndexProvider>
          );
        },
      } as ColumnProps<any>;
    });
  return columns;
};

const BaseTable: React.FC<TableProps> = observer((props) => {
  const field = useField<ArrayField>();
  const columns = useTableColumns();
  return (
    <Table
      rowKey="id"
      {...props}
      columns={columns}
      data={field.value?.slice()}
    />
  );
});

export const ListTable: ComposedListTable = observer(
  (props: TableProps<any>) => {
    const ctx = useContext(ListContext);
    const field = useField<ArrayField>();
    const schema = useFieldSchema();
    const form = useMemo(() => createForm(), []);
    const f = useAttach(
      form.createArrayField({ ...field.props, basePath: "" })
    );
    useEffect(() => {
      form.setInitialValues({
        [schema.name as string]: ctx.tableProps?.data,
      });
    }, [ctx.tableProps?.data]);

    return (
      <FormContext.Provider value={form}>
        <FieldContext.Provider value={f}>
          <BaseTable
            {...ctx.tableProps}
            {...props}
            onChange={ctx.tableProps?.onChange}
            rowSelection={
              props.rowSelection
                ? {
                    ...props.rowSelection,
                    onChange: ctx.setSelectedRowKeys,
                  }
                : undefined
            }
          />
        </FieldContext.Provider>
      </FormContext.Provider>
    );
  }
);

ListTable.Column = () => {
  return <Fragment />;
};

export default ListTable;
