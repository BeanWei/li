import { useEffect, useRef, useState } from "react";
import { Card, Grid, Message } from "@arco-design/web-react";
import LogicFlow from "@logicflow/core";
import ToolbarPanel from "./components/ToolbarPanel";
import NodePanel from "./components/NodePanel";
import { registerNodes } from "./nodes";
import "@logicflow/core/dist/style/index.css";

export const LiFlow: React.FC = () => {
  const [lf, setLf] = useState<LogicFlow>();
  const ref = useRef<HTMLDivElement>(null);

  const initEvent = (lf: LogicFlow) => {
    lf.on("element:click", ({ data }) => {
      console.log(JSON.stringify(lf.getGraphData()));
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
        <Grid.Col span={2} style={{ height: 400, overflow: "scroll" }}>
          {lf && <NodePanel lf={lf} />}
        </Grid.Col>
        <Grid.Col span={22} style={{ height: 400 }}>
          <div ref={ref} style={{ height: "100%" }} />
        </Grid.Col>
      </Grid.Row>
    </Card>
  );
};

export default LiFlow;
