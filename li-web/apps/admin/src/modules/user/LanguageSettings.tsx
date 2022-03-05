import { useState } from "react";
import { Menu, Select } from "@arco-design/web-react";
import { useTranslation } from "react-i18next";

export const LanguageSettings = () => {
  const { t, i18n } = useTranslation();
  const [open, setOpen] = useState(false);
  return (
    <Menu.Item
      onClick={() => {
        setOpen(true);
      }}
      key="lng"
    >
      {t("Language")}{" "}
      <Select
        style={{ minWidth: 100 }}
        bordered={false}
        onVisibleChange={(open) => {
          setOpen(open);
        }}
        options={[
          { label: "简体中文", value: "zh-CN" },
          { label: "English", value: "en-US" },
        ]}
        value={i18n.language}
        onChange={async (lang) => {
          await i18n.changeLanguage(lang);
          window.location.reload();
        }}
      />
    </Menu.Item>
  );
};
