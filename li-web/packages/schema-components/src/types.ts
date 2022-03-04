import React from "react";
import { Form } from "@formily/core";
import {
  IRecursionFieldProps,
  ISchemaFieldProps,
  SchemaReactComponents,
} from "@formily/react";

export interface ISchemaComponentContext {
  scope?: any;
  components?: SchemaReactComponents;
  refresh?: () => void;
  reset?: () => void;
  SchemaField?: React.FC<ISchemaFieldProps>;
}

export interface ISchemaComponentProvider {
  form?: Form;
  scope?: any;
  components?: SchemaReactComponents;
}

export interface IRecursionComponentProps extends IRecursionFieldProps {
  scope?: any;
  components?: SchemaReactComponents;
}

export interface ISchemaComponentOptionsProps {
  scope?: any;
  components?: SchemaReactComponents;
  inherit?: boolean;
}
