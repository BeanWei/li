import { connect, mapProps, mapReadPretty } from "@formily/react";
import { Select, Tag } from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import { useCompile } from "../../hooks/useCompile";
import "@arco-design/web-react/lib/Select/style/index";
import "@arco-design/web-react/lib/Tag/style/index";

const colors: Record<string, any> = {
  red: '{{t("Red")}}',
  orangered: '{{t("Orangered")}}',
  orange: '{{t("Orange")}}',
  gold: '{{t("Gold")}}',
  lime: '{{t("Lime")}}',
  green: '{{t("Green")}}',
  cyan: '{{t("Cyan")}}',
  blue: '{{t("Blue")}}',
  arcobule: '{{t("Arco Blue")}}',
  purple: '{{t("Purple")}}',
  pinkpurple: '{{t("Pink Purple")}}',
  magenta: '{{t("Magenta")}}',
  gray: '{{t("Gray")}}',
};

export const ColorSelect = connect(
  (props) => {
    const compile = useCompile();
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
    const compile = useCompile();
    const { value } = props;
    if (!colors[value]) {
      return null;
    }
    return <Tag color={value}>{compile(colors[value] || colors.default)}</Tag>;
  })
);

export default ColorSelect;
