import { createContext, useContext, useEffect } from "react";
import { useRequest } from "pro-utils";
import { Navigate } from "react-router";
import { useLocalStorageState } from "ahooks";
import { ConfigProvider, Layout as ArcoLayout } from "@arco-design/web-react";
import { useParams } from "react-router";
import { useNavigate } from "react-router-dom";
import { SchemaComponent, zhCN, enUS } from "schema-components";
import Logo from "@/assets/logo.svg";
import { useRoute } from "../route-switch/hooks";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import styles from "./index.module.less";
import { Loading } from "../components";

const getOpenKeysFromMenuData = (curKey: string, menuData?: any[]) => {
  return (menuData || []).reduce((pre, item) => {
    if (curKey == item.key) {
      const paths = item.path.split(".");
      return pre.concat(paths.slice(0, paths.length - 1));
    }
    if (item.children) {
      const newArray: string[] = pre.concat(
        getOpenKeysFromMenuData(curKey, item.children) || []
      );
      return newArray;
    }
    return pre;
  }, [] as string[]);
};

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
  const navigate = useNavigate();
  const params = useParams();

  const {
    title = "Li Admin",
    logo,
    navitems = [],
    home,
    menus = [],
    entry,
  } = route.config;

  const [theme, setTheme] = useLocalStorageState("li-theme", {
    defaultValue: window.matchMedia?.("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light",
  });
  const [lang, setLang] = useLocalStorageState("li-lang", {
    defaultValue: navigator.language,
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
    return <Loading />;
  }
  if (result.error) {
    return <Navigate to={entry + "/sign"} />;
  }

  const curKey = params?.["*"] || home || menus[0]?.key;
  const onClickMenuItem = (key: string) => {
    navigate(entry + `/${key}`);
  };

  const global = {
    app: {
      title,
      logo,
      home,
      entry,
    },
    lang,
    currentUser: result.data,
  };

  return (
    <ConfigProvider
      locale={lang === "en-US" ? enUS : zhCN}
      componentConfig={{
        Card: {
          bordered: false,
        },
        List: {
          bordered: false,
        },
        Table: {
          border: false,
        },
      }}
    >
      <LayoutContext.Provider
        value={{
          ...global,
          lang,
          setLang,
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
                          defaultSelectedKeys: [curKey],
                          defaultOpenKeys: getOpenKeysFromMenuData(
                            curKey,
                            menus
                          ),
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
    </ConfigProvider>
  );
};
