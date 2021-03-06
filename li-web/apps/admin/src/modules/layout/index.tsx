import { useContext, useEffect, useState } from "react";
import { Navigate } from "react-router";
import { Layout as ArcoLayout } from "@arco-design/web-react";
import { useParams } from "react-router";
import { useNavigate } from "react-router-dom";
import { isEmpty } from "@formily/shared";
import { IconMenuFold, IconMenuUnfold } from "@arco-design/web-react/icon";
import NProgress from "nprogress";
import { useTranslation } from "react-i18next";
import { SchemaComponent } from "schema-components";
import { useRequest } from "pro-utils";
import Logo from "@/assets/logo.svg";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import styles from "./index.module.less";
import { Loading } from "../components";
import { GlobalContext } from "../../context";

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

export const Layout = () => {
  const navigate = useNavigate();
  const params = useParams();
  const global = useContext(GlobalContext);
  const result = useRequest("@getCurrentUser");
  const { t } = useTranslation();
  const [collapsed, setCollapsed] = useState(false);

  const curKey = params?.["*"] || global.app?.home || global.app?.menus[0]?.key;
  const defaultSelectedKeys = [curKey];
  const defaultOpenKeys = getOpenKeysFromMenuData(curKey, global.app?.menus);
  const onClickMenuItem = (key: string) => {
    NProgress.start();
    navigate(global.app?.entry + `/${key}`);
    NProgress.done();
  };

  useEffect(() => {
    if (!isEmpty(result.data)) {
      global.setCurrentUser?.(result.data);
    }
  }, [result.data]);

  if (result.loading) {
    return <Loading />;
  }
  if (result.error) {
    return <Navigate to={global.app?.entry + "/sign"} />;
  }

  return (
    <ArcoLayout className={styles.layout}>
      <div className={styles["layout-navbar"]}>
        <div className={styles.navbar}>
          <div className={styles.left}>
            <div className={styles.logo}>
              {global.app?.logo ? <img src={global.app?.logo} /> : <Logo />}
              <div className={styles["logo-name"]}>{global.app?.title}</div>
            </div>
          </div>
          <ul className={styles.right}>
            {global.app?.navitems.map((item: any, i: number) => {
              return (
                <li key={i.toString()}>
                  <SchemaComponent schema={item} scope={{ global, t }} />
                </li>
              );
            })}
          </ul>
        </div>
      </div>
      <ArcoLayout>
        <ArcoLayout.Sider
          className={styles["layout-sider"]}
          collapsed={collapsed}
          onCollapse={setCollapsed}
          trigger={null}
          collapsible
          breakpoint="xl"
          width={collapsed ? 48 : 220}
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
                      defaultSelectedKeys,
                      defaultOpenKeys,
                      menuData: global.app?.menus,
                      collapse: collapsed,
                      onClickMenuItem,
                    },
                  },
                },
              }}
              scope={{ global, t }}
            />
          </div>
          <div
            className={styles["collapse-btn"]}
            onClick={() => setCollapsed((collapsed) => !collapsed)}
          >
            {collapsed ? <IconMenuUnfold /> : <IconMenuFold />}
          </div>
        </ArcoLayout.Sider>
        <ArcoLayout
          className={styles["layout-content"]}
          style={{
            paddingTop: 60,
            paddingLeft: collapsed ? 48 : 220,
          }}
        >
          <div className={styles["layout-content-wrapper"]}>
            <ArcoLayout.Content>
              <RemoteSchemaComponent uid={curKey} />
            </ArcoLayout.Content>
          </div>
        </ArcoLayout>
      </ArcoLayout>
    </ArcoLayout>
  );
};
