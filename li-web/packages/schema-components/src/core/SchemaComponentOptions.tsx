import {
  SchemaExpressionScopeContext,
  SchemaOptionsContext,
} from "@formily/react";
import { FormPath } from "@formily/shared";
import React, { useContext } from "react";
import { ISchemaComponentOptionsProps } from "../types";

export const SchemaComponentOptions: React.FC<
  React.PropsWithChildren<ISchemaComponentOptionsProps>
> = (props) => {
  let options = useContext(SchemaOptionsContext);
  const expressionScope = useContext(SchemaExpressionScopeContext);
  const scope = { ...options?.scope, ...expressionScope, ...props.scope };

  return (
    <SchemaOptionsContext.Provider
      value={{
        scope,
        getComponent(name) {
          const propsComponent = FormPath.getIn(props.components, name);
          if (propsComponent) return propsComponent;
          return options.getComponent(name);
        },
      }}
    >
      <SchemaExpressionScopeContext.Provider value={scope}>
        {props.children}
      </SchemaExpressionScopeContext.Provider>
    </SchemaOptionsContext.Provider>
  );
};
