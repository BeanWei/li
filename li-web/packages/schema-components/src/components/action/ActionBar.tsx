import { observer, RecursionField, useFieldSchema } from "@formily/react";
import { Space } from "@arco-design/web-react";

export const ActionBar = observer((props: any) => {
  const { layout = "tow-columns", style, ...others } = props;
  const fieldSchema = useFieldSchema();
  if (layout === "one-column") {
    return (
      <div style={{ display: "flex", ...style }} {...others}>
        {props.children && (
          <div style={{ marginRight: 8 }}>{props.children}</div>
        )}
      </div>
    );
  }
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        ...style,
      }}
      {...others}
    >
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          width: "100%",
        }}
      >
        <Space>
          {fieldSchema.mapProperties((schema, key) => {
            if (schema["x-align"] !== "left") {
              return null;
            }
            return <RecursionField key={key} name={key} schema={schema} />;
          })}
        </Space>
        <Space>
          {fieldSchema.mapProperties((schema, key) => {
            if (schema["x-align"] === "left") {
              return null;
            }
            return <RecursionField key={key} name={key} schema={schema} />;
          })}
        </Space>
      </div>
    </div>
  );
});
