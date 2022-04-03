import { Locale as ArcoLocale } from "@arco-design/web-react/es/locale/interface";

export type Locale = ArcoLocale & {
  List: {
    search: string;
    reset: string;
    confirmDelete: string;
  };
};
