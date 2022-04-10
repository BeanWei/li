import React from "react";
import { Card } from "@arco-design/web-react";
import FormGrid, { IFormGridProps } from "../form-grid";
import ChartLine from "./Chart.Line";
import ChartColumn from "./Chart.Column";
import ChartPie from "./Chart.Pie";

export type ComposedChart = React.FC<IFormGridProps> & {
  Line?: React.FC<any>;
  Column?: React.FC<any>;
  Pie?: React.FC<any>;
};

export const Chart: ComposedChart = (props) => {
  return (
    <Card bodyStyle={{ padding: 0 }}>
      <FormGrid {...props}>{props.children}</FormGrid>
    </Card>
  );
};

Chart.Column = ChartColumn;
Chart.Line = ChartLine;
Chart.Pie = ChartPie;

export default Chart;
