import React from "react";
import { Card } from "@arco-design/web-react";
import FormGrid, { IFormGridProps } from "../form-grid";
import ChartAutoChart from "./Chart.AutoChart";

export type ComposedChart = React.FC<IFormGridProps> & {
  AutoChart?: React.FC<any>;
};

export const Chart: ComposedChart = (props) => {
  return (
    <Card bodyStyle={{ padding: 0 }}>
      <FormGrid {...props}>{props.children}</FormGrid>
    </Card>
  );
};

Chart.AutoChart = ChartAutoChart;

export default Chart;
