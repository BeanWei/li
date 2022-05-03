import { BaseNodeModel, CircleNode, CircleNodeModel, h } from "@logicflow/core";
import { node } from "../config";

class EndEventModel extends CircleNodeModel {
  static extendKey = "EndEventModel";
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
    const rules = super.getConnectedSourceRules();
    const notAsSource = {
      message: "结束节点只能连入，不能连出",
      validate: (source: BaseNodeModel, target: BaseNodeModel) => {
        let isValid = true;
        if (target) {
          isValid = false;
        }
        return isValid;
      },
    };
    // @ts-ignore
    rules.push(notAsSource);
    return rules;
  }
}

class EndEventView extends CircleNode {
  static extendKey = "EndEventNode";
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
        href: node.endevent.imgsrc,
      }),
    ]);
  }
}

const EndEvent = {
  type: node.endevent.type,
  view: EndEventView,
  model: EndEventModel,
};

export { EndEventModel, EndEventView };
export default EndEvent;
