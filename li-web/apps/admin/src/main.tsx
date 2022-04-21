import { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import { ConfigProvider, Link } from "@arco-design/web-react";
import { I18nextProvider, useTranslation } from "react-i18next";
import { BrowserRouter, NavLink } from "react-router-dom";
import { useLocalStorageState } from "ahooks";
import {
  SchemaComponentProvider,
  UiSchemaComponentProvider,
  zhCN,
  enUS,
} from "schema-components";
import { useRequest } from "pro-utils";
import {
  compose,
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
import { GlobalContext } from "./context";
import "@arco-design/web-react/dist/css/arco.css";
import "./style/global.less";

const providers = [
  [I18nextProvider, { i18n }],
  [
    SchemaComponentProvider,
    { components: { Link, NavLink, LangSwitch, ThemeSwitch } },
  ],
  UiSchemaComponentProvider,
  [DocumentTitleProvider, { addonAfter: "Li" }],
  [RouteSwitchProvider, { components: { SignPage, Layout } }],
];

const App = compose(...providers)(() => {
  const [theme, setTheme] = useLocalStorageState("li-theme", {
    defaultValue: window.matchMedia?.("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light",
  });
  const [lang, setLang] = useLocalStorageState("li-lang", {
    defaultValue: navigator.language,
  });
  const [currentUser, setCurrentUser] = useState({});
  const { t } = useTranslation();
  const { data, loading } = useRequest("@getAppConfig");
  const entry = data?.entry || "/admin";

  useEffect(() => {
    if (theme === "dark") {
      document.body.setAttribute("arco-theme", "dark");
    } else {
      document.body.removeAttribute("arco-theme");
    }
  }, [theme]);

  if (loading) {
    return <Loading />;
  }

  const contextValue = {
    app: {
      ...data,
      title: t(data?.title || "Li Admin"),
      entry,
    },
    lang,
    setLang,
    theme,
    setTheme,
    currentUser,
    setCurrentUser,
  };

  return (
    <BrowserRouter>
      <ConfigProvider
        locale={lang === "en-US" ? enUS : zhCN}
        componentConfig={{
          Card: {
            bordered: false,
          },
          List: {
            bordered: false,
          },
          Table: {
            border: false,
          },
        }}
      >
        <GlobalContext.Provider value={contextValue}>
          <RouteSwitch
            routes={[
              {
                path: "/",
                redirect: entry,
                title: contextValue.app.title,
              },
              {
                path: entry + "/sign",
                component: "SignPage",
                title: contextValue.app.title,
              },
              {
                path: entry + "/*",
                component: "Layout",
                title: contextValue.app.title,
              },
            ]}
          />
        </GlobalContext.Provider>
      </ConfigProvider>
    </BrowserRouter>
  );
});

ReactDOM.render(<App />, document.getElementById("root"));
