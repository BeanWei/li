import * as components from ".";
import { SchemaComponentOptions } from "../core/SchemaComponentOptions";

export const AntdSchemaComponentProvider = (props: any) => {
  const { children } = props;
  return (
    <SchemaComponentOptions components={{ ...components }}>
      {children}
    </SchemaComponentOptions>
  );
};
