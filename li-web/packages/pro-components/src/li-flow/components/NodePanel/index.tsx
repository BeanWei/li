import LogicFlow from "@logicflow/core";
import useLocale from "../../../locale";
import { node } from "../../config";

const NodePanel: React.FC<{ lf: LogicFlow }> = (props) => {
  const local = useLocale();
  const nodes = Object.values(node);

  return (
    <div
      style={{ width: "100%", padding: "16px 0" }}
      className="arco-space arco-space-vertical arco-space-align-center"
    >
      {nodes.map((item, key) => {
        return (
          <div
            key={key}
            style={{
              width: "100%",
              marginBottom: key === nodes.length - 1 ? 0 : 8,
            }}
            className="arco-space arco-space-vertical arco-space-align-center"
          >
            <img
              src={item.imgsrc}
              style={{
                ...item.style,
                cursor: "grab",
              }}
              draggable={false}
              onMouseDown={() => {
                props.lf.dnd.startDrag({
                  type: item.type,
                });
              }}
            />
            <span style={{ fontSize: 12 }}>{local.LiFlow[item.label]}</span>
          </div>
        );
      })}
    </div>
  );
};

export default NodePanel;
