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
  Layout,
  RouteSwitch,
  RouteSwitchProvider,
  ThemeSwitch,
} from "./modules";

const providers = [
  [I18nextProvider, { i18n }],
  [ConfigProvider, { remoteLocale: true }],
  [SchemaComponentProvider, { components: { Link, NavLink, ThemeSwitch } }],
  UiSchemaComponentProvider,
  [DocumentTitleProvider, { addonAfter: "Li" }],
  [RouteSwitchProvider, { components: { Layout } }],
];

const App = compose(...providers)(() => {
  const { data, loading } = useRequest("@getAppConfig");

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
            component: "Layout",
            title: data?.data?.title || "Li Admin",
            config: data?.data,
          },
          {
            type: "route",
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
