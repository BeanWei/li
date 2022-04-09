import React from "react";
import { RecursionField, useFieldSchema } from "@formily/react";
import ChartLine from "./Chart.Line";
import ChartColumn from "./Chart.Column";

export type ComposedChart = React.FC & {
  Line?: React.FC<any>;
  Column?: React.FC<any>;
};

export const Chart: ComposedChart = (props) => {
  const fieldSchema = useFieldSchema();

  return (
    <RecursionField schema={fieldSchema["items"] as any} onlyRenderProperties />
  );
};

Chart.Column = ChartColumn;
Chart.Line = ChartLine;

export default Chart;
