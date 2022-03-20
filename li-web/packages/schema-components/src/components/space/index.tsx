import React from "react";
import {
  Divider,
  Space as ArcoSpace,
  SpaceProps,
} from "@arco-design/web-react";
import { RecursionField, useFieldSchema } from "@formily/react";
import { useFormLayout } from "../form-layout";

export const Space: React.FC<SpaceProps> = (props) => {
  const fieldSchema = useFieldSchema();
  const layout = useFormLayout();
  let { split } = props;
  if (split === "divider") {
    split = <Divider type="vertical" style={{ margin: "0 2px" }} />;
  }
  return (
    <ArcoSpace size={props.size ?? layout?.spaceGap} {...props} split={split}>
      {fieldSchema.mapProperties((schema, key) => {
        return <RecursionField key={key} name={key} schema={schema} />;
      })}
    </ArcoSpace>
  );
};

export default Space;
