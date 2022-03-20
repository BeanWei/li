import React from "react";
import * as ICONS from "@arco-design/web-react/icon";
import { IconProps } from "@arco-design/web-react/icon";

const allicons = ICONS as Record<string, React.FC>;

export const Icon: React.FC<IconProps & { type?: string }> = (props) => {
  const { type, ...rest } = props;
  if (!type) {
    return null;
  }
  const Icon = allicons[type];
  if (!Icon) {
    return null;
  }
  return <Icon {...rest} />;
};

export default Icon;
