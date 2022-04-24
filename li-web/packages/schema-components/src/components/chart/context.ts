import { createContext } from "react";

export type ChartItemContextProps = {
  data: Record<string, any>[];
  loading: boolean;
  chartRef?: React.MutableRefObject<any>;
};

export const ChartItemContext = createContext<ChartItemContextProps>({
  data: [],
  loading: false,
});
