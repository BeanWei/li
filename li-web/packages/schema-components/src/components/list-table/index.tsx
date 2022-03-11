import { Fragment, useMemo, useState } from "react";
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

type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

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
        return buf.concat([s]);
      }
    }, [])
    .map((s: Schema) => {
      return {
        title: s["x-component-props"]["title"] || s.title,
        dataIndex: s.name,
        key: s.name,
        render: (v, record) => {
          const index = field.value?.indexOf(record);
          return (
            <RecordIndexProvider index={index}>
              <RecordProvider record={record}>
                <RecursionField schema={s} name={index} onlyRenderProperties />
              </RecordProvider>
            </RecordIndexProvider>
          );
        },
      } as ColumnProps<any>;
    });
  return columns;
};

const BaseTable: React.FC<any> = observer((props) => {
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
    const field = useField<ArrayField>();
    const schema = useFieldSchema();
    const form = useMemo(() => createForm(), []);
    const f = useAttach(
      form.createArrayField({ ...field.props, basePath: "" })
    );
    form.setValues({
      [schema.name as string]: [
        { id: 1, name: "阿璃0", ok: true },
        { id: 2, name: "阿璃00", ok: false },
      ],
    });

    const [selectedRowKeys, setSelectedRowKeys] = useState([]);

    return (
      <FormContext.Provider value={form}>
        <FieldContext.Provider value={f}>
          <BaseTable {...props} />
        </FieldContext.Provider>
      </FormContext.Provider>
    );
  }
);

ListTable.Column = () => {
  return <Fragment />;
};

export default ListTable;
