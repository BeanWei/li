import { useContext } from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import { IconMoonFill, IconSunFill } from "@arco-design/web-react/icon";
import { observer } from "@formily/react";
import { LayoutContext } from "../../layout";

export const ThemeSwitch: React.FC<ButtonProps> = observer((props) => {
  const { theme, setTheme } = useContext(LayoutContext);

  return (
    <Button
      shape="circle"
      type="secondary"
      style={{
        fontSize: "16px",
      }}
      {...props}
      icon={theme === "dark" ? <IconMoonFill /> : <IconSunFill />}
      onClick={() => setTheme?.(theme === "light" ? "dark" : "light")}
    />
  );
});

export default ThemeSwitch;
