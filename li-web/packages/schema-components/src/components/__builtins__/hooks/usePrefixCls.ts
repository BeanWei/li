import { useContext } from "react";
import { ConfigProvider } from "@arco-design/web-react";

export const usePrefixCls = (
  tag: string,
  props?: {
    prefixCls?: string;
  }
) => {
  const { getPrefixCls } = useContext(ConfigProvider.ConfigContext);
  return getPrefixCls?.(tag, props?.prefixCls);
};
