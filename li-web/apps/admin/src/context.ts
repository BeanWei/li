import { createContext } from "react";

export const GlobalContext = createContext<{
  app?: Record<string, any>;
  lang?: string;
  setLang?: (value: string) => void;
  theme?: string;
  setTheme?: (value: string) => void;
  currentUser?: Record<string, any>;
  setCurrentUser?: (value: Record<string, any>) => void;
}>({});
