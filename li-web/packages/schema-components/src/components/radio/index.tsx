import { connect, mapProps, mapReadPretty, useField } from "@formily/react";
import { Radio as ArcoRadio, Tag } from "@arco-design/web-react";
import type { RadioProps, RadioGroupProps } from "@arco-design/web-react";
import { isValid } from "@formily/shared";

type ComposedRadio = React.FC<RadioProps> & {
  Group?: React.FC<RadioGroupProps>;
  __ARCO_RADIO?: boolean;
};

export const Radio: ComposedRadio = connect(
  ArcoRadio,
  mapProps({
    value: "checked",
    onInput: "onChange",
  })
);

Radio.__ARCO_RADIO = true;

Radio.Group = connect(
  ArcoRadio.Group,
  mapProps({
    dataSource: "options",
  }),
  mapReadPretty((props) => {
    const { value } = props;
    if (!isValid(value)) {
      return <div>-</div>;
    }
    const field = useField<any>();
    const dataSource = field.dataSource || [];

    return (
      <div>
        {dataSource
          .filter((option: any) => option.value === value)
          .map((option: any, key: any) => (
            <Tag key={key} color={option.color}>
              {option.label}
            </Tag>
          ))}
      </div>
    );
  })
);

export default Radio;
