import { useState } from "react";
import { Layout } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useRoute } from "../route-switch/hooks";
import { useDocumentTitle } from "../document-title";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import Navbar from "./NavBar";
import { AdminLayoutProvider } from "./AdminLayoutProvider";
import styles from "@/style/layout.module.less";

const InternalLayout = () => {
  const route = useRoute();
  const history = useHistory();
  const match = useRouteMatch<any>();
  const { setTitle } = useDocumentTitle();
  const defaultSelectedKeys = match.params.name ?? [];
  const [, setSchema] = useState({});
  const onClickMenuItem = (key: string, _: any, paths: any) => {
    console.log(key, paths);
    // const schema = item.props.schema;
    // console.log("onSelect", schema);
    // setSchema(schema);
    // setTitle?.(schema.title);
    // history.push(`/admin/${schema["x-data"]}`);
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
              scope={{ onClickMenuItem, defaultSelectedKeys }}
            />
          </div>
        </Layout.Sider>
        <Layout className={styles["layout-content"]}>
          <div className={styles["layout-content-wrapper"]}>
            <Layout.Content>
              <RemoteSchemaComponent uid={match.params.name} />
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
