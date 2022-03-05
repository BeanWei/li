import { Switch as ArcoSwitch } from "@arco-design/web-react";
import { connect, mapProps } from "@formily/react";

export const Switch = connect(
  ArcoSwitch,
  mapProps({
    value: "checked",
  })
);

export default Switch;
