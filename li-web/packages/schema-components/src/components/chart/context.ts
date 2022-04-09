import { createContext } from "react";
import { IDataItem } from "./types";

export type ChartItemContextProps = {
  data: IDataItem[];
  loading: boolean;
};

export const ChartItemContext = createContext<ChartItemContextProps>({
  data: [],
  loading: false,
});
