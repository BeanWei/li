import { Fragment } from "react";
import ActionFormDrawer from "./Action.FormDrawer";
import ActionFormModal from "./Action.FormModal";
import { ComposedAction } from "./types";

export const Action: ComposedAction = (props) => {
  return <Fragment />;
};

Action.FormDrawer = ActionFormDrawer;
Action.FormModal = ActionFormModal;

export default Action;
