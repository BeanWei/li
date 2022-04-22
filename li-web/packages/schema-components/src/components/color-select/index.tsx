import { connect, mapProps, mapReadPretty } from "@formily/react";
import { Select, Tag } from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import { useCompile } from "../../hooks/useCompile";
import { getLocale } from "../__builtins__";

const useColorEnum = (): Record<string, string> => {
  const locale = getLocale();
  return {
    red: locale.ColorSelect.red,
    orangered: locale.ColorSelect.orangered,
    orange: locale.ColorSelect.orange,
    gold: locale.ColorSelect.gold,
    lime: locale.ColorSelect.lime,
    green: locale.ColorSelect.green,
    cyan: locale.ColorSelect.cyan,
    blue: locale.ColorSelect.blue,
    arcobule: locale.ColorSelect.arcobule,
    purple: locale.ColorSelect.purple,
    pinkpurple: locale.ColorSelect.pinkpurple,
    magenta: locale.ColorSelect.magenta,
    gray: locale.ColorSelect.gray,
  };
};

export const ColorSelect = connect(
  (props) => {
    const compile = useCompile();
    const colors = useColorEnum();
    return (
      <Select {...props}>
        {Object.keys(colors).map((color) => (
          <Select.Option value={color}>
            <Tag color={color}>{compile(colors[color] || colors.default)}</Tag>
          </Select.Option>
        ))}
      </Select>
    );
  },
  mapProps((props, field: any) => {
    return {
      ...props,
      suffix: (
        <span>
          {field?.["loading"] || field?.["validating"] ? (
            <IconLoading />
          ) : (
            props.suffix
          )}
        </span>
      ),
    };
  }),
  mapReadPretty((props) => {
    const { value } = props;
    const compile = useCompile();
    const colors = useColorEnum();
    if (!colors[value]) {
      return null;
    }
    return <Tag color={value}>{compile(colors[value] || colors.default)}</Tag>;
  })
);

export default ColorSelect;
