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
    // fallbackLng: "en-US",
    // preload: ['en-US', 'zh-CN'],
    // backend: {
    //   loadPath: '/locales/{{lng}}.json',
    // },
  });
