import { Tooltip } from "@arco-design/web-react";
import { IconMoonFill, IconSunFill } from "@arco-design/web-react/icon";
import { useContext } from "react";
import { AdminLayoutContext } from "../../AdminLayoutProvider";
import IconButton from "../IconButton";

const ActionTheme: React.FC = () => {
  const { theme, setTheme } = useContext(AdminLayoutContext);

  return (
    <Tooltip
      content={theme === "light" ? "点击切换为暗黑模式" : "点击切换为亮色模式"}
    >
      <IconButton
        icon={theme !== "dark" ? <IconMoonFill /> : <IconSunFill />}
        onClick={() => setTheme?.(theme === "light" ? "dark" : "light")}
      />
    </Tooltip>
  );
};

export default ActionTheme;
