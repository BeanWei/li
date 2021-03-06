import React, {
  useState,
  useEffect,
  useRef,
  useImperativeHandle,
  useContext,
} from "react";
import { connect } from "@formily/react";
import { specToG2Plot } from "@antv/antv-spec";
import { Advice, Advisor } from "@antv/chart-advisor";
import { AdviseParams } from "@antv/chart-advisor/lib/advisor";
import { Spin } from "@arco-design/web-react";
import { getLocale, getPrefixCls } from "../__builtins__";
import { ChartItemContext } from "./context";
import "./index.less";

const ChartRender: React.FC<{
  spec?: any;
  chartRef?: React.MutableRefObject<any>;
}> = ({ spec, chartRef }) => {
  const plotRef = useRef<HTMLDivElement>(null);
  const [chartType, setChartType] = useState<string | null>(null);
  const [plot, setPlot] = useState<any>(null);
  const prefixCls = getPrefixCls();

  useEffect(() => {
    if (spec && plotRef?.current) {
      const plot = specToG2Plot(spec, plotRef?.current);
      setPlot(plot);
      setChartType(plot.constructor.name);
    }
  }, [spec]);

  useEffect(() => {
    if (spec === null) {
      setChartType(null);
      if (plot) plot.destroy();
      setPlot(null);
    }
  }, [spec]);

  useImperativeHandle(chartRef, () => {
    return {
      chartType,
      plot,
    };
  });

  return (
    <div className={`${prefixCls}-autochart-canvas-layer`}>
      <div className="canvas-content">
        <div className="feedback-layer" ref={plotRef}></div>
      </div>
    </div>
  );
};

const EmptyContent: React.FC = () => {
  const prefixCls = getPrefixCls();
  const locale = getLocale();
  return (
    <div className={`${prefixCls}-autochart-nodata-layer`}>
      <div className={`${prefixCls}-autochart-nodata-content`}>
        <div style={{ marginBottom: 16 }}>
          <img
            src="https://gw.alipayobjects.com/zos/basement_prod/9a59280d-8f23-4234-b5cf-02956a91b6ff.svg"
            alt=""
          />
        </div>
        <div>{locale.Chart?.nodata}</div>
      </div>
    </div>
  );
};

const AutoChart: React.FC<
  AdviseParams & { chartRef?: React.MutableRefObject<any> }
> = ({ chartRef, ...rest }) => {
  const containerRef = useRef<HTMLDivElement>(null);
  const myAdvisor = new Advisor();
  const [advices, setAdvices] = useState<Advice[]>([]);
  const [currentAdviceIndex, setCurrentAdviceIndex] = useState<number>(0);
  const prefixCls = getPrefixCls();

  useEffect(() => {
    if (rest.data?.length > 0) {
      const myAdvices = myAdvisor.advise(rest);
      setAdvices(myAdvices as any);
      setCurrentAdviceIndex(0);
    }
  }, [rest.data]);

  return (
    <div className={`${prefixCls}-autochart-container`} ref={containerRef}>
      {rest.data?.length ? (
        <ChartRender
          chartRef={chartRef}
          spec={advices[currentAdviceIndex]?.spec || null}
        />
      ) : (
        <EmptyContent />
      )}
    </div>
  );
};

export const ChartAutoChart = connect((props) => {
  const ctx = useContext(ChartItemContext);
  return (
    <Spin loading={ctx.loading} style={{ display: "block" }}>
      <AutoChart
        {...props}
        chartRef={ctx.chartRef}
        data={ctx.data.length ? ctx.data : props.data}
      />
    </Spin>
  );
});

export default ChartAutoChart;
