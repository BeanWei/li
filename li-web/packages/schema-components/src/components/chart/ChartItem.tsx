import { useRef } from "react";
import { connect, useFieldSchema } from "@formily/react";
import {
  Button,
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
        <Grid.Row
          justify="space-between"
          align="center"
          style={{ marginBottom: 16 }}
        >
          <Grid.Col flex="auto">
            <Typography.Title heading={6} style={{ marginBottom: 0 }}>
              {title || fieldSchema.title}
            </Typography.Title>
          </Grid.Col>
          <Grid.Col flex="60px">
            <Space size={8}>
              <Button
                type="text"
                iconOnly
                style={{ color: "var(--color-text-2)" }}
              >
                <IconDownload
                  onClick={() => {
                    // @ts-ignore
                    ref.current?.downloadImage();
                  }}
                />
              </Button>
              <Button
                type="text"
                iconOnly
                style={{ color: "var(--color-text-2)" }}
              >
                <IconRefresh onClick={() => run()} />
              </Button>
            </Space>
          </Grid.Col>
        </Grid.Row>
        {props.children}
      </Card>
    </ChartItemContext.Provider>
  );
});
