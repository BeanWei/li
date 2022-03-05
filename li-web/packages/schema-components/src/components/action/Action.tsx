import { useState } from "react";
import { css } from "@emotion/css";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { Popover } from "@arco-design/web-react";
import ActionContainer from "./Action.Container";
import { ActionDrawer } from "./Action.Drawer";
import { ActionModal } from "./Action.Modal";
import { ActionPage } from "./Action.Page";
import { useActionContext } from "./hooks";
import { ActionContext } from "./context";
import { ComposedAction } from "./types";

export const Action: ComposedAction = observer((props: any) => {
  const { popover, openMode, containerRefKey } = props;
  const [visible, setVisible] = useState(false);
  const field = useField();
  const fieldSchema = useFieldSchema();
  return (
    <ActionContext.Provider
      value={{ visible, setVisible, openMode, containerRefKey }}
    >
      {popover && (
        <RecursionField
          basePath={field.address}
          onlyRenderProperties
          schema={fieldSchema}
        />
      )}
      {!popover && props.children}
    </ActionContext.Provider>
  );
});

Action.Popover = observer((props) => {
  const { button, visible, setVisible } = useActionContext();
  return (
    <Popover
      {...props}
      popupVisible={visible}
      onVisibleChange={(visible) => {
        setVisible(visible);
      }}
      content={props.children}
    >
      {button}
    </Popover>
  );
});

Action.Popover.Footer = observer((props) => {
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

Action.Link = observer((props) => {
  return <Action {...props} component={"a"} className={"li-action-link"} />;
});

Action.Drawer = ActionDrawer;
Action.Modal = ActionModal;
Action.Container = ActionContainer;
Action.Page = ActionPage;

export default Action;
