import { PolylineEdge, PolylineEdgeModel } from "@logicflow/core";

class SequenceFlowModel extends PolylineEdgeModel {
  static extendKey = "SequenceFlowModel";
  getEdgeStyle() {
    const style = super.getEdgeStyle();
    // svg属性
    style.strokeWidth = 1;
    style.stroke = "#ababac";
    return style;
  }
  /**
   * 重写此方法，使保存数据是能带上锚点数据。
   */
  // getData() {
  //   const data = super.getData();
  //   data.sourceAnchorId = this.sourceAnchorId;
  //   data.targetAnchorId = this.targetAnchorId;
  //   return data;
  // }
}

class SequenceFlowView extends PolylineEdge {
  static extendKey = "SequenceFlowEdge";
}

const SequenceFlow = {
  type: "SequenceFlow:SequenceFlow",
  view: SequenceFlowView,
  model: SequenceFlowModel,
};

export { SequenceFlowView, SequenceFlowModel };
export default SequenceFlow;
