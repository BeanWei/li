import LogicFlow from "@logicflow/core";
import EndEvent from "./EndEvent";
import SequenceFlow from "./SequenceFlow";
import StartEvent from "./StartEvent";
import ApproverTask from "./ApproverTask";
import ExclusiveGateway from "./ExclusiveGateway";
import WebhookTask from "./WebhookTask";
import { eletype } from "../config";

export const registerNodes = (lf: LogicFlow) => {
  lf.register(StartEvent);
  lf.register(EndEvent);
  lf.register(ApproverTask);
  lf.register(WebhookTask);
  lf.register(SequenceFlow);
  lf.register(ExclusiveGateway);
  lf.setDefaultEdgeType(eletype.sequenceflow);
};
