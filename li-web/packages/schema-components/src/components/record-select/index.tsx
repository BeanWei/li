import { useCallback } from "react";
import {
  Avatar,
  List,
  Select,
  SelectProps,
  Space,
  Spin,
  Tag,
} from "@arco-design/web-react";
import { isValid, toArr } from "@formily/shared";
import { connect, mapReadPretty, useField } from "@formily/react";
import { debounce, unionBy } from "lodash";
import { useRequest } from "pro-utils";
import { isUrl } from "../__builtins__";

type RecordSelectProps = Omit<SelectProps, "onChange" | "mode" | "value"> & {
  value?: Record<string, any> | Record<string, any>[];
  onChange?: (value: Record<string, any> | Record<string, any>[]) => void;
  multiple?: boolean;
  fieldNames?: {
    title?: string;
    value?: string;
    avatar?: string;
    description?: string[];
  };
  searchConfig: {
    operation: string;
    variables?: Record<string, any>;
  };
};

export const BaseRecordSelect: React.FC<RecordSelectProps> = (props) => {
  const {
    value: value_,
    onChange,
    multiple,
    fieldNames: fieldNames_,
    searchConfig,
    ...rest
  } = props;
  const fieldNames = { title: "id", value: "id", ...fieldNames_ };
  const value = toArr(value_).filter((item: any) => !!item[fieldNames.value]);

  const {
    data = {},
    loading,
    run,
  } = useRequest(searchConfig.operation, searchConfig.variables, {
    manual: true,
  });
  const debouncedSearch = useCallback(
    debounce((inputValue: string) => {
      run({
        ...searchConfig.variables,
        query: inputValue,
      });
    }, 500),
    []
  );
  const options = unionBy(value, data.list || [], fieldNames.value);

  return (
    <Select
      {...rest}
      mode={multiple ? "multiple" : undefined}
      allowClear
      filterOption={false}
      notFoundContent={
        loading ? (
          <div
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <Spin style={{ margin: 12 }} />
          </div>
        ) : null
      }
      onSearch={debouncedSearch}
      renderFormat={(option) => {
        return options.find((v) => v[fieldNames.value] === option?.value)?.[
          fieldNames.title
        ];
      }}
      options={options.map((item: any) => {
        return {
          label: (
            <List.Item key={item[fieldNames.value]}>
              <List.Item.Meta
                avatar={
                  fieldNames.avatar && item[fieldNames.avatar] ? (
                    <Avatar shape="square">
                      {isUrl(item[fieldNames.avatar]) ? (
                        <img
                          src={item[fieldNames.avatar]}
                          alt={item[fieldNames.title]}
                        />
                      ) : (
                        item[fieldNames.avatar]
                      )}
                    </Avatar>
                  ) : null
                }
                title={item[fieldNames.title]}
                description={
                  fieldNames.description
                    ? fieldNames.description
                        .map((field) => item[field])
                        .filter((v) => isValid(v))
                        .join(", ")
                    : null
                }
              />
            </List.Item>
          ),
          value: item[fieldNames.value],
        };
      })}
      value={
        multiple
          ? value.map((item: any) => item[fieldNames.value])
          : value?.[0]?.[fieldNames.value]
      }
      onChange={(changed) => {
        if (!changed || !changed.length) {
          onChange?.([]);
        } else if (Array.isArray(changed)) {
          const values = options?.filter((v: any) =>
            changed.includes(v[fieldNames.value])
          );
          onChange?.(values);
        }
      }}
    />
  );
};

export const RecordSelect = connect(
  BaseRecordSelect,
  mapReadPretty((props: RecordSelectProps) => {
    const field = useField<any>();
    if (!isValid(props.value)) {
      return <div>-</div>;
    }
    if (props.multiple && !field?.value?.length) {
      return <div>-</div>;
    }
    const fieldNames = { title: "id", value: "id", ...props.fieldNames };
    const values = toArr(field.value).filter(
      (item: any) => !!item[fieldNames.value]
    );

    return (
      <Space wrap>
        {values.map((item: any) => (
          <Tag key={item[fieldNames.value]}>{item[fieldNames.title]}</Tag>
        ))}
      </Space>
    );
  })
);

export default RecordSelect;
