import {
  Icon,
  Menu as ArcoMenu,
  MenuProps as ArcoMenuProps,
} from "@arco-design/web-react";
import { uid } from "@formily/shared";
import { isUrl } from "../__builtins__";

export type MenuDataItem = {
  /** @name 对应页面的key */
  key?: string;
  /** @name 子菜单 */
  children?: MenuDataItem[];
  /** @name 菜单的icon */
  icon?: string;
  /** @name 菜单的名字 */
  name?: string;
  /** @name disable 菜单选项 */
  disabled?: boolean;
  /** @name 指定外链打开形式，同a标签 */
  target?: string;
};

const getNavMenuItems = (menusData: MenuDataItem[] = []): React.ReactNode[] => {
  const getSubMenuOrItem = (item: MenuDataItem): React.ReactNode => {
    if (Array.isArray(item.children) && item && item.children.length > 0) {
      return (
        <ArcoMenu.SubMenu
          key={item.key || item.target || uid()}
          title={
            <>
              <Icon style={{ marginRight: 5 }} type={item.icon} />
              {item.name}
            </>
          }
        >
          {getNavMenuItems(item.children)}
        </ArcoMenu.SubMenu>
      );
    }
    return (
      <ArcoMenu.Item
        key={item.key || item.target || uid()}
        disabled={item.disabled}
      >
        {getMenuItemPath(item)}
      </ArcoMenu.Item>
    );
  };

  const getMenuItemPath = (item: MenuDataItem) => {
    const itemPath = conversionPath(item.target || "/");
    // Is it a http link
    if (isUrl(itemPath)) {
      return (
        <span
          title={item.name}
          onClick={(e) => {
            e.stopPropagation();
            e.preventDefault();
            window?.open?.(itemPath);
          }}
        >
          <Icon style={{ marginRight: 5 }} type={item.icon} />
          {item.name}
        </span>
      );
    }
    return (
      <>
        <Icon style={{ marginRight: 5 }} type={item.icon} />
        {item.name}
      </>
    );
  };

  const conversionPath = (path: string) => {
    if (path && path.indexOf("http") === 0) {
      return path;
    }
    return `/${path || ""}`.replace(/\/+/g, "/");
  };

  return menusData.map((item) => getSubMenuOrItem(item)).filter((item) => item);
};

type MenuProps = ArcoMenuProps & {
  menuData?: MenuDataItem[];
};

export const Menu: React.FC<MenuProps> = (props) => {
  const { menuData, ...rest } = props;
  return <ArcoMenu {...rest}>{getNavMenuItems(menuData)}</ArcoMenu>;
};

export default Menu;
