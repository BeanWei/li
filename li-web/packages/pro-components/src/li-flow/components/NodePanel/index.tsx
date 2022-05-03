import LogicFlow from "@logicflow/core";
import useLocale from "../../../locale";
import { node } from "../../config";

const NodePanel: React.FC<{ lf: LogicFlow }> = (props) => {
  const local = useLocale();

  return (
    <div
      style={{ width: "100%", padding: "16px 0" }}
      className="arco-space arco-space-horizontal arco-space-align-center"
    >
      {Object.values(node).map((item, key) => {
        return (
          <div
            key={key}
            style={{ width: "100%" }}
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
            <span>{local.LiFlow[item.label]}</span>
          </div>
        );
      })}
    </div>
  );
};

export default NodePanel;
