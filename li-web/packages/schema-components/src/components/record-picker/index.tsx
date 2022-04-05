import { useState } from "react";
import {
  connect,
  mapReadPretty,
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { toArr } from "@formily/shared";
import { Field } from "@formily/core";
import { Drawer, Select, SelectProps, Space } from "@arco-design/web-react";
import { RecordProvider, SchemaComponentOptions, useRecord } from "../../core";
import { ActionFormDrawerProps, ActionFormModalProps } from "../action/types";
import ActionFormDrawer from "../action/Action.FormDrawer";
import ActionFormModal from "../action/Action.FormModal";
import { getLocale } from "../__builtins__";

type RecordPickerProps = Omit<SelectProps, "onChange" | "mode" | "value"> & {
  value?: Record<string, any> | Record<string, any>[];
  onChange?: (value: Record<string, any> | Record<string, any>[]) => void;
  multiple?: boolean;
  fieldNames?: {
    label?: string;
    value?: string;
  };
};

type ComposedRecordPicker = React.FC<RecordPickerProps> & {
  RecordFormDrawer?: React.FC<
    ActionFormDrawerProps & { fieldNames: RecordPickerProps["fieldNames"] }
  >;
  RecordFormModal?: React.FC<
    ActionFormModalProps & { fieldNames: RecordPickerProps["fieldNames"] }
  >;
};

export const RecordPicker: ComposedRecordPicker = connect(
  (props: RecordPickerProps) => {
    const {
      value,
      multiple,
      onChange,
      fieldNames: fieldNames_,
      ...reset
    } = props;
    const fieldNames = { label: "id", value: "id", ...fieldNames_ };
    const values = toArr(value);
    const [visible, setVisible] = useState(false);
    const locale = getLocale();
    const fieldSchema = useFieldSchema();
    fieldSchema.reduceProperties((b, s) => {
      if (s["x-component"] === "List") {
        s["x-component-props"] = {
          ...s["x-component-props"],
          useProps: "{{useListSelectionProps}}",
        };
      } else {
        s["x-hidden"] = true;
      }
    });

    return (
      <>
        <Select
          {...reset}
          mode={multiple ? "multiple" : undefined}
          allowClear
          popupVisible={false}
          onVisibleChange={(open) => {
            setVisible(open);
          }}
          options={values.map((item: any) => {
            return {
              label: item[fieldNames.label],
              value: item[fieldNames.value],
            };
          })}
          value={
            multiple
              ? values.map((item: any) => item[fieldNames.value])
              : values?.[0]?.[fieldNames.value]
          }
          onChange={(changed) => {
            if (!changed || !changed.length) {
              onChange?.([]);
            } else if (Array.isArray(changed)) {
              const values = value?.filter((v: any) =>
                changed.includes(v[fieldNames.value])
              );
              onChange?.(values);
            }
          }}
        />
        <Drawer
          // FIXME: locale.RecordPicker is undefined
          title={locale.RecordPicker?.drawerTitle || "Please select"}
          width="80%"
          mountOnEnter
          unmountOnExit
          visible={visible}
          footer={null}
          onCancel={() => setVisible(false)}
        >
          <SchemaComponentOptions
            scope={{
              useListSelectionProps: () => {
                return {
                  selection: {
                    enableCheckAll: true,
                    multiple,
                    defaultSelectedKeys: values?.map(
                      (item: any) => item[fieldNames.value]
                    ),
                    preserveSelectedKeys: true,
                    onChange: (_: any, selected: any[]) => {
                      onChange?.(selected);
                    },
                  },
                };
              },
            }}
          >
            <RecursionField schema={fieldSchema} onlyRenderProperties />
          </SchemaComponentOptions>
        </Drawer>
      </>
    );
  },
  mapReadPretty((props: RecordPickerProps) => {
    const fieldNames = { label: "id", value: "id", ...props.fieldNames };
    const field = useField<Field>();
    const fieldSchema = useFieldSchema();
    fieldSchema.reduceProperties((b, s) => {
      if (s["x-component"] === "List") {
        s["x-hidden"] = true;
      } else if (
        s["x-component"] === "RecordPicker.RecordFormDrawer" ||
        s["x-component"] === "RecordPicker.RecordFormModal"
      ) {
        s["x-component-props"] = {
          ...s["x-component-props"],
          fieldNames,
        };
      }
    });

    return (
      <Space
        size={0}
        split={<span style={{ marginRight: 4, color: "#aaa" }}>, </span>}
      >
        {toArr(field.value).map((record) => {
          return (
            <RecordProvider record={record} key={record[fieldNames.value]}>
              <RecursionField schema={fieldSchema} onlyRenderProperties />
            </RecordProvider>
          );
        })}
      </Space>
    );
  })
);

RecordPicker.RecordFormDrawer = observer((props) => {
  const { fieldNames: fieldNames_, ...rest } = props;
  const fieldNames = { label: "id", value: "id", ...fieldNames_ };
  const forInitVariables = useRecord();
  return (
    <ActionFormDrawer
      {...rest}
      isMenuItem
      actionText={
        <a style={{ cursor: "pointer", color: "rgb(var(--primary-6))" }}>
          {forInitVariables[fieldNames.label]}
        </a>
      }
      drawerProps={{
        ...rest.drawerProps,
        title: forInitVariables[fieldNames.label],
      }}
      forInitVariables={forInitVariables}
    />
  );
});

RecordPicker.RecordFormModal = observer((props) => {
  const { fieldNames: fieldNames_, ...rest } = props;
  const fieldNames = { label: "id", value: "id", ...fieldNames_ };
  const forInitVariables = useRecord();
  return (
    <ActionFormModal
      {...rest}
      isMenuItem
      actionText={
        <a style={{ cursor: "pointer", color: "rgb(var(--primary-6))" }}>
          {forInitVariables[fieldNames.label]}
        </a>
      }
      forInitVariables={forInitVariables}
    />
  );
});

export default RecordPicker;
