import { useEffect, useRef, useState } from "react";
import { Card, Grid, Message } from "@arco-design/web-react";
import LogicFlow from "@logicflow/core";
import ToolbarPanel from "./components/ToolbarPanel";
import NodePanel from "./components/NodePanel";
import { registerNodes } from "./nodes";
import "@logicflow/core/dist/style/index.css";
import PropertyPanel from "./components/PropertyPanel";

export const LiFlow: React.FC = () => {
  const [lf, setLf] = useState<LogicFlow>();
  const [nodeData, setNodeData] = useState();
  const ref = useRef<HTMLDivElement>(null);

  const initEvent = (lf: LogicFlow) => {
    lf.on("blank:click", () => {
      setNodeData(undefined);
    });
    lf.on("element:click", ({ data }) => {
      console.log(data);
      setNodeData(data);
      // console.log(JSON.stringify(lf.getGraphData()));
    });
    lf.on("connection:not-allowed", (data: any) => {
      Message.error(data.msg);
    });
  };

  useEffect(() => {
    if (!ref.current) return;
    const lf = new LogicFlow({
      stopScrollGraph: true,
      stopZoomGraph: true,
      nodeTextEdit: false,
      edgeTextEdit: false,
      grid: {
        size: 10,
        visible: true,
        type: "mesh",
        config: {
          color: "#DCDCDC", // 设置网格的颜色
        },
      },
      keyboard: { enabled: true },
      container: ref.current,
    });
    setLf(lf);
    registerNodes(lf);
    lf.render();
    initEvent(lf);
  }, []);

  return (
    <Card title={lf && <ToolbarPanel lf={lf} />} bodyStyle={{ padding: 0 }}>
      <Grid.Row>
        <Grid.Col span={2} style={{ height: 400, overflowY: "scroll" }}>
          {lf && <NodePanel lf={lf} />}
        </Grid.Col>
        <Grid.Col span={22} style={{ height: 400 }}>
          <div ref={ref} style={{ height: "100%" }} />
          {lf && nodeData && (
            <Card
              style={{
                position: "absolute",
                right: 0,
                top: 0,
                height: "calc(100% - 16px)",
                margin: "8px 8px -8px -8px",
                zIndex: 101,
                boxShadow: "0 0 10px 1px #e4e0db",
              }}
            >
              <PropertyPanel
                activeNode={nodeData}
                onChange={(id, values) => {
                  const node = lf.graphModel.nodesMap[id];
                  const edge = lf.graphModel.edgesMap[id];
                  if (node) {
                    node.model.setProperties(
                      Object.assign(node.model.properties, values)
                    );
                    node.model.updateText(values.name);
                  } else if (edge) {
                    edge.model.setProperties(
                      Object.assign(edge.model.properties, values)
                    );
                    edge.model.updateText(values.name);
                  }
                }}
              />
            </Card>
          )}
        </Grid.Col>
      </Grid.Row>
    </Card>
  );
};

export default LiFlow;
