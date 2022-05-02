import LogicFlow from "@logicflow/core";
import StartEvent from "./StartEvent";

export const registerNodes = (lf: LogicFlow) => {
  lf.register(StartEvent);
};
