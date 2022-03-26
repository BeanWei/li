import {
  ConfigProvider as ArcoConfigProvider,
  ConfigProviderProps,
} from "@arco-design/web-react";
import { useTranslation } from "react-i18next";
import zhCN from "@arco-design/web-react/es/locale/zh-CN";
import enUS from "@arco-design/web-react/es/locale/en-US";

export function ConfigProvider(props: ConfigProviderProps) {
  const { i18n } = useTranslation();
  return (
    <ArcoConfigProvider
      {...props}
      locale={i18n.language === "zh-CN" ? zhCN : enUS}
    >
      {props.children}
    </ArcoConfigProvider>
  );
}
