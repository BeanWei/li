import React from "react";
import { RecursionField, useFieldSchema } from "@formily/react";
import ChartLine from "./Chart.Line";

export type ComposedChart = React.FC & {
  Line?: React.FC<any>;
};

export const Chart: ComposedChart = (props) => {
  const fieldSchema = useFieldSchema();

  return (
    <RecursionField schema={fieldSchema["items"] as any} onlyRenderProperties />
  );
};

Chart.Line = ChartLine;

export default Chart;
