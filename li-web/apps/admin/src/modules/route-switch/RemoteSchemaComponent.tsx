import { Schema } from "@formily/react";
import { Spin } from "@arco-design/web-react";
import React from "react";
import { useSchemaComponentContext } from "schema-components/src/hooks";
import { useRequest } from "pro-utils";
import { SchemaComponent } from "schema-components";

export interface RemoteSchemaComponentProps {
  scope?: any;
  uid?: string;
  onSuccess?: any;
  schemaTransform?: (schema: Schema) => Schema;
  render?: any;
  hidden?: any;
}

const defaultTransform = (s: Schema) => s;

const RequestSchemaComponent: React.FC<RemoteSchemaComponentProps> = (
  props
) => {
  const {
    hidden,
    scope,
    uid,
    onSuccess,
    schemaTransform = defaultTransform,
  } = props;
  const { reset } = useSchemaComponentContext();
  const { data, loading } = useRequest(
    "@getAppView",
    { uid },
    {
      refreshDeps: [uid],
      onSuccess(data: any) {
        onSuccess && onSuccess(data);
        reset && reset();
      },
    }
  );

  if (loading) {
    return <Spin />;
  }
  if (hidden) {
    return <Spin />;
  }
  return (
    <SchemaComponent
      memoized
      scope={scope}
      schema={schemaTransform(data || ({} as any))}
    />
  );
};

export const RemoteSchemaComponent: React.FC<RemoteSchemaComponentProps> = (
  props
) => {
  return props.uid ? <RequestSchemaComponent {...props} /> : null;
};
