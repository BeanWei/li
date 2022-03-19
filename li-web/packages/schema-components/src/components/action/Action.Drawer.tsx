import { css } from "@emotion/css";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { Drawer } from "@arco-design/web-react";
import { createPortal } from "react-dom";
import { useActionContext } from "./hooks";
import { ComposedActionDrawer } from "./types";

export const ActionDrawer: ComposedActionDrawer = observer((props) => {
  const { footerNodeName = "Action.Drawer.Footer", ...rest } = props;
  const { visible, setVisible } = useActionContext();
  const schema = useFieldSchema();
  const field = useField();
  const footerSchema: any = schema.reduceProperties((buf, s) => {
    if (s["x-component"] === footerNodeName) {
      return s;
    }
    return buf;
  });

  return createPortal(
    <Drawer
      width="50%"
      title={field.title}
      autoFocus={false}
      {...rest}
      visible={visible}
      unmountOnExit
      onCancel={() => setVisible(false)}
      className={rest.className}
      footer={
        footerSchema ? (
          <div
            className={css`
              display: flex;
              justify-content: flex-end;
              width: 100%;
              .arco-btn {
                margin-right: 8px;
              }
            `}
          >
            <RecursionField
              basePath={field.address}
              schema={schema}
              onlyRenderProperties
              filterProperties={(s) => {
                return s["x-component"] === footerNodeName;
              }}
            />
          </div>
        ) : null
      }
    >
      <RecursionField
        basePath={field.address}
        schema={schema}
        onlyRenderProperties
        filterProperties={(s) => {
          return s["x-component"] !== footerNodeName;
        }}
      />
    </Drawer>,
    document.body
  );
});

ActionDrawer.Footer = observer(() => {
  const field = useField();
  const schema = useFieldSchema();

  return (
    <RecursionField
      basePath={field.address}
      schema={schema}
      onlyRenderProperties
    />
  );
});

export default ActionDrawer;
