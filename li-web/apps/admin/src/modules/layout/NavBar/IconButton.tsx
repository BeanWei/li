import React, { forwardRef } from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import styles from "./style/icon-button.module.less";
import cs from "classnames";

function IconButton(props: ButtonProps, ref: React.Ref<unknown> | undefined) {
  const { icon, className, ...rest } = props;

  return (
    <Button
      ref={ref}
      icon={icon}
      shape="circle"
      type="secondary"
      className={cs(styles["icon-button"], className)}
      {...rest}
    />
  );
}

export default forwardRef(IconButton);
