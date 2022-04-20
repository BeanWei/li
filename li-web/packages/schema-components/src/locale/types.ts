import { Locale as ArcoLocale } from "@arco-design/web-react/es/locale/interface";

export type Locale = ArcoLocale & {
  List: {
    query: string;
    reset: string;
    search: string;
    confirmDelete: string;
  };
  RecordPicker: {
    drawerTitle: string;
  };
};
