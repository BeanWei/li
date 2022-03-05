import { Link, Spin } from "@arco-design/web-react";
import { I18nextProvider } from "react-i18next";
import { NavLink } from "react-router-dom";
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
        Layout,
      },
    },
  ],
];

const App = compose(...providers)(() => {
  const { data, loading } = useRequest({
    operation: "getAppRoutes",
  });
  if (loading) {
    return <Spin />;
  }
  return (
    <div>
      <RouteSwitch routes={data?.data || []} />
    </div>
  );
});

export default App;
