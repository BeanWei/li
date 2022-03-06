import { Link, Spin } from "@arco-design/web-react";
import { I18nextProvider } from "react-i18next";
import { BrowserRouter, NavLink } from "react-router-dom";
import {
  SchemaComponentProvider,
  UiSchemaComponentProvider,
} from "schema-components";
import { useRequest } from "pro-utils";
import {
  compose,
  ConfigProvider,
  DocumentTitleProvider,
  i18n,
  AdminLayout,
  BlankLayout,
  RouteSwitch,
  RouteSwitchProvider,
} from "./modules";

const providers = [
  [I18nextProvider, { i18n }],
  [ConfigProvider, { remoteLocale: true }],
  // SystemSettingsProvider,
  // [
  //   PluginManagerProvider,
  //   {
  //     components: {
  //       ACLShortcut,
  //       DesignableSwitch,
  //       CollectionManagerShortcut,
  //       SystemSettingsShortcut,
  //     },
  //   },
  // ],
  [SchemaComponentProvider, { components: { Link, NavLink } }],
  UiSchemaComponentProvider,
  [DocumentTitleProvider, { addonAfter: "Li" }],
  [
    RouteSwitchProvider,
    {
      components: {
        AdminLayout,
        BlankLayout,
      },
    },
  ],
];

const App = compose(...providers)(() => {
  const { data, loading } = useRequest({
    operation: "getAppMenuSchema",
  });

  if (loading) {
    return <Spin />;
  }
  return (
    <BrowserRouter>
      <RouteSwitch
        routes={[
          {
            type: "redirect",
            from: "/",
            to: "/admin",
            exact: true,
          },
          {
            type: "route",
            path: "/admin/:name(.+)?",
            component: "AdminLayout",
            title: "Li Admin",
            schema: data,
          },
          {
            type: "route",
            component: "BlankLayout",
            routes: [
              {
                type: "route",
                path: "/signin",
                component: "SigninPage",
              },
              {
                type: "route",
                path: "/signup",
                component: "SignupPage",
              },
            ],
          },
        ]}
      />
    </BrowserRouter>
  );
});

export default App;
