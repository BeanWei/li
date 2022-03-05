import { Link, Notification, Spin } from "@arco-design/web-react";
import { I18nextProvider } from "react-i18next";
import { NavLink } from "react-router-dom";
import { SchemaComponentProvider } from "schema-components";
import {
  APIClient,
  APIClientProvider,
  compose,
  DocumentTitleProvider,
  Layout,
  RouteSwitch,
  RouteSwitchProvider,
  useRequest,
} from "./modules";

const apiClient = new APIClient({
  baseURL: `http://localhost:3000/api/`,
});

apiClient.axios.interceptors.response.use(
  (response) => response,
  (error) => {
    Notification.error({
      content: error?.response?.data?.errors?.map?.((error: any) => {
        return <div>{error.message}</div>;
      }),
    });
    throw error;
  }
);

const providers = [
  // [HashRouter],
  // [MemoryRouter, { initialEntries: ['/'] }],
  [APIClientProvider, { apiClient }],
  // [I18nextProvider, { i18n }],
  // [AntdConfigProvider, { remoteLocale: true }],
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
    url: "uiRoutes:getAccessible",
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
