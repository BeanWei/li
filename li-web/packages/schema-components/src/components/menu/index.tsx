import { RecursionField, useField, useFieldSchema } from "@formily/react";
import {
  Icon,
  Menu as ArcoMenu,
  MenuItemProps,
  MenuProps,
  MenuSubMenuProps,
} from "@arco-design/web-react";

type ComposedMenu = React.FC<MenuProps> & {
  Item?: React.FC<MenuItemProps & { icon: string }>;
  SubMenu?: React.FC<MenuSubMenuProps & { icon: string }>;
  URL?: React.FC<MenuItemProps & { icon: string; href: string }>;
};

export const Menu: ComposedMenu = (props) => {
  const schema = useFieldSchema();
  return (
    <ArcoMenu {...props}>
      <RecursionField schema={schema} onlyRenderProperties />
    </ArcoMenu>
  );
};

Menu.Item = (props) => {
  const { icon, ...rest } = props;
  const schema = useFieldSchema();
  const field = useField();
  return (
    <ArcoMenu.Item {...rest} key={schema.name as string}>
      <Icon type={icon} style={{ marginRight: 5 }} />
      {field.title}
    </ArcoMenu.Item>
  );
};

Menu.SubMenu = (props) => {
  const { icon, ...rest } = props;
  const schema = useFieldSchema();
  const field = useField();
  return (
    <ArcoMenu.SubMenu
      {...rest}
      key={schema.name as string}
      title={
        <>
          <Icon style={{ marginRight: 5 }} type={icon} />
          {field.title}
        </>
      }
    >
      <RecursionField schema={schema} onlyRenderProperties />
    </ArcoMenu.SubMenu>
  );
};

Menu.URL = (props) => {
  const { icon, ...rest } = props;
  const schema = useFieldSchema();
  const field = useField();
  return (
    <ArcoMenu.Item
      {...rest}
      key={schema.name as string}
      onClick={() => {
        window.open(props.href, "_blank");
      }}
    >
      <Icon style={{ marginRight: 5 }} type={icon} />
      {field.title}
    </ArcoMenu.Item>
  );
};

export default Menu;
