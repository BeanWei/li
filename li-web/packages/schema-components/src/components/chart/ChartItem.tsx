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

export type ChartItemProps = CardProps & {
  subTitle?: string;
  forInit?: string;
  forInitVariables?: Record<string, any>;
  gridSpan?: number;
};

export const ChartItem = connect((props: ChartItemProps) => {
  const {
    forInit = "",
    forInitVariables,
    title,
    subTitle,
    gridSpan,
    ...rest
  } = props;
  const fieldSchema = useFieldSchema();
  const ref = useRef();

  const {
    data = [],
    loading,
    run,
  } = useRequest(forInit, forInitVariables, {
    refreshDeps: [forInitVariables],
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
      <Card {...rest} data-grid-span={gridSpan}>
        <Grid.Row
          justify="space-between"
          align="center"
          style={{ marginBottom: 16 }}
        >
          <Grid.Col flex="auto">
            <Typography.Paragraph
              style={{
                marginBottom: 0,
                fontSize: 16,
                fontWeight: 500,
              }}
            >
              {title || fieldSchema.title}
              {subTitle && (
                <span
                  style={{
                    fontSize: 12,
                    fontWeight: 400,
                    marginLeft: 4,
                    color: "var(--color-text-3)",
                  }}
                >
                  {subTitle}
                </span>
              )}
            </Typography.Paragraph>
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
