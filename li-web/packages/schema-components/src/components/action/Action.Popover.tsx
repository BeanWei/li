import { Popover } from "@arco-design/web-react";
import { css } from "@emotion/css";
import { observer } from "@formily/reactive-react";
import { useActionContext } from "./hooks";
import { ComposedActionPopover } from "./types";

export const ActionPopover: ComposedActionPopover = observer((props) => {
  const { button, visible, setVisible } = useActionContext();
  return (
    <Popover
      {...props}
      popupVisible={visible}
      onVisibleChange={(visible: boolean) => {
        setVisible(visible);
      }}
      content={props.children}
    >
      {button}
    </Popover>
  );
});

ActionPopover.Footer = observer((props) => {
  return (
    <div
      className={css`
        display: flex;
        justify-content: flex-end;
        width: 100%;
      `}
    >
      {props.children}
    </div>
  );
});

export default ActionPopover;
