import ArcoZhCN from "@arco-design/web-react/es/locale/en-US";
import { Locale } from "./types";

export const zhCN: Locale = {
  ...ArcoZhCN,
  Popconfirm: {
    okText: "确定",
    cancelText: "取消",
  },
  List: {
    search: "查询",
    reset: "重置",
    confirmDelete: "你确定要删除吗？",
  },
  RecordPicker: {
    drawerTitle: "请选择",
  },
};
