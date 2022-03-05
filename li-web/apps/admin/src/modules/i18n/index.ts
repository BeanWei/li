import i18next from "i18next";
import { initReactI18next } from "react-i18next";

// const zhCN = require("../locale/zh_CN");
// const enUS = require("../locale/en_US");

export const i18n = i18next.createInstance();

i18n.use(initReactI18next).init({
  lng: localStorage.getItem("locale") || "en-US",
  debug: false,
  defaultNS: "client",
  // resources: {
  //   "en-US": {
  //     client: {
  //       ...enUS,
  //     },
  //   },
  //   "zh-CN": {
  //     client: {
  //       ...zhCN,
  //     },
  //   },
  // },
});

i18n.on("languageChanged", (lng) => {
  localStorage.setItem("locale", lng);
});
