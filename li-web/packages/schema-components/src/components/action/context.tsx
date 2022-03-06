import { createContext } from "react";

export const ActionContext = createContext<{
  button?: any;
  visible?: boolean;
  setVisible?: (v: boolean) => void;
  openMode?: "drawer" | "modal" | "page";
  containerRefKey?: string;
}>({});
