import { createContext, useContext, useEffect } from "react";
import { useRequest } from "pro-utils";
import { Redirect } from "react-router";
import { useLocalStorageState } from "ahooks";
import { Layout as ArcoLayout, Spin } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import Logo from "@/assets/logo.svg";
import { useRoute } from "../route-switch/hooks";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import styles from "./index.module.less";

export const LayoutContext = createContext<{
  app?: Record<string, any>;
  lang?: string;
  setLang?: (value: string) => void;
  theme?: string;
  setTheme?: (value: string) => void;
  currentUser?: Record<string, any>;
  setCurrentUser?: (value: Record<string, any>) => void;
}>({});

export const useLayoutContext = () => {
  return useContext(LayoutContext);
};

export const Layout = () => {
  const route = useRoute();
  const history = useHistory();
  const match = useRouteMatch<any>();

  const {
    title = "Li Admin",
    logo,
    navitems = [],
    home,
    menus = [],
    entry,
  } = route.config;

  const [theme, setTheme] = useLocalStorageState("li-theme", {
    defaultValue: "light",
  });

  useEffect(() => {
    if (theme === "dark") {
      document.body.setAttribute("arco-theme", "dark");
    } else {
      document.body.removeAttribute("arco-theme");
    }
  }, [theme]);

  const result = useRequest("@getCurrentUser");
  if (result.loading) {
    return <Spin />;
  }
  if (result.error) {
    return <Redirect to={entry + "/sign"} />;
  }

  const curKey =
    match.params.name ??
    route?.["schema"]?.["x-component-props"]?.["defaultSelectedKeys"]?.[0];
  const onClickMenuItem = (key: string) => {
    history.push(entry + `/${key}`);
  };

  const global = {
    app: {
      title,
      logo,
      home,
      entry,
    },
    lang: result?.data?.language,
    currentUser: result.data.data,
  };

  return (
    <LayoutContext.Provider
      value={{
        ...global,
        theme,
        setTheme,
      }}
    >
      <ArcoLayout className={styles.layout}>
        <div className={styles["layout-navbar"]}>
          <div className={styles.navbar}>
            <div className={styles.left}>
              <div className={styles.logo}>
                {logo ? <img src={logo} /> : <Logo />}
                <div className={styles["logo-name"]}>{title}</div>
              </div>
            </div>
            <ul className={styles.right}>
              {navitems.map((item: any, i: number) => {
                return (
                  <li key={i.toString()}>
                    <SchemaComponent schema={item} scope={{ global }} />
                  </li>
                );
              })}
            </ul>
          </div>
        </div>
        <ArcoLayout>
          <ArcoLayout.Sider
            className={styles["layout-sider"]}
            trigger={null}
            collapsible
            breakpoint="xl"
            width={220}
            style={{ paddingTop: 60 }}
          >
            <div className={styles["menu-wrapper"]}>
              <SchemaComponent
                schema={{
                  type: "void",
                  properties: {
                    menu: {
                      "x-component": "Menu",
                      "x-component-props": {
                        defaultSelectedKeys: [home || menus[0]?.key],
                        onClickMenuItem: "{{ onClickMenuItem }}",
                        menuData: menus,
                      },
                    },
                  },
                }}
                scope={{ onClickMenuItem, global }}
              />
            </div>
          </ArcoLayout.Sider>
          <ArcoLayout
            className={styles["layout-content"]}
            style={{ paddingLeft: 220, paddingTop: 60 }}
          >
            <div className={styles["layout-content-wrapper"]}>
              <ArcoLayout.Content>
                <RemoteSchemaComponent uid={curKey} />
              </ArcoLayout.Content>
            </div>
          </ArcoLayout>
        </ArcoLayout>
      </ArcoLayout>
    </LayoutContext.Provider>
  );
};
