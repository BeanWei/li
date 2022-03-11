import { Fragment, useMemo, useState } from "react";
import { Table } from "@arco-design/web-react";
import { ColumnProps, TableProps } from "@arco-design/web-react/es/Table";
import { ArrayField } from "@formily/core";
import {
  RecursionField,
  Schema,
  useField,
  useFieldSchema,
} from "@formily/react";
import { observer } from "@formily/reactive-react";
import { RecordIndexProvider, RecordProvider } from "../../core";

type ComposedListTable = React.FC<TableProps<any>> & {
  Column?: React.FC<ColumnProps<any>>;
};

const isColumnComponent = (schema: Schema) => {
  return schema["x-component"]?.endsWith(".Column") > -1;
};

export const ListTable: ComposedListTable = observer(
  (props: TableProps<any>) => {
    const field = useField<ArrayField>();
    const schema = useFieldSchema();
    const columns = useMemo(() => {
      return schema
        .reduceProperties((buf: any, s) => {
          if (isColumnComponent(s)) {
            return buf.concat([s]);
          }
        }, [])
        .map((s: Schema) => {
          return {
            title: <RecursionField name={s.name} schema={s} onlyRenderSelf />,
            dataIndex: s.name,
            key: s.name,
            render: (v, record) => {
              const index = field.value?.indexOf(record);
              return (
                <RecordIndexProvider index={index}>
                  <RecordProvider record={record}>
                    <RecursionField
                      schema={s}
                      name={index}
                      onlyRenderProperties
                    />
                  </RecordProvider>
                </RecordIndexProvider>
              );
            },
          } as ColumnProps<any>;
        });
    }, [schema]);
    const [selectedRowKeys, setSelectedRowKeys] = useState([]);

    return <Table {...props} columns={columns} />;
  }
);

ListTable.Column = () => {
  return <Fragment />;
};

export default ListTable;
