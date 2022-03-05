import { useRef, useState } from "react";
import { css } from "@emotion/css";
import { Layout as ArcoLayout } from "@arco-design/web-react";
import { useHistory, useRouteMatch } from "react-router-dom";
import { Schema } from "@formily/react";
import { useRoute } from "./hooks";
import { useDocumentTitle } from "../document-title";
import { CurrentUser, CurrentUserProvider } from "../user";
import { RemoteSchemaComponent } from "./RemoteSchemaComponent";

export function findByUid(schema: Schema, uid: string): any {
  if (!Schema.isSchemaInstance(schema)) {
    schema = new Schema(schema);
  }
  return schema.reduceProperties((buffter, s) => {
    if (s["x-uid"] === uid) {
      return s;
    }
    const ss = findByUid(s, uid);
    if (ss) {
      return ss;
    }
    return buffter;
  }, null);
}

export function findMenuItem(schema: Schema): any {
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

export function findKeysByUid(schema: Schema, uid: string) {
  return findKeys(findByUid(schema, uid));
}

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
    history.push(`/admin/${schema["x-uid"]}`);
  };
  const [hidden, setHidden] = useState(false);
  // const result = useSystemSettings();
  return (
    <ArcoLayout>
      <ArcoLayout.Header
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
          <RemoteSchemaComponent
            hidden={hidden}
            uid={route.uiSchemaUid}
            scope={{ onSelect, sideMenuRef, defaultSelectedUid }}
            schemaTransform={(data: any) => {
              if (!data) {
                return data;
              }
              data["x-component-props"] = data["x-component-props"] || {};
              data["x-component-props"]["defaultSelectedUid"] =
                defaultSelectedUid;
              return data;
            }}
            onSuccess={(data: any) => {
              if (defaultSelectedUid) {
                const s = findByUid(data?.data, defaultSelectedUid);
                if (s) {
                  setTitle?.(s.title);
                }
                return;
              }
              setHidden(true);
              setTimeout(() => setHidden(false), 11);
              const s = findMenuItem(data?.data);
              if (s) {
                setSchema(s);
                setTitle?.(s.title);
                history.push(`/admin/${s["x-uid"]}`);
              }
            }}
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
      </ArcoLayout.Header>
      <ArcoLayout>
        <ArcoLayout.Sider
          style={{ display: "none" }}
          theme={"light"}
          ref={sideMenuRef}
        ></ArcoLayout.Sider>
        <ArcoLayout.Content style={{ minHeight: "calc(100vh - 46px)" }}>
          <RemoteSchemaComponent onlyRenderProperties uid={match.params.name} />
        </ArcoLayout.Content>
      </ArcoLayout>
    </ArcoLayout>
  );
};

export const Layout = () => {
  return (
    <CurrentUserProvider>
      <InternalLayout />
    </CurrentUserProvider>
  );
};

export default Layout;
