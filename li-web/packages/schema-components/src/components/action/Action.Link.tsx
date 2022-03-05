import { observer } from "@formily/react";
import Action from "./Action";
import { ComposedAction } from "./types";

export const ActionLink: ComposedAction = observer((props: any) => {
  return <Action {...props} component={"a"} className={"li-action-link"} />;
});

export default ActionLink;
