import { useRef } from "react";
import { connect, useFieldSchema } from "@formily/react";
import {
  Card,
  CardProps,
  Grid,
  Space,
  Typography,
} from "@arco-design/web-react";
import { IconDownload, IconRefresh } from "@arco-design/web-react/icon";
import { useRequest } from "pro-utils";
import { ChartItemContext } from "./context";
import { IRequest } from "./types";
import "./index.less";

export type ChartItemProps = CardProps & {
  request?: IRequest;
};

export const ChartItem = connect((props: ChartItemProps) => {
  const { request, title, ...rest } = props;
  const fieldSchema = useFieldSchema();
  const ref = useRef();

  const {
    data = [],
    loading,
    run,
  } = useRequest(request?.operation || "", request?.variables, {
    refreshDeps: [request],
  });

  return (
    <ChartItemContext.Provider
      value={{
        data,
        loading,
        setChartRef: (plot) => {
          ref.current = plot;
        },
      }}
    >
      <Card {...rest}>
        <Grid.Row justify="space-between">
          <Grid.Col flex="auto">
            <Typography.Title heading={6}>
              {title || fieldSchema.title}
            </Typography.Title>
          </Grid.Col>
          <Grid.Col flex="60px">
            <Space>
              <span className="li-chartitem-extra-icon">
                <IconDownload
                  style={{ cursor: "pointer" }}
                  onClick={() => {
                    // @ts-ignore
                    ref.current?.downloadImage();
                  }}
                />
              </span>
              <span className="li-chartitem-extra-icon">
                <IconRefresh
                  style={{ cursor: "pointer" }}
                  onClick={() => run()}
                />
              </span>
            </Space>
          </Grid.Col>
        </Grid.Row>
        {props.children}
      </Card>
    </ChartItemContext.Provider>
  );
});
