import { useContext } from "react";
import { ConfigProvider } from "@arco-design/web-react";
import { Locale } from "../../../locale/types";

export const getLocale = (): Locale => {
  const { locale } = useContext(ConfigProvider.ConfigContext);
  return locale as Locale;
};
