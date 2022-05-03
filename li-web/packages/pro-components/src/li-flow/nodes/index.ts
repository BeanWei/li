import LogicFlow from "@logicflow/core";
import EndEvent from "./EndEvent";
import SequenceFlow from "./SequenceFlow";
import StartEvent from "./StartEvent";
import UserTask from "./UserTask";
import ExclusiveGateway from "./ExclusiveGateway";

export const registerNodes = (lf: LogicFlow) => {
  lf.register(StartEvent);
  lf.register(EndEvent);
  lf.register(UserTask);
  lf.register(SequenceFlow);
  lf.register(ExclusiveGateway);
  lf.setDefaultEdgeType("SequenceFlow:SequenceFlow");
};
