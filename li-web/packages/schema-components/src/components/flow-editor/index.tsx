import { CSSProperties, useRef } from "react";
import { connect, mapReadPretty } from "@formily/react";
import { LiFlow } from "pro-components";
import { useRequest } from "pro-utils";

type FlowEditorProps = {
  height?: CSSProperties["height"];
  fetchUserConfig?: {
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
    const { fetchUserConfig, ...rest } = props;
    const { data, run } = useRequest(
      fetchUserConfig?.operation || "",
      fetchUserConfig?.variables,
      {
        manual: true,
      }
    );
    const ref = useRef(false);
    if (!ref.current && props.fetchUserConfig) {
      run();
      ref.current = true;
    }

    return (
      <LiFlow
        {...rest}
        userOptions={(data && Array.isArray(data) ? data : data?.list)?.map(
          (item: any) => {
            return {
              label: item[fetchUserConfig?.fieldNames.label || "id"],
              value: item[fetchUserConfig?.fieldNames.value || "id"],
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
