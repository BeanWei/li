import { observer, RecursionField, useFieldSchema } from "@formily/react";
import { Space } from "@arco-design/web-react";

export const ActionBar = observer((props: any) => {
  const { style, ...rest } = props;
  const fieldSchema = useFieldSchema();
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        ...style,
      }}
      {...rest}
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
            if (schema["x-component-props"]?.["position"] !== "left") {
              return null;
            }
            return <RecursionField key={key} name={key} schema={schema} />;
          })}
        </Space>
        <Space>
          {fieldSchema.mapProperties((schema, key) => {
            if (schema["x-component-props"]?.["position"] === "left") {
              return null;
            }
            return <RecursionField key={key} name={key} schema={schema} />;
          })}
        </Space>
      </div>
    </div>
  );
});
