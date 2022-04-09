import { connect, useFieldSchema } from "@formily/react";
import { Card, CardProps, Space, Typography } from "@arco-design/web-react";
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
      }}
    >
      <Card {...rest}>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <Typography.Title heading={6}>
            {title || fieldSchema.title}
          </Typography.Title>
          <Space>
            <span className="li-chartitem-extra-icon">
              <IconDownload style={{ cursor: "pointer" }} />
            </span>
            <span className="li-chartitem-extra-icon">
              <IconRefresh
                style={{ cursor: "pointer" }}
                onClick={() => run()}
              />
            </span>
          </Space>
        </div>
        {props.children}
      </Card>
    </ChartItemContext.Provider>
  );
});
