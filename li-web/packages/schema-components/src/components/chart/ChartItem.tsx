import React, { useRef } from "react";
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
import { useTranslation } from "react-i18next";
import { useRequest } from "pro-utils";
import { ChartItemContext } from "./context";
import { downloadImage } from "./utils";

export type ChartItemProps = CardProps & {
  subTitle?: string;
  forInit?: string;
  forInitVariables?: Record<string, any>;
  gridSpan?: number;
};

export const ChartItem = connect(
  (props: React.PropsWithChildren<ChartItemProps>) => {
    const {
      forInit = "",
      forInitVariables,
      title,
      subTitle,
      gridSpan,
      ...rest
    } = props;
    const fieldSchema = useFieldSchema();
    const chartRef = useRef();
    const { t } = useTranslation();

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
          chartRef,
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
                {t(title || fieldSchema.title)}
                {subTitle && (
                  <span
                    style={{
                      fontSize: 12,
                      fontWeight: 400,
                      marginLeft: 4,
                      color: "var(--color-text-3)",
                    }}
                  >
                    {t(subTitle)}
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
                      downloadImage(
                        // @ts-ignore
                        chartRef.current?.plot.chart,
                        t(title || fieldSchema.title)
                      );
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
  }
);
