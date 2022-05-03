import LogicFlow from "@logicflow/core";

// 将 LogicFlow 中的 Node 数据转换为 LiFlow 元素数据
function convertNodeToLiFlowElement(node: any) {
  const { id, x, y, text = "", properties } = node;
  const [type, flow_type] = node.type.split(":");
  return {
    incoming: [],
    outgoing: [],
    type,
    flow_type,
    properties: {
      ...properties,
      name: (text && text.value) || "",
      x,
      y,
      text,
    },
    key: id,
  };
}

// 将 LogicFlow 中的 Edge 数据转换为 LiFlow 元素数据
function convertEdgeToLiFlowElement(edge: any) {
  const {
    id,
    sourceNodeId,
    targetNodeId,
    startPoint,
    endPoint,
    text = "",
    properties,
  } = edge;
  const [type, flow_type] = edge.type.split(":");
  return {
    incoming: [sourceNodeId],
    outgoing: [targetNodeId],
    type,
    flow_type,
    properties: {
      ...properties,
      name: (text && text.value) || "",
      text,
      startPoint: JSON.stringify(startPoint),
      endPoint: JSON.stringify(endPoint),
      pointsList: JSON.stringify(text),
    },
    key: id,
  };
}

// 将 LogicFlow 中数据转换为 LiFlow 数据
export function toLiFlowModel(data: any) {
  const nodeMap = new Map();
  const flowElementList: any[] = [];
  data.nodes.forEach((node: any) => {
    const flowElement = convertNodeToLiFlowElement(node);
    flowElementList.push(flowElement);
    nodeMap.set(node.id, flowElement);
  });
  data.edges.forEach((edge: any) => {
    const flowElement = convertEdgeToLiFlowElement(edge);
    const sourceElement = nodeMap.get(edge.sourceNodeId);
    sourceElement.outgoing.push(flowElement.key);
    const targetElement = nodeMap.get(edge.targetNodeId);
    targetElement.incoming.push(flowElement.key);
    flowElementList.push(flowElement);
  });
  return flowElementList;
}

// 将 LiFlow 元素数据转换为 LogicFlow 中的 Edge 数据
function convertFlowElementToEdge(element: any) {
  const { incoming, outgoing, properties, key, type, flow_type } = element;
  const { text, name, startPoint, endPoint, pointsList } = properties;
  const edge: any = {
    id: key,
    type: `${type}:${flow_type}`,
    sourceNodeId: incoming[0],
    targetNodeId: outgoing[0],
    text: text || name,
    properties: {},
  };
  if (startPoint) {
    // @ts-ignore
    edge.startPoint = JSON.parse(startPoint);
  }
  if (endPoint) {
    // @ts-ignore
    edge.endPoint = JSON.parse(endPoint);
  }
  if (pointsList) {
    // @ts-ignore
    edge.endPoint = JSON.parse(pointsList);
  }
  // 这种转换方式，在自定义属性中不能与excludeProperties中的属性重名，否则将在转换过程中丢失
  const excludeProperties = ["startPoint", "endPoint", "pointsList", "text"];
  Object.keys(element.properties).forEach((property) => {
    if (excludeProperties.indexOf(property) === -1) {
      edge.properties[property] = element.properties[property];
    }
  });
  return edge;
}

// 将 LiFlow 元素数据转换为 LogicFlow 中的 Node 数据
function convertFlowElementToNode(element: any) {
  const { properties, key, type, flow_type, bounds } = element;
  let { x, y } = properties;
  if (x === undefined) {
    const [{ x: x1, y: y1 }, { x: x2, y: y2 }] = bounds;
    x = (x1 + x2) / 2;
    y = (y1 + y2) / 2;
  }
  const node: any = {
    id: key,
    type: `${type}:${flow_type}`,
    x,
    y,
    text: properties.text,
    properties: {},
  };
  // 这种转换方式，在自定义属性中不能与excludeProperties中的属性重名，否则将在转换过程中丢失
  const excludeProperties = ["x", "y", "text"];
  Object.keys(element.properties).forEach((property) => {
    if (excludeProperties.indexOf(property) === -1) {
      node.properties[property] = element.properties[property];
    }
  });
  return node;
}

// 将 LiFlow 元素数据转换为 LogicFlow 数据
export function toLogicflowData(data: any) {
  const lfData: any = {
    nodes: [],
    edges: [],
  };
  Array.isArray(data) &&
    data.length > 0 &&
    data.forEach((element: any) => {
      if (element.flow_type === "SequenceFlow") {
        const edge = convertFlowElementToEdge(element);
        lfData.edges.push(edge);
      } else {
        const node = convertFlowElementToNode(element);
        lfData.nodes.push(node);
      }
    });
  return lfData;
}

class LiFlowAdapter {
  static pluginName = "liflowAdapter";
  constructor({ lf }: { lf: LogicFlow }) {
    lf.adapterIn = this.adapterIn;
    lf.adapterOut = this.adapterOut;
  }
  adapterOut(logicflowData: any) {
    if (logicflowData) {
      return toLiFlowModel(logicflowData);
    }
  }
  adapterIn(liflowData: any) {
    if (liflowData) {
      return toLogicflowData(liflowData);
    }
  }
}

export default LiFlowAdapter;
