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
  Tag,
} from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import { isValid, toArr } from "@formily/shared";
import { isArrayField } from "@formily/core";
import { getCurrentOptions } from "./shared";
import "@arco-design/web-react/lib/Select/style/index";
import "@arco-design/web-react/lib/Tag/style/index";

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
  const fieldNames = { ...props.fieldNames };
  const field = useField<any>();
  if (!isValid(props.value)) {
    return <div />;
  }
  if (isArrayField(field) && field?.value?.length === 0) {
    return <div />;
  }
  const dataSource = field.dataSource || props.options || [];
  const options = getCurrentOptions(field.value, dataSource);
  return (
    <div>
      {options.map((option: any, key: any) => (
        <Tag key={key} color={option.color}>
          {option.label}
        </Tag>
      ))}
    </div>
  );
});

export const Select: React.FC<SelectProps> = connect(
  (props: SelectProps) => {
    const { objectValue, ...others } = props;
    if (objectValue) {
      return <ObjectSelect {...others} />;
    }
    return <ArcoSelect {...others} value={others.value || undefined} />;
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
