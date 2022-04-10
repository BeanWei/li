import { useContext } from "react";
import { Pie as ACPie, PieConfig } from "@ant-design/plots";
import { connect } from "@formily/react";
import { ChartItemContext } from "./context";
import { ICommonProps, IDataItem, IMetaItem } from "./types";
import { splitMeta } from "./utils";

export type PieProps = Omit<
  PieConfig,
  keyof ICommonProps | "angleField" | "colorField"
> &
  ICommonProps;

export function generateConfig(
  meta: IMetaItem[],
  data: IDataItem[]
): PieConfig {
  const { metaDim, metaInd } = splitMeta(meta);

  if (metaInd.length === 1 && metaDim.length === 1) {
    // case 1: 单指标，单维度 => 维度作为 colorField，指标作为 angleField
    const colorField = metaDim.shift()?.id as string;
    const angleField = metaInd.shift()?.id as string;
    return {
      data,
      colorField,
      angleField,
    };
  }
  return { data, colorField: "", angleField: "" };
}

const Pie: React.FC<PieProps> = ({
  className,
  style,
  meta = [],
  data = [],
  ...props
}) => {
  return <ACPie {...generateConfig(meta, data)} {...props} />;
};

export const ChartPie = connect((props: PieProps) => {
  const ctx = useContext(ChartItemContext);
  return (
    <Pie
      {...props}
      chartRef={(plot) => ctx.setChartRef?.(plot)}
      loading={ctx.loading}
      data={ctx.data.length ? ctx.data : props.data}
    />
  );
});

export default ChartPie;
