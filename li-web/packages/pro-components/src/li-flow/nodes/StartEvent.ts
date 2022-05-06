import { BaseNodeModel, CircleNode, CircleNodeModel, h } from "@logicflow/core";
import { eletype, node } from "../config";

class StartEventModel extends CircleNodeModel {
  static extendKey = "StartEventModel";
  constructor(data: any, graphModel: any) {
    if (!data.text) {
      data.text = "";
    }
    if (typeof data.text === "string") {
      data.text = {
        value: data.text,
        x: data.x,
        y: data.y + 40,
      };
    }
    super(data, graphModel);
  }
  setAttributes() {
    this.r = 21;
  }
  getConnectedTargetRules() {
    const rules = super.getConnectedTargetRules();
    const notAsTarget = {
      message: "开始节点只能连出，不能连入",
      validate: (source: BaseNodeModel, target: BaseNodeModel) => {
        let isValid = true;
        if (target) {
          isValid = false;
        }
        return isValid;
      },
    };
    // @ts-ignore
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
      h("image", {
        ...style,
        x: model.x - (1 / 2) * model.width,
        y: model.y - (1 / 2) * model.height,
        width: model.width,
        height: model.height,
        // 根据宽高缩放
        preserveAspectRatio: "none meet",
        href: node.startevent.imgsrc,
      }),
    ]);
  }
}

const StartEvent = {
  type: eletype.startevent,
  view: StartEventView,
  model: StartEventModel,
};

export { StartEventModel, StartEventView };
export default StartEvent;
