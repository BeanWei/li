import ArcoEnUS from "@arco-design/web-react/es/locale/en-US";
import { Locale } from "./types";

export const enUS: Locale = {
  ...ArcoEnUS,
  Popconfirm: {
    okText: "OK",
    cancelText: "Cancel",
  },
  List: {
    query: "Query",
    reset: "Reset",
    search: "Search",
    confirmDelete: "Are you sure you want to delete?",
  },
  RecordPicker: {
    drawerTitle: "Please select",
  },
};
