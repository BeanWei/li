import React, { useContext } from "react";
import { observer } from "@formily/react";
import { IconLanguage } from "@arco-design/web-react/icon";
import { Button, Select, TriggerProps } from "@arco-design/web-react";
import { GlobalContext } from "../../../context";

export const LangSwitch: React.FC<{
  triggerElementStyle?: React.CSSProperties;
  triggerProps?: TriggerProps;
}> = observer((props) => {
  const { lang, setLang } = useContext(GlobalContext);

  return (
    <Select
      triggerElement={
        <Button
          shape="circle"
          type="secondary"
          style={{
            ...props.triggerElementStyle,
            fontSize: "16px",
          }}
          icon={<IconLanguage />}
        />
      }
      options={[
        { label: "中文", value: "zh-CN" },
        { label: "English", value: "en-US" },
      ]}
      value={lang}
      triggerProps={{
        ...props.triggerProps,
        autoAlignPopupWidth: false,
        autoAlignPopupMinWidth: true,
      }}
      trigger="hover"
      onChange={(value) => {
        setLang?.(value);
      }}
    />
  );
});

export default LangSwitch;
