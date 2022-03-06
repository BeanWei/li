import { useContext } from "react";
import { ConfigProvider } from "@arco-design/web-react";

export const getPrefixCls = () => {
  const { prefixCls } = useContext(ConfigProvider.ConfigContext);
  return prefixCls;
};
