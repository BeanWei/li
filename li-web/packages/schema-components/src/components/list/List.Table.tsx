import { Fragment, useContext, useEffect, useMemo } from "react";
import { Table } from "@arco-design/web-react";
import { TableProps } from "@arco-design/web-react/es/Table";
import { ArrayField, createForm } from "@formily/core";
import {
  FieldContext,
  FormContext,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { observer } from "@formily/reactive-react";
import { RecordIndexProvider, RecordProvider } from "../../core";
import { useAttach } from "../../hooks";
import { isColumnComponent, useArrayTableSources } from "../array-table";
import { ComposedListTable } from "./types";
import { ListContext } from "./context";

const useListTableColumns = (dataSource: any[]): TableProps<any>["columns"] => {
  const source = useArrayTableSources();
  return source.reduce((buf, { name, columnProps, schema, display }, key) => {
    if (display && display !== "visible") return buf;
    if (!isColumnComponent(schema)) return buf;
    schema.reduceProperties((buf, s) => {
      s.title = "";
      s["x-read-pretty"] = true;
    });
    return buf.concat({
      ...columnProps,
      // @ts-ignore
      key,
      dataIndex: name,
      render: (value: any, record: any) => {
        const index = dataSource.indexOf(record);
        return (
          <RecordIndexProvider index={index}>
            <RecordProvider record={record}>
              <RecursionField
                schema={schema}
                name={index}
                onlyRenderProperties
              />
            </RecordProvider>
          </RecordIndexProvider>
        );
      },
    });
  }, []);
};

const BaseTable: React.FC<TableProps> = observer((props) => {
  const field = useField<ArrayField>();
  const columns = useListTableColumns(field.value?.slice());
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
