import { CircleNode, CircleNodeModel, h } from "@logicflow/core";
import { nanoid } from "pro-utils";
import { img } from "../config";

class StartEventModel extends CircleNodeModel {
  static extendKey = "StartEventModel";
  constructor(data: any, graphModel: any) {
    if (!data.id) {
      data.id = `StartEvent_${nanoid()}`;
    }
    super(data, graphModel);
  }
  getConnectedTargetRules() {
    const rules = super.getConnectedTargetRules();
    const notAsTarget = {
      message: "起始节点不能作为边的终点",
      validate: () => false,
    };
    rules.push(notAsTarget);
    return rules;
  }
}

class StartEventView extends CircleNode {
  static extendKey = "StartEventNode";
  getShape(): any {
    const { model } = this.props;
    const style = model.getNodeStyle();
    return h("g", {}, [
      h("img", {
        ...style,
        src: img.startevent,
      }),
    ]);
  }
}

const StartEvent = {
  type: "StartEvent:StartEvent",
  view: StartEventView,
  model: StartEventModel,
};

export { StartEventModel, StartEventView };
export default StartEvent;
