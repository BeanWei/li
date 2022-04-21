import { useContext } from "react";
import { observer } from "@formily/react";
import { IconLanguage } from "@arco-design/web-react/icon";
import { Button, Select } from "@arco-design/web-react";
import { GlobalContext } from "../../../context";

export const LangSwitch: React.FC = observer((props) => {
  const { lang, setLang } = useContext(GlobalContext);

  return (
    <Select
      triggerElement={
        <Button
          shape="circle"
          type="secondary"
          style={{
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
