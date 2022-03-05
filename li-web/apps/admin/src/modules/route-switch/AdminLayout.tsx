import { useRef, useState } from "react";
import { css } from "@emotion/css";
import { Layout } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { SchemaComponent } from "schema-components";
import { useRoute } from "./hooks";
import { useDocumentTitle } from "../document-title";
import { CurrentUser, CurrentUserProvider } from "../user";
import { RemoteSchemaComponent } from "./RemoteSchemaComponent";

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
  // const result = useSystemSettings();
  return (
    <Layout>
      <Layout.Header
        style={{
          height: 46,
          lineHeight: "46px",
          position: "relative",
          paddingLeft: 0,
        }}
      >
        <div style={{ display: "flex", height: "100%" }}>
          <div
            style={{
              width: 200,
              display: "inline-flex",
              color: "#fff",
              padding: "0",
              alignItems: "center",
            }}
          >
            <img
              className={css`
                height: 20px;
                padding: 0 16px;
              `}
              // src={result?.data?.data?.logo?.url}
              src="https://unpkg.byted-static.com/latest/byted/arco-config/assets/favicon.ico"
            />
            {/* {result?.data?.data?.title} */}
          </div>
          <SchemaComponent
            hidden={hidden}
            schema={route.schema}
            scope={{ onSelect, sideMenuRef, defaultSelectedUid }}
          />
        </div>
        <div style={{ position: "absolute", top: 0, right: 0 }}>
          {/* <PluginManager.Toolbar
            items={[
              { component: 'DesignableSwitch', pin: true },
              { component: 'CollectionManagerShortcut', pin: true },
              { component: 'ACLShortcut', pin: true },
              { component: 'SystemSettingsShortcut' },
            ]}
          /> */}
          <CurrentUser />
        </div>
      </Layout.Header>
      <Layout>
        <Layout.Sider
          style={{ display: "none" }}
          theme={"light"}
          ref={sideMenuRef}
        ></Layout.Sider>
        <Layout.Content style={{ minHeight: "calc(100vh - 46px)" }}>
          <RemoteSchemaComponent uid={match.params.name} />
        </Layout.Content>
      </Layout>
    </Layout>
  );
};

export const AdminLayout = () => {
  return (
    <CurrentUserProvider>
      <InternalLayout />
    </CurrentUserProvider>
  );
};

export default AdminLayout;
