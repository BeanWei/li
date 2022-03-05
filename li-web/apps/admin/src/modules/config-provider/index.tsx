import {
  ConfigProvider as ArcoConfigProvider,
  ConfigProviderProps,
  Spin,
} from "@arco-design/web-react";
import { useTranslation } from "react-i18next";
import zhCN from "@arco-design/web-react/es/locale/zh-CN";
import enUS from "@arco-design/web-react/es/locale/en-US";
import { useRequest } from "pro-utils";

export function ConfigProvider(props: ConfigProviderProps) {
  const { i18n } = useTranslation();
  const { loading } = useRequest("app:getLang", {
    onSuccess(data) {
      const locale = localStorage.getItem("locale");
      if (data?.data?.lang && locale !== data?.data?.lang) {
        i18n.changeLanguage(data?.data?.lang);
      }
    },
    manual: true, // !remoteLocale,
  });
  if (loading) {
    return <Spin />;
  }
  return (
    <ArcoConfigProvider
      {...props}
      locale={i18n.language === "zh-CN" ? zhCN : enUS}
    >
      {props.children}
    </ArcoConfigProvider>
  );
}
