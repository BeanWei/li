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
  lang?: string;
  setLang?: (value: string) => void;
  theme?: string;
  setTheme?: (value: string) => void;
  profile?: Record<string, any>;
  setProfile?: (value: Record<string, any>) => void;
}>({});

export const useLayoutContext = () => {
  return useContext(LayoutContext);
};

export const LayoutProvider = (props: any) => {
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

  const result = useRequest("userGetProfile");
  if (result.loading) {
    return <Spin />;
  }
  if (result.error) {
    return <Redirect to={"/signin"} />;
  }
  return (
    <LayoutContext.Provider
      value={{
        lang: result?.data?.language,
        profile: result.data,
        theme,
        setTheme,
      }}
    >
      {props.children}
    </LayoutContext.Provider>
  );
};

export const Layout = () => {
  const route = useRoute();
  const history = useHistory();
  const match = useRouteMatch<any>();
  const curKey =
    match.params.name ??
    route?.["schema"]?.["x-component-props"]?.["defaultSelectedKeys"]?.[0];
  const onClickMenuItem = (key: string) => {
    history.push(`/admin/${key}`);
  };

  const {
    title = "Li Admin",
    logo,
    navitems = [],
    home,
    menus = [],
  } = route.config;

  return (
    <LayoutProvider>
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
                console.log(item);
                return (
                  <li key={i.toString()}>
                    <SchemaComponent schema={item} />
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
                scope={{ onClickMenuItem }}
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
    </LayoutProvider>
  );
};
