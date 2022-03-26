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
  LangSwitch,
  Layout,
  RouteSwitch,
  RouteSwitchProvider,
  SignPage,
  ThemeSwitch,
} from "./modules";
import { Loading } from "./modules/components";

const providers = [
  [I18nextProvider, { i18n }],
  [ConfigProvider, { remoteLocale: true }],
  [
    SchemaComponentProvider,
    { components: { Link, NavLink, LangSwitch, ThemeSwitch } },
  ],
  UiSchemaComponentProvider,
  [DocumentTitleProvider, { addonAfter: "Li" }],
  [RouteSwitchProvider, { components: { SignPage, Layout } }],
];

const App = compose(...providers)(() => {
  const { data, loading } = useRequest("@getAppConfig");
  const entry = data?.entry || "/admin";

  if (loading) {
    return <Loading />;
  }
  return (
    <BrowserRouter>
      <RouteSwitch
        routes={[
          {
            type: "redirect",
            from: "/",
            to: entry,
            exact: true,
          },
          {
            type: "route",
            path: entry + "/sign",
            component: "SignPage",
            title: "Sign",
            config: {
              title: data?.title,
              logo: data?.logo,
              body: data?.binding.signform,
              footer: data?.copyright,
            },
          },
          {
            type: "route",
            path: entry + "/:name(.+)?",
            component: "Layout",
            title: data?.title || "Li Admin",
            config: {
              ...data,
              title: data?.title || "Li Admin",
              entry,
            },
          },
        ]}
      />
    </BrowserRouter>
  );
});

export default App;
