import { Schema } from "@formily/react";
import { Spin } from "@arco-design/web-react";
import React from "react";
import { useSchemaComponentContext } from "schema-components/src/hooks";
import { useRequest } from "pro-utils";
import { SchemaComponent } from "schema-components";
import { trimStart, unset } from "lodash";

export interface RemoteSchemaComponentProps {
  scope?: any;
  uid?: string;
  onSuccess?: any;
  schemaTransform?: (schema: Schema) => Schema;
  render?: any;
  hidden?: any;
}

const defaultTransform = (s: Schema) => s;

const parseSchemaStr = (
  s: string,
  removes?: string[],
  pageKey?: string
): any => {
  const o = JSON.parse(s);
  removes?.forEach((path) => {
    unset(o, pageKey ? trimStart(path, pageKey + ".") : path);
  });
  return o;
};

const Loading: React.FC = () => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
      }}
    >
      <Spin />
    </div>
  );
};

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
  // const cacheKey = "@getAppView|" + uid;
  const { reset } = useSchemaComponentContext();
  const { data, loading } = useRequest(
    "@getAppView",
    { key: uid },
    {
      refreshDeps: [uid],
      onSuccess(data: any) {
        onSuccess && onSuccess(data);
        reset && reset();
      },
    }
  );

  if (loading) {
    return <Loading />;
  }
  if (hidden) {
    return <Loading />;
  }
  return (
    <SchemaComponent
      memoized
      scope={scope}
      schema={schemaTransform(
        data ? parseSchemaStr(data.schema, data.removes, uid) : {}
      )}
    />
  );
};

export const RemoteSchemaComponent: React.FC<RemoteSchemaComponentProps> = (
  props
) => {
  return props.uid ? <RequestSchemaComponent {...props} /> : null;
};
