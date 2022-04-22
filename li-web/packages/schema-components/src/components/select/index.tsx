import React from "react";
import {
  connect,
  mapReadPretty,
  mapProps,
  observer,
  useField,
} from "@formily/react";
import {
  Select as ArcoSelect,
  SelectProps as ArcoSelectProps,
  Space,
  Tag,
} from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import { isValid, toArr } from "@formily/shared";
import { isArrayField } from "@formily/core";
import { useTranslation } from "react-i18next";
import { getCurrentOptions } from "./shared";

type SelectProps = ArcoSelectProps & {
  objectValue?: boolean;
  onChange?: (v: any, option: any) => void;
};

const isEmptyObject = (val: any) =>
  !isValid(val) || (typeof val === "object" && Object.keys(val).length === 0);

const ObjectSelect = (props: SelectProps) => {
  const { value, options, onChange, mode, ...others } = props;
  const toValue = (v: any) => {
    if (isEmptyObject(v)) {
      return;
    }
    const values = toArr(v)
      .filter((item) => item)
      .map((val) => {
        return typeof val === "object" ? val.value : val;
      });
    const current = getCurrentOptions(values, options)?.map((val: any) => {
      return {
        label: val.label,
        value: val.value,
      };
    });
    if (mode && ["tags", "multiple"].includes(mode)) {
      return current;
    }
    return current.shift();
  };
  return (
    <ArcoSelect
      value={toValue(value)}
      allowClear
      labelInValue
      options={options}
      onChange={(changed) => {
        const current = getCurrentOptions(
          toArr(changed).map((v) => v.value),
          options
        );
        if (mode && ["tags", "multiple"].includes(mode)) {
          onChange?.(current, options);
        } else {
          onChange?.(current.shift(), options);
        }
      }}
      mode={mode}
      {...others}
    />
  );
};

const ReadPretty = observer((props: any) => {
  const field = useField<any>();
  if (!isValid(props.value)) {
    return <div>-</div>;
  }
  if (isArrayField(field) && field?.value?.length === 0) {
    return <div>-</div>;
  }
  const dataSource = field.dataSource || props.options || [];
  const options = getCurrentOptions(field.value, dataSource);
  return (
    <Space wrap>
      {options.map((option: any, key: any) => (
        <Tag key={key} color={option.color}>
          {option.label}
        </Tag>
      ))}
    </Space>
  );
});

export const Select: React.FC<SelectProps> = connect(
  (props: SelectProps) => {
    const { objectValue, ...rest } = props;
    const { t } = useTranslation();
    if (objectValue) {
      return (
        <ObjectSelect
          {...rest}
          placeholder={rest.placeholder ? t(rest.placeholder) : undefined}
        />
      );
    }
    return (
      <ArcoSelect
        {...rest}
        placeholder={rest.placeholder ? t(rest.placeholder) : undefined}
        value={rest.value || undefined}
      />
    );
  },
  mapProps(
    {
      dataSource: "options",
      loading: true,
    },
    (props, field: any) => {
      return {
        ...props,
        suffixIcon:
          field?.["loading"] || field?.["validating"] ? (
            <IconLoading />
          ) : (
            props.suffixIcon
          ),
      };
    }
  ),
  mapReadPretty(ReadPretty)
);

export default Select;
