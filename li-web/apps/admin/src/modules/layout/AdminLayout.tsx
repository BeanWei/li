import { Layout } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useRoute } from "../route-switch/hooks";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import Navbar from "./NavBar";
import { AdminLayoutProvider } from "./AdminLayoutProvider";
import styles from "@/style/layout.module.less";

const InternalLayout = () => {
  const route = useRoute();
  const history = useHistory();
  const match = useRouteMatch<any>();
  const curKey =
    match.params.name ??
    route?.["schema"]?.["x-component-props"]?.["defaultSelectedKeys"]?.[0];
  const onClickMenuItem = (key: string) => {
    history.push(`/admin/${key}`);
  };

  return (
    <Layout className={styles.layout}>
      <div className={styles["layout-navbar"]}>
        <Navbar />
      </div>
      <Layout>
        <Layout.Sider
          className={styles["layout-sider"]}
          trigger={null}
          collapsible
          breakpoint="xl"
          width={220}
          style={{ paddingTop: 60 }}
        >
          <div className={styles["menu-wrapper"]}>
            <SchemaComponent
              schema={route.schema}
              scope={{ onClickMenuItem }}
            />
          </div>
        </Layout.Sider>
        <Layout
          className={styles["layout-content"]}
          style={{ paddingLeft: 220, paddingTop: 60 }}
        >
          <div className={styles["layout-content-wrapper"]}>
            <Layout.Content>
              <RemoteSchemaComponent uid={curKey} />
            </Layout.Content>
          </div>
        </Layout>
      </Layout>
    </Layout>
  );
};

export const AdminLayout = () => {
  return (
    <AdminLayoutProvider>
      <InternalLayout />
    </AdminLayoutProvider>
  );
};

export default AdminLayout;
