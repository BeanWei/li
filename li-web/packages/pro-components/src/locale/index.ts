import { useContext } from "react";
import { ConfigProvider } from "@arco-design/web-react";
import { Locale } from "./types";
import enUS from "./en-US";
import zhCN from "./zh-CN";

function useLocale(): Locale {
  const lang = useContext(ConfigProvider.ConfigContext).locale?.locale;
  if (lang === "en-US") {
    return enUS;
  }
  return zhCN;
}

export default useLocale;
