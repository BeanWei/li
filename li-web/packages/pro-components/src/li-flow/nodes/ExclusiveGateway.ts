import { h, PolygonNode, PolygonNodeModel } from "@logicflow/core";
import { eletype, node } from "../config";

class ExclusiveGatewayModel extends PolygonNodeModel {
  static extendKey = "ExclusiveGatewayModel";
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
    this.points = [
      [25, 0],
      [50, 25],
      [25, 50],
      [0, 25],
    ];
  }
}

class ExclusiveGatewayView extends PolygonNode {
  static extendKey = "ExclusiveGatewayNode";
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
        href: node.exclusivegateway.imgsrc,
      }),
    ]);
  }
}

const ExclusiveGateway = {
  type: eletype.exclusivegateway,
  view: ExclusiveGatewayView,
  model: ExclusiveGatewayModel,
};

export { ExclusiveGatewayView, ExclusiveGatewayModel };
export default ExclusiveGateway;
