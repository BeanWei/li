import { Fragment, useContext, useEffect, useMemo } from "react";
import {
  Button,
  ConfigProvider,
  DatePicker,
  Input,
  Table,
  TimePicker,
} from "@arco-design/web-react";
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
import { isColumnComponent, useArrayTableSources } from "../array-table";
import { ComposedListTable } from "./types";
import { ListContext } from "./context";
import { isObject } from "lodash";
import { IconFilter } from "@arco-design/web-react/icon";

type FilterConfig = Pick<
  ColumnProps,
  "filters" | "filterIcon" | "filterMultiple" | "filterDropdown"
>;
const getLightFilterConfig = (schema: Schema): FilterConfig => {
  const { prefixCls, locale } = useContext(ConfigProvider.ConfigContext);
  const filterConfig: FilterConfig = {};
  if (Array.isArray(schema.enum) && schema.enum.length) {
    filterConfig.filters = schema.enum.map((opt) => {
      if (isObject(opt)) {
        return {
          text: opt?.label || opt?.value,
          value: opt?.value,
        };
      }
      return {
        text: opt,
        value: opt,
      };
    });
    filterConfig.filterMultiple = true;
  } else if (
    schema["x-component"] === "Switch" ||
    schema["x-component"] === "Checkbox"
  ) {
    filterConfig.filters = [
      { text: "Yes", value: true },
      { text: "No", value: false },
    ];
    filterConfig.filterMultiple = false;
  } else {
    filterConfig.filterIcon = <IconFilter />;
    filterConfig.filterDropdown = ({ filterKeys, setFilterKeys, confirm }) => {
      let filterNode = null;
      switch (schema["x-component"]) {
        case "DatePicker":
        case "DatePicker.RangePicker":
          filterNode = (
            <DatePicker.RangePicker
              size="small"
              mode={schema["x-component-props"]?.mode}
              value={filterKeys}
              onChange={(values) => setFilterKeys?.(values)}
            />
          );
          break;
        case "TimePicker":
        case "TimePicker.RangePicker":
          filterNode = (
            <TimePicker.RangePicker
              size="small"
              value={filterKeys}
              onChange={(values) => setFilterKeys?.(values)}
            />
          );
          break;
        case "InputNumber":
        case "Money":
        case "Rate":
        case "Slider":
        // TODO: 数字范围选择器
        default:
          filterNode = (
            <Input.Search
              size="small"
              value={Array.isArray(filterKeys) ? filterKeys[0] : filterKeys}
              onChange={(value) => setFilterKeys?.(value ? [value] : [])}
            />
          );
          break;
      }
      return (
        <div className={`${prefixCls}-table-filters-popup`}>
          <div
            style={{
              padding: "4px 12px",
              fontSize: "14PX",
            }}
          >
            {filterNode}
          </div>
          <div className={`${prefixCls}-table-filters-btn`}>
            <Button
              onClick={() => setFilterKeys?.([])}
              size="mini"
              style={{ marginRight: 8 }}
            >
              {locale?.Table.resetText}
            </Button>
            <Button onClick={() => confirm?.()} type="primary" size="mini">
              {locale?.Table.okText}
            </Button>
          </div>
        </div>
      );
    };
  }
  return filterConfig;
};

const useListTableColumns = (
  dataSource: any[],
  filterRender?: boolean
): TableProps<any>["columns"] => {
  const source = useArrayTableSources();
  return source.reduce((buf, { name, columnProps, schema, display }, key) => {
    if (display && display !== "visible") return buf;
    if (!isColumnComponent(schema)) return buf;
    let firstFieldSchema: any;
    schema.reduceProperties((buf, s) => {
      if (!firstFieldSchema) {
        firstFieldSchema = s;
      }
      s["x-decorator"] = undefined;
      s["x-decorator-props"] = undefined;
      s["x-read-pretty"] = true;
    });
    return buf.concat({
      align: "center",
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
      ...(firstFieldSchema && filterRender && columnProps.filterable
        ? getLightFilterConfig(firstFieldSchema)
        : undefined),
    });
  }, []);
};

const BaseTable: React.FC<TableProps & { filter?: true | "light" }> = observer(
  (props) => {
    const field = useField<ArrayField>();
    const columns = useListTableColumns(
      field.value?.slice(),
      props.filter === "light"
    );
    return (
      <Table
        rowKey="id"
        {...props}
        columns={columns}
        data={field.value?.slice()}
      />
    );
  }
);

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
