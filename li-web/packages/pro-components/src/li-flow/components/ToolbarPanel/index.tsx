import { Button, Divider, Space, Tooltip } from "@arco-design/web-react";
import {
  IconRedo,
  IconUndo,
  IconZoomIn,
  IconZoomOut,
  IconFullscreenExit,
} from "@arco-design/web-react/icon";
import LogicFlow from "@logicflow/core";
import { useEffect, useState } from "react";
import useLocale from "../../../locale";

const ToolbarPanel: React.FC<{
  lf: LogicFlow;
}> = (props) => {
  const local = useLocale();
  const [lfData, setLfData] = useState<Record<string, boolean>>();

  useEffect(() => {
    props.lf.on("history:change", ({ data }) => {
      setLfData(data);
    });
  }, []);

  return (
    <Space>
      <Tooltip
        position="bottom"
        trigger="hover"
        content={local.LiFlow.tooltipUndo}
      >
        <Button
          style={
            lfData?.undoAble
              ? { color: "var(--color-text-2)" }
              : { color: "var(--color-text-4)" }
          }
          icon={<IconUndo />}
          type="text"
          iconOnly
          disabled={!!!lfData?.undoAble}
          onClick={() => {
            props.lf.undo();
          }}
        />
      </Tooltip>
      <Tooltip
        position="bottom"
        trigger="hover"
        content={local.LiFlow.tooltipRedo}
      >
        <Button
          style={
            lfData?.redoAble
              ? { color: "var(--color-text-2)" }
              : { color: "var(--color-text-4)" }
          }
          icon={<IconRedo />}
          type="text"
          iconOnly
          disabled={!!!lfData?.redoAble}
          onClick={() => {
            props.lf.redo();
          }}
        />
      </Tooltip>
      <Divider type="vertical" />
      <Tooltip
        position="bottom"
        trigger="hover"
        content={local.LiFlow.tooltipZoomIn}
      >
        <Button
          style={{ color: "var(--color-text-2)" }}
          icon={<IconZoomIn />}
          type="text"
          iconOnly
          onClick={() => {
            props.lf.zoom(true);
          }}
        />
      </Tooltip>
      <Tooltip
        position="bottom"
        trigger="hover"
        content={local.LiFlow.tooltipZoomOut}
      >
        <Button
          style={{ color: "var(--color-text-2)" }}
          icon={<IconZoomOut />}
          type="text"
          iconOnly
          onClick={() => {
            props.lf.zoom(false);
          }}
        />
      </Tooltip>
      <Tooltip
        position="bottom"
        trigger="hover"
        content={local.LiFlow.tooltipAutoFit}
      >
        <Button
          style={{ color: "var(--color-text-2)" }}
          icon={<IconFullscreenExit />}
          type="text"
          iconOnly
          onClick={() => {
            props.lf.resetZoom();
          }}
        />
      </Tooltip>
    </Space>
  );
};

export default ToolbarPanel;
