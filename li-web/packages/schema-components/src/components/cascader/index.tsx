import React from "react";
import { connect, mapReadPretty, mapProps, useField } from "@formily/react";
import {
  Cascader as ArcoCascader,
  CascaderProps,
} from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import { toArr } from "@formily/shared";
import { ArrayField } from "@formily/core";
import { DefaultFieldNames } from "@arco-design/web-react/es/Cascader/cascader";

interface CascaderRef {
  focus: () => void;
  blur: () => void;
}

type FixArcoCascaderType = React.ForwardRefExoticComponent<
  CascaderProps &
    React.RefAttributes<CascaderRef> & {
      suffixIcon?: React.ReactNode;
    }
>;

export const Cascader = connect(
  ArcoCascader as FixArcoCascaderType,
  mapProps(
    {
      dataSource: "options",
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
  mapReadPretty((props) => {
    const { fieldNames = DefaultFieldNames } = props;
    const values = toArr(props.value);
    const len = values.length;
    const field = useField<ArrayField>();
    let dataSource = field.dataSource;
    const data = [];
    for (const item of values) {
      if (typeof item === "object") {
        data.push(item);
      } else {
        const curr = dataSource?.find((v) => v[fieldNames.value] === item);
        dataSource = curr?.[fieldNames.children] || [];
        data.push(curr || { label: item, value: item });
      }
    }
    return (
      <div>
        {data.map((item, index) => {
          return (
            <span key={index}>
              {typeof item === "object" ? item[fieldNames.label] : item}
              {len > index + 1 && " / "}
            </span>
          );
        })}
      </div>
    );
  })
);

export default Cascader;
