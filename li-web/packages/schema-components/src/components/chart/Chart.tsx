import React from "react";
import { Card, CardProps } from "@arco-design/web-react";
import ChartLine from "./Chart.Line";
import ChartColumn from "./Chart.Column";
import ChartPie from "./Chart.Pie";

export type ComposedChart = React.FC<CardProps> & {
  Line?: React.FC<any>;
  Column?: React.FC<any>;
  Pie?: React.FC<any>;
};

export const Chart: ComposedChart = (props) => {
  return (
    <Card {...props} bodyStyle={{ ...props.bodyStyle, padding: 0 }}>
      {props.children}
    </Card>
  );
};

Chart.Column = ChartColumn;
Chart.Line = ChartLine;
Chart.Pie = ChartPie;

export default Chart;
