import { useRef, useState } from "react";
import { Layout } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useRoute } from "../route-switch/hooks";
import { useDocumentTitle } from "../document-title";
import { RemoteSchemaComponent } from "../route-switch/RemoteSchemaComponent";
import Navbar from "./NavBar";
import { AdminLayoutProvider } from "./AdminLayoutProvider";
import styles from "@/style/layout.module.less";

const InternalLayout = (props: any) => {
  const route = useRoute();
  const history = useHistory();
  const match = useRouteMatch<any>();
  const { setTitle } = useDocumentTitle();
  const sideMenuRef = useRef();
  const defaultSelectedUid = match.params.name;
  const [schema, setSchema] = useState({});
  const onSelect = ({ item }: any) => {
    const schema = item.props.schema;
    console.log("onSelect", schema);
    setSchema(schema);
    setTitle?.(schema.title);
    history.push(`/admin/${schema["x-data"]}`);
  };
  const [hidden, setHidden] = useState(false);

  return (
    <Layout className={styles.layout}>
      <div className={styles["layout-navbar"]}>
        <Navbar />
      </div>
      <Layout>
        <Layout.Sider
          style={{ display: "none" }}
          theme={"light"}
          ref={sideMenuRef}
        >
          <SchemaComponent
            hidden={hidden}
            schema={route.schema}
            scope={{ onSelect, sideMenuRef, defaultSelectedUid }}
          />
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
