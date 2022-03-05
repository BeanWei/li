import {
  observer,
  RecursionField,
  Schema,
  SchemaExpressionScopeContext,
  useField,
  useFieldSchema,
} from "@formily/react";
import { createContext, useContext, useEffect, useState } from "react";
import { Icon, Menu as ArcoMenu } from "@arco-design/web-react";
import { createPortal } from "react-dom";

function findByUid(schema: Schema, uid: string): any {
  if (!Schema.isSchemaInstance(schema)) {
    schema = new Schema(schema);
  }
  return schema.reduceProperties((buffter, s) => {
    if (s["x-data"] === uid) {
      return s;
    }
    const ss = findByUid(s, uid);
    if (ss) {
      return ss;
    }
    return buffter;
  }, null);
}

function findMenuItem(schema: Schema): any {
  if (!Schema.isSchemaInstance(schema)) {
    schema = new Schema(schema);
  }
  for (const { schema: s } of Schema.getOrderProperties(schema)) {
    if (s["x-component"] === "Menu.Item") {
      return s;
    }
    const ss = findMenuItem(s);
    if (ss) {
      return ss;
    }
  }
  return null;
}

function findKeys(schema: Schema) {
  if (!schema) {
    return;
  }
  const keys = [];
  keys.push(schema.name);
  while (schema.parent) {
    if (schema.parent["x-component"] === "Menu") {
      break;
    }
    keys.push(schema.parent.name);
    schema = schema.parent;
  }
  return keys.reverse();
}

function findKeysByUid(schema: Schema, uid: string) {
  return findKeys(findByUid(schema, uid));
}

type ComposedMenu = React.FC<any> & {
  Item?: React.FC<any>;
  URL?: React.FC<any>;
  SubMenu?: React.FC<any>;
};

const MenuModeContext = createContext<any>({});

const useSideMenuRef = () => {
  const schema = useFieldSchema();
  const scope = useContext(SchemaExpressionScopeContext);
  const scopeKey = schema?.["x-component-props"]?.["sideMenuRefScopeKey"];
  if (!scopeKey) {
    return;
  }
  return scope[scopeKey];
};

export const Menu: ComposedMenu = observer((props) => {
  let {
    onClickMenuItem,
    mode,
    defaultSelectedUid,
    sideMenuRefScopeKey,
    defaultSelectedKeys: dSelectedKeys,
    defaultOpenKeys: dOpenKeys,
    ...rest
  } = props;
  const schema = useFieldSchema();
  const sideMenuRef = useSideMenuRef();
  const [defaultSelectedKeys, setDefaultSelectedKeys] = useState(() => {
    if (dSelectedKeys) {
      return dSelectedKeys;
    }
    if (defaultSelectedUid) {
      return findKeysByUid(schema, defaultSelectedUid);
    }
    return [];
  });
  const [loading, setLoading] = useState(false);
  const [defaultOpenKeys, setDefaultOpenKeys] = useState(() => {
    if (["inline", "mix"].includes(mode)) {
      return dOpenKeys || defaultSelectedKeys;
    }
    return dOpenKeys;
  });
  // @ts-ignore
  const [sideMenuSchema, setSideMenuSchema] = useState<Schema>(() => {
    const key = defaultSelectedKeys?.[0] || null;
    if (mode === "mix" && key) {
      const s = schema.properties?.[key];
      if (s?.["x-component"] === "Menu.SubMenu") {
        return s;
      }
    }
    return null;
  });
  useEffect(() => {
    if (["inline", "mix"].includes(mode)) {
      setDefaultOpenKeys(defaultSelectedKeys);
    }
  }, [defaultSelectedKeys]);
  useEffect(() => {
    const sideMenuElement = sideMenuRef?.current as HTMLElement;
    if (!sideMenuElement) {
      return;
    }
    sideMenuElement.style.display =
      sideMenuSchema?.["x-component"] === "Menu.SubMenu" ? "block" : "none";
  }, [sideMenuSchema?.name, sideMenuRef]);

  return (
    <MenuModeContext.Provider value={mode}>
      <ArcoMenu
        {...rest}
        style={{
          width: mode === "mix" ? "100%" : undefined,
        }}
        onClickMenuItem={(info: any) => {
          const s = schema.properties?.[info.key];
          if (s && mode === "mix") {
            setSideMenuSchema(s);
            if (s["x-component"] !== "Menu.SubMenu") {
              onClickMenuItem && onClickMenuItem(info);
            } else {
              const menuItemSchema = findMenuItem(s);
              if (!menuItemSchema) {
                return;
              }
              // TODO
              setLoading(true);
              const keys = findKeysByUid(schema, menuItemSchema["x-data"]);
              setDefaultSelectedKeys(keys);
              setTimeout(() => {
                setLoading(false);
              }, 100);
              onClickMenuItem &&
                onClickMenuItem({
                  key: menuItemSchema.name,
                  item: {
                    props: {
                      schema: menuItemSchema,
                    },
                  },
                });
            }
          } else {
            onClickMenuItem && onClickMenuItem(info);
          }
        }}
        mode={mode === "mix" ? "horizontal" : mode}
        defaultOpenKeys={defaultOpenKeys}
        defaultSelectedKeys={defaultSelectedKeys}
      >
        <RecursionField schema={schema} onlyRenderProperties />
      </ArcoMenu>
      {loading
        ? null
        : mode === "mix" &&
          sideMenuSchema?.["x-component"] === "Menu.SubMenu" &&
          sideMenuRef?.current?.firstChild &&
          createPortal(
            <MenuModeContext.Provider value="vertical">
              <ArcoMenu
                mode="vertical"
                defaultOpenKeys={defaultOpenKeys}
                defaultSelectedKeys={defaultSelectedKeys}
                onClickMenuItem={(info) => {
                  onClickMenuItem && onClickMenuItem(info);
                }}
              >
                <RecursionField schema={sideMenuSchema} onlyRenderProperties />
              </ArcoMenu>
            </MenuModeContext.Provider>,
            sideMenuRef.current.firstChild
          )}
    </MenuModeContext.Provider>
  );
});

Menu.Item = observer((props) => {
  const { icon, ...others } = props;
  const schema = useFieldSchema();
  const field = useField();
  return (
    <ArcoMenu.Item
      {...others}
      key={schema.name}
      eventKey={schema.name}
      schema={schema}
    >
      <Icon type={icon} style={{ marginRight: 5 }} />
      {field.title}
    </ArcoMenu.Item>
  );
});

Menu.URL = observer((props) => {
  const { icon, ...others } = props;
  const schema = useFieldSchema();
  const field = useField();
  return (
    <ArcoMenu.Item
      {...others}
      key={schema.name}
      eventKey={schema.name}
      schema={schema}
      onClick={() => {
        window.open(props.href, "_blank");
      }}
    >
      <Icon style={{ marginRight: 5 }} type={icon} />
      {field.title}
    </ArcoMenu.Item>
  );
});

Menu.SubMenu = observer((props) => {
  const { icon, ...others } = props;
  const schema = useFieldSchema();
  const field = useField();
  const mode = useContext(MenuModeContext);
  if (mode === "mix") {
    return <ArcoMenu.Item {...props} />;
  }
  return (
    <ArcoMenu.SubMenu
      {...others}
      key={schema.name}
      eventKey={schema.name}
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
});

export default Menu;
