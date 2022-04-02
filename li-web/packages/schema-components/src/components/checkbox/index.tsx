import React from "react";
import { connect, mapProps, mapReadPretty, useField } from "@formily/react";
import { Checkbox as ArcoCheckbox, Tag } from "@arco-design/web-react";
import { IconCheck, IconClose } from "@arco-design/web-react/icon";
import type { CheckboxProps, CheckboxGroupProps } from "@arco-design/web-react";
import { isValid } from "@formily/shared";
import { uniq } from "lodash";

type ComposedCheckbox = React.FC<CheckboxProps> & {
  Group?: React.FC<CheckboxGroupProps<any>>;
  __ARCO_CHECKBOX?: boolean;
};

export const Checkbox: ComposedCheckbox = connect(
  ArcoCheckbox,
  mapProps({
    value: "checked",
    onInput: "onChange",
  }),
  mapReadPretty((props) => {
    if (!isValid(props.value)) {
      return <div>-</div>;
    }
    return props.value ? (
      <IconCheck style={{ color: "#52c41a" }} />
    ) : (
      <IconClose style={{ color: "#f5222d" }} />
    );
  })
);

Checkbox.__ARCO_CHECKBOX = true;

Checkbox.Group = connect(
  ArcoCheckbox.Group,
  mapProps({
    dataSource: "options",
  }),
  mapReadPretty((props) => {
    if (!isValid(props.value)) {
      return <div>-</div>;
    }

    const field = useField<any>();
    const dataSource = field.dataSource || [];
    const value = uniq(field.value ? field.value : []);

    return (
      <div>
        {dataSource
          .filter((option: any) => value.includes(option.value))
          .map((option: any, key: any) => (
            <Tag key={key} color={option.color}>
              {option.label}
            </Tag>
          ))}
      </div>
    );
  })
);

export default Checkbox;
