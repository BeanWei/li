import React from "react";
import {
  Divider,
  Space as ArcoSpace,
  SpaceProps,
} from "@arco-design/web-react";
import { useFormLayout } from "../form-layout";
import "@arco-design/web-react/lib/Divider/style/index";
import "@arco-design/web-react/lib/Space/style/index";

export const Space: React.FC<SpaceProps> = (props) => {
  let { split } = props;
  if (split === "divider") {
    split = <Divider type="vertical" style={{ margin: "0 2px" }} />;
  }
  const layout = useFormLayout();
  return React.createElement(ArcoSpace, {
    size: props.size ?? layout?.spaceGap,
    ...props,
  });
};

export default Space;