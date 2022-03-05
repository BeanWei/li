import * as components from "../components";
import { SchemaComponentOptions } from "./SchemaComponentOptions";

export const UiSchemaComponentProvider = (props: any) => {
  const { children } = props;
  return (
    // TODO: IMPROVE ME
    // @ts-ignore
    <SchemaComponentOptions components={{ ...components }}>
      {children}
    </SchemaComponentOptions>
  );
};
