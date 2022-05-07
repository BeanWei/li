import { CSSProperties, useRef } from "react";
import { connect, mapReadPretty } from "@formily/react";
import { LiFlow } from "pro-components";
import { useRequest } from "pro-utils";
import { SelectProps } from "@arco-design/web-react";

type FlowEditorProps = {
  height?: CSSProperties["height"];
  notifyChannelOptions?: SelectProps["options"];
  fetchUserConfig?: {
    operation: string;
    variables?: Record<string, any>;
    fieldNames: {
      label: string;
      value: string;
    };
  };
  fetchRoleConfig?: {
    operation: string;
    variables?: Record<string, any>;
    fieldNames: {
      label: string;
      value: string;
    };
  };
};

export const FlowEditor = connect(
  (props: FlowEditorProps) => {
    const { fetchUserConfig, fetchRoleConfig, ...rest } = props;
    const { data: data1, run: run1 } = useRequest(
      fetchUserConfig?.operation || "",
      fetchUserConfig?.variables,
      {
        manual: true,
      }
    );
    const { data: data2, run: run2 } = useRequest(
      fetchRoleConfig?.operation || "",
      fetchRoleConfig?.variables,
      {
        manual: true,
      }
    );
    const ref = useRef(false);
    if (!ref.current) {
      if (props.fetchUserConfig) run1();
      if (props.fetchRoleConfig) run2();
      ref.current = true;
    }

    return (
      <LiFlow
        {...rest}
        userOptions={(data1 && Array.isArray(data1) ? data1 : data1?.list)?.map(
          (item: any) => {
            return {
              label: item[fetchUserConfig?.fieldNames.label || "id"],
              value: item[fetchUserConfig?.fieldNames.value || "id"],
            };
          }
        )}
        roleOptions={(data2 && Array.isArray(data2) ? data2 : data2?.list)?.map(
          (item: any) => {
            return {
              label: item[fetchRoleConfig?.fieldNames.label || "id"],
              value: item[fetchRoleConfig?.fieldNames.value || "id"],
            };
          }
        )}
      />
    );
  },
  mapReadPretty((props) => {
    return <LiFlow readOnly value={props.value} />;
  })
);

export default FlowEditor;
