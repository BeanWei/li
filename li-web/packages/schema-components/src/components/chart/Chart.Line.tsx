import { useContext } from "react";
import {
  Area,
  AreaConfig,
  DualAxes,
  DualAxesConfig,
  Line as ACLine,
  LineConfig,
} from "@ant-design/plots";
import { connect } from "@formily/react";
import { ChartItemContext } from "./context";
import { ICommonProps, IDataItem, IMetaItem } from "./types";
import { splitMeta, strip } from "./utils";

export type LineProps = Omit<
  LineConfig,
  keyof ICommonProps | "yField" | "xField" | "seriesField"
> &
  ICommonProps & {
    /**
     * 以面积图展示，默认 `false`
     * - 注意面积图默认堆叠展示，如不需要可以传入 `isStack={false}` 覆盖
     * - 开启面积图后方可使用 `areaStyle` `startOnZero` `isPercent` 属性
     */
    withArea?: boolean;
  };
export type AreaProps = Omit<
  AreaConfig,
  keyof ICommonProps | "yField" | "xField" | "seriesField"
> &
  ICommonProps;
export type DualAxesProps = Omit<
  DualAxesConfig,
  keyof ICommonProps | "yField" | "xField" | "seriesField"
> &
  ICommonProps;

export function generateConfig(
  meta: IMetaItem[],
  data: IDataItem[]
): AreaConfig | LineConfig | DualAxesConfig {
  const { metaDim, metaInd } = splitMeta(meta);

  if (metaInd.length === 1 && metaDim.length === 1) {
    // case 1: 单指标、单维度 => 维度作为 x 轴，指标作为 y 轴
    const xFieldMeta = metaDim.shift() as IMetaItem;
    const yFieldMeta = metaInd.shift() as IMetaItem;
    const xField = xFieldMeta.id;
    const yField = yFieldMeta.id;
    return {
      data,
      xField,
      yField,
      yAxis: {
        label: {
          formatter: (v) => {
            return yFieldMeta.isRate ? `${strip(100 * Number(v))}%` : v;
          },
        },
      },
      tooltip: {
        formatter: ({ [xField]: type, [yField]: value }) => {
          return {
            name: yFieldMeta.name,
            value: yFieldMeta.isRate ? `${strip(100 * Number(value))}%` : value,
          };
        },
      },
      meta: {
        [yField]: { alias: meta.find(({ id }) => id === yField)?.name },
      },
    };
  } else if (metaInd.length === 1 && metaDim.length === 2) {
    // case 2: 单指标、双维度 => 第一维度作为 x 轴，指标作为 y 轴，第二维度作为 系列
    const xFieldMeta = metaDim.shift() as IMetaItem;
    const yFieldMeta = metaInd.shift() as IMetaItem;
    const seriesFieldMeta = metaDim.shift() as IMetaItem;
    const xField = xFieldMeta.id;
    const yField = yFieldMeta.id;
    const seriesField = seriesFieldMeta.id;
    return {
      data,
      xField,
      yField,
      seriesField,
      yAxis: {
        label: {
          formatter: (v) => {
            return yFieldMeta.isRate ? `${strip(100 * Number(v))}%` : v;
          },
        },
      },
      tooltip: {
        formatter: ({ [seriesField]: type, [yField]: value }) => {
          return {
            name: type,
            value: yFieldMeta.isRate ? `${strip(100 * Number(value))}%` : value,
          };
        },
      },
    };
  } else if (metaInd.length === 2 && metaDim.length === 2) {
    // case 3: 双指标、双维度 => 第一维度作为 x 轴，第二维度作为 系列，第一指标作为左 y 轴，第二指标作为右 y 轴
    const data1 = data.map(
      ({ [metaInd[0].id]: _, [metaDim[1].id]: metaValue, ...item }) => ({
        [metaDim[1].id]: `${metaValue}-${metaInd[0].name}`,
        ...item,
      })
    );
    const data2 = data.map(
      ({ [metaInd[1].id]: _, [metaDim[1].id]: metaValue, ...item }) => ({
        [metaDim[1].id]: `${metaValue}-${metaInd[1].name}`,
        ...item,
      })
    );
    return {
      data: [data2, data1],
      geometryOptions: [
        { geometry: "line", seriesField: metaDim[1].id },
        {
          geometry: "line",
          seriesField: metaDim[1].id,
          lineStyle: { lineDash: [5, 5] },
        },
      ],
      xField: metaDim.shift()?.id as string,
      yField: [metaInd.shift()?.id as string, metaInd.shift()?.id as string],
      seriesField: metaDim.shift()?.id,
    };
  } else if (metaInd.length > 1 && metaDim.length === 1) {
    // case 4: 多指标、单维度 => 维度作为 x 轴，指标名作为系列，指标值作为 y 轴
    // 需要把 data 做一下转化，例：从 { ds, uv, pv }[] 转为 [{ ds, type: uvName, value: xxx }, { ds, type: pvName, value: xxx }]
    const xField = metaDim.shift()?.id as string;
    const yField = "value";
    const seriesField = "type";
    return {
      data: data
        .map((item) => {
          return metaInd.map(({ id, name }) => {
            return {
              [xField]: item[xField],
              [yField]: item[id],
              [seriesField]: name,
            };
          });
        })
        .flat(),
      xField,
      yField,
      seriesField,
    };
  }
  return { data };
}

export const Line: React.FC<LineProps | AreaProps | DualAxesProps> = ({
  className,
  style,
  meta = [],
  data = [],
  // @ts-ignore
  withArea,
  ...props
}) => {
  const config = generateConfig(meta, data);

  // 面积图展示
  if (withArea) {
    return (
      <Area
        {...(config as AreaConfig)}
        {...(props as Omit<
          AreaConfig,
          keyof ICommonProps | "yField" | "xField" | "seriesField"
        >)}
      />
    );
  }

  // 双轴图展示
  if (Array.isArray(config.yField)) {
    return (
      <DualAxes
        {...(config as DualAxesConfig)}
        {...(props as Omit<
          DualAxesConfig,
          keyof ICommonProps | "yField" | "xField" | "seriesField"
        >)}
      />
    );
  }

  // 普通折线图
  return (
    <ACLine
      {...(config as LineConfig)}
      {...(props as Omit<
        LineConfig,
        keyof ICommonProps | "yField" | "xField" | "seriesField"
      >)}
    />
  );
};

export const ChartLine = connect(
  (props: LineProps | AreaProps | DualAxesProps) => {
    const ctx = useContext(ChartItemContext);
    return (
      <Line
        {...props}
        chartRef={(plot) => ctx.setChartRef?.(plot)}
        loading={ctx.loading}
        data={ctx.data.length ? ctx.data : props.data}
      />
    );
  }
);

export default ChartLine;
