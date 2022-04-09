import { createContext } from "react";
import { IDataItem } from "./types";

export type ChartItemContextProps = {
  data: IDataItem[];
  loading: boolean;
  setChartRef?: (plot: any) => void;
};

export const ChartItemContext = createContext<ChartItemContextProps>({
  data: [],
  loading: false,
});
