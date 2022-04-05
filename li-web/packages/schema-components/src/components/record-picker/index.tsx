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
import { IconLink } from "@arco-design/web-react/icon";
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
  RecordFormDrawer?: React.FC<ActionFormDrawerProps>;
  RecordFormModal?: React.FC<ActionFormModalProps>;
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
    const [visible, setVisible] = useState(false);
    const [selected, setSelected] = useState(toArr(value));
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
          arrowIcon={<IconLink />}
          mode={multiple ? "multiple" : undefined}
          allowClear
          popupVisible={false}
          onVisibleChange={(open) => {
            setVisible(open);
          }}
          options={selected.map((item: any) => {
            return {
              label: item[fieldNames.label],
              value: item[fieldNames.value],
            };
          })}
          value={
            multiple
              ? selected.map((item: any) => item[fieldNames.value])
              : selected?.[0]?.[fieldNames.value]
          }
          onChange={(changed) => {
            if (!changed || !changed.length) {
              onChange?.([]);
              setSelected([]);
            } else if (Array.isArray(changed)) {
              const values = value?.filter((v: any) =>
                changed.includes(v[fieldNames.value])
              );
              onChange?.(values);
              setSelected(values);
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
                    defaultSelectedKeys: selected?.map(
                      (item: any) => item[fieldNames.value]
                    ),
                    preserveSelectedKeys: true,
                    onChange: (_: any, selected: any[]) => {
                      setSelected(selected);
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
      if (s["x-component"] !== "List") {
        s.title = field.value[fieldNames.label];
        s["x-component-props"] = {
          ...s["x-component-props"],
          drawerProps: {
            ...s["x-component-props"]?.["drawerProps"],
            title: s.title,
          },
          modalProps: {
            ...s["x-component-props"]?.["modalProps"],
            title: s.title,
          },
          type: "text",
          icon: null,
        };
      } else {
        s["x-hidden"] = true;
      }
    });

    return (
      <Space>
        {toArr(field.value).map((record) => {
          return (
            <RecordProvider record={record}>
              <RecursionField schema={fieldSchema} onlyRenderProperties />
            </RecordProvider>
          );
        })}
      </Space>
    );
  })
);

RecordPicker.RecordFormDrawer = observer((props) => {
  const forInitVariables = useRecord();
  return <ActionFormDrawer {...props} forInitVariables={forInitVariables} />;
});

RecordPicker.RecordFormModal = observer((props) => {
  const forInitVariables = useRecord();
  return <ActionFormModal {...props} forInitVariables={forInitVariables} />;
});

export default RecordPicker;
