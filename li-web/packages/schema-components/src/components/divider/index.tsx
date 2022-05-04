import { Divider as ArcoDivider } from "@arco-design/web-react";
import React from "react";

export const Divider: React.FC<React.PropsWithChildren<{}>> = (props) => {
  return (
    <ArcoDivider style={{ marginBottom: 24 }} {...props}>
      {props.children}
    </ArcoDivider>
  );
};

export default Divider;
