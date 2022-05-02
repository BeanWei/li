import { Space } from "@arco-design/web-react";
import LogicFlow from "@logicflow/core";
import useLocale from "../../../locale";
import { nodes } from "../../config";

const NodePanel: React.FC<{ lf: LogicFlow }> = (props) => {
  const local = useLocale();

  return (
    <Space>
      {nodes.map((node, key) => {
        return (
          <Space key={key} direction="vertical" align="center">
            <img
              src={node.src}
              style={node.style}
              onMouseDown={() => {
                props.lf.dnd.startDrag({
                  type: node.type,
                });
              }}
            />
            <span>{local.LiFlow[node.label]}</span>
          </Space>
        );
      })}
    </Space>
  );
};

export default NodePanel;
