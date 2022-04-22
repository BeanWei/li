import { Locale as ArcoLocale } from "@arco-design/web-react/es/locale/interface";

export type Locale = ArcoLocale & {
  ColorSelect: {
    red: string;
    orangered: string;
    orange: string;
    gold: string;
    lime: string;
    green: string;
    cyan: string;
    blue: string;
    arcobule: string;
    purple: string;
    pinkpurple: string;
    magenta: string;
    gray: string;
  };
  List: {
    query: string;
    reset: string;
    search: string;
    confirmDelete: string;
  };
  ListFilter: {
    yes: string;
    no: string;
  };
  RecordPicker: {
    drawerTitle: string;
  };
};
