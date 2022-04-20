import React, { Fragment, useContext, useEffect, useMemo } from "react";
import {
  Button,
  ConfigProvider,
  DatePicker,
  Form,
  Input,
  Select,
  SelectProps,
  Table,
  TimePicker,
} from "@arco-design/web-react";
import {
  IconDown,
  IconFilter,
  IconRefresh,
  IconSearch,
  IconUp,
} from "@arco-design/web-react/icon";
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
import { isObject, pickBy, sortBy } from "lodash";
import { isValid } from "@formily/shared";
import { RecordIndexProvider, RecordProvider } from "../../core";
import { useAttach } from "../../hooks";
import {
  isColumnComponent,
  ObservableColumnSource,
  useArrayTableSources,
} from "../array-table";
import { ComposedListTable, ListProps } from "./types";
import { ListContext, ListContextProps } from "./context";
import { getLocale, useCollapseGrid } from "../__builtins__";
import FormGrid from "../form-grid";
import FormButtonGroup from "../form-button-group";
import { BaseRecordSelect } from "../record-select";

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
        case "RecordSelect":
          // toJson 复制一份出来，防止污染原 schema
          const schema_ = schema.toJSON(true);
          filterNode = (
            <BaseRecordSelect
              {...schema_["x-component-props"]}
              multiple
              value={filterKeys}
              // @ts-ignore
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
  source: ObservableColumnSource[],
  dataSource: any[],
  filterRender?: boolean
): TableProps<any>["columns"] => {
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

const FilterForm: React.FC<{
  source: ObservableColumnSource[];
  onSearch?: (values: Record<string, any>) => void;
}> = observer(({ source, onSearch }) => {
  const locale = getLocale();
  const { grid, toggle, expanded, type } = useCollapseGrid(2);
  const [form] = Form.useForm();
  const fieldSchemas: Schema[] = [];
  source.forEach((item) => {
    if (item.schema["x-component-props"]?.filterable) {
      item.schema.reduceProperties((b: any, s) => {
        if (!s.title) {
          s.title = item.schema["x-component-props"]?.title;
        }
        fieldSchemas.push(s);
      });
    }
  });

  const handleSubmit = () => {
    const values = form.getFieldsValue();
    onSearch?.(pickBy(values, isValid));
  };

  const handleReset = () => {
    form.resetFields();
    onSearch?.({});
  };

  const renderActions = () => {
    return (
      <FormButtonGroup
        align="right"
        style={{
          marginBottom: 16,
          width: "100%",
          justifyContent: "end",
        }}
      >
        <Button
          key="submit"
          type="primary"
          icon={<IconSearch />}
          onClick={handleSubmit}
        >
          {locale.List.query}
        </Button>
        <Button key="reset" icon={<IconRefresh />} onClick={handleReset}>
          {locale.List.reset}
        </Button>
      </FormButtonGroup>
    );
  };

  const renderButtonGroup = () => {
    if (type === "collapsible") {
      return (
        <>
          {renderActions()}
          <FormButtonGroup
            align="right"
            style={{ marginLeft: 8, marginTop: 14 }}
          >
            <Button
              type="text"
              icon={expanded ? <IconUp /> : <IconDown />}
              onClick={toggle}
            />
          </FormButtonGroup>
        </>
      );
    }
    return renderActions();
  };

  return (
    <div
      style={{
        borderBottom: "1px solid var(--color-border-1)",
        display: "flex",
        marginBottom: 16,
      }}
    >
      <Form form={form} layout="vertical">
        <FormGrid grid={grid}>
          {fieldSchemas.map((schema) => {
            let filterNode = null;
            let span = 1;
            switch (schema["x-component"]) {
              case "Checkbox":
              case "Switch":
                filterNode = (
                  <Select
                    options={
                      [
                        { label: "Yes", value: true },
                        { label: "No", value: false },
                      ] as any
                    }
                    allowClear
                  />
                );
                break;
              case "DatePicker":
              case "DatePicker.RangePicker":
                filterNode = (
                  <DatePicker.RangePicker
                    allowClear
                    mode={schema["x-component-props"]?.mode}
                  />
                );
                span = 2;
                break;
              case "TimePicker":
              case "TimePicker.RangePicker":
                filterNode = <TimePicker.RangePicker allowClear />;
                span = 2;
                break;
              case "RecordSelect":
                // toJson 复制一份出来，防止污染原 schema
                const schema_ = schema.toJSON(true);
                filterNode = (
                  <BaseRecordSelect
                    {...schema_["x-component-props"]}
                    multiple
                  />
                );
                break;
              case "InputNumber":
              case "Money":
              case "Rate":
              case "Slider":
              // TODO: 数字范围选择器
              default:
                if (schema.enum?.length) {
                  filterNode = (
                    <Select
                      options={schema.enum as SelectProps["options"]}
                      mode="multiple"
                      allowClear
                    />
                  );
                } else {
                  filterNode = <Input allowClear />;
                }
                break;
            }
            return (
              <FormGrid.GridColumn gridSpan={span} key={schema.name}>
                <Form.Item
                  label={schema.title}
                  field={schema.name}
                  style={{
                    display: "block",
                    marginBottom: 16,
                  }}
                >
                  {filterNode}
                </Form.Item>
              </FormGrid.GridColumn>
            );
          })}
          <FormGrid.GridColumn
            gridSpan={-1}
            style={{
              display: "flex",
              justifyContent: "space-end",
              alignItems: "flex-end",
            }}
          >
            {renderButtonGroup()}
          </FormGrid.GridColumn>
        </FormGrid>
      </Form>
    </div>
  );
});

const BaseTable: React.FC<
  TableProps & {
    filter?: ListProps["filter"];
    reload?: ListContextProps["reload"];
  }
> = observer((props) => {
  const { reload, ...rest } = props;
  const field = useField<ArrayField>();
  const schema = useFieldSchema();
  const source = useArrayTableSources();
  const columns = useListTableColumns(
    source,
    field.value?.slice(),
    props.filter === "light"
  );
  return (
    <>
      {props.filter === true && (
        <FilterForm
          source={source}
          onSearch={(values) => reload?.({ filter: values })}
        />
      )}
      <RecursionField schema={schema} onlyRenderProperties />
      <Table
        rowKey="id"
        {...rest}
        columns={sortBy(
          // @ts-ignore
          columns?.filter((col) => !!!col.hideInTable),
          "order"
        )}
        data={field.value?.slice()}
      />
    </>
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
      form.setValues({
        [schema.name as string]: ctx.tableProps?.data,
      });
    }, [ctx.tableProps?.data]);

    return (
      <FormContext.Provider value={form}>
        <FieldContext.Provider value={f}>
          <BaseTable {...props} {...ctx.tableProps} reload={ctx.reload} />
        </FieldContext.Provider>
      </FormContext.Provider>
    );
  }
);

ListTable.Column = () => {
  return <Fragment />;
};

export default ListTable;
