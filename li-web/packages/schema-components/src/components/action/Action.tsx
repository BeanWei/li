import { useState } from "react";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import ActionContainer from "./Action.Container";
import { ActionDrawer } from "./Action.Drawer";
import { ActionModal } from "./Action.Modal";
import { ActionPage } from "./Action.Page";
import ActionPopover from "./Action.Popover";
import ActionLink from "./Action.Link";
import { ActionContext } from "./context";
import { ComposedAction } from "./types";
import { Button, Modal } from "@arco-design/web-react";
import { useA } from "./hooks";

// TODO: Improve Typing
export const Action: ComposedAction = observer((props: any) => {
  const {
    popover,
    confirm,
    openMode,
    containerRefKey,
    useAction = useA,
    onClick,
    ...rest
  } = props;
  const [visible, setVisible] = useState(false);
  const field = useField();
  const { run } = useAction();
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
      {!popover && (
        <Button
          {...rest}
          onClick={(e) => {
            e.preventDefault();
            e.stopPropagation();
            const onOk = () => {
              onClick?.(e);
              setVisible(true);
              run();
            };
            if (confirm) {
              Modal.confirm({
                ...confirm,
                onOk,
              });
            } else {
              onOk();
            }
          }}
        >
          {field.title}
        </Button>
      )}
      {!popover && props.children}
    </ActionContext.Provider>
  );
});

Action.Page = ActionPage;
Action.Container = ActionContainer;
Action.Drawer = ActionDrawer;
Action.Modal = ActionModal;
Action.Popover = ActionPopover;
Action.Link = ActionLink;

export default Action;
