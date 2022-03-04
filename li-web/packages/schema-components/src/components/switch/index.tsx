import { Switch as ArcoSwitch } from "@arco-design/web-react";
import { connect, mapProps } from "@formily/react";
import "@arco-design/web-react/lib/Switch/style/index";

export const Switch = connect(
  ArcoSwitch,
  mapProps({
    value: "checked",
  })
);

export default Switch;
