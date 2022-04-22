import i18next from "i18next";
import { initReactI18next } from "react-i18next";
import HttpApi from "i18next-http-backend";

const getLocalStorage = (): string | undefined => {
  try {
    return JSON.parse(localStorage.getItem("li-lang") || '""');
  } catch {
    return;
  }
};

export const i18n = i18next.createInstance();

i18n
  .use(HttpApi)
  .use(initReactI18next)
  .init({
    debug: false,
    lng: getLocalStorage() || navigator.language,
    supportedLngs: ["en-US", "zh-CN"],
    preload: ["en-US", "zh-CN"],
    backend: {
      loadPath: "/locales/{{lng}}.json",
    },
    // lient 默认的文案
    resources: {
      "en-US": {
        translation: {
          addNew: "Add New",
          bulkDelete: "Bulk Delete",
          confirmDelete: "Are you sure you want to delete it?",
          columnAction: "Actions",
          viewDrawerTitle: "View",
          editDrawerTitle: "Edit",
        },
      },
      "zh-CN": {
        translation: {
          addNew: "新建",
          bulkDelete: "删除",
          confirmDelete: "你确定要删除吗？",
          columnAction: "操作",
          viewDrawerTitle: "查看",
          editDrawerTitle: "编辑",
        },
      },
    },
  });
