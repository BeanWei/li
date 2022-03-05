import { css } from "@emotion/css";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { Drawer } from "@arco-design/web-react";
import cls from "classnames";
import { createPortal } from "react-dom";
import { useActionContext } from "./hooks";
import { ComposedActionDrawer } from "./types";

export const ActionDrawer: ComposedActionDrawer = observer((props) => {
  const { footerNodeName = "Action.Drawer.Footer", ...others } = props;
  const { visible, setVisible } = useActionContext();
  const schema = useFieldSchema();
  const field = useField();
  const footerSchema = schema.reduceProperties((buf, s) => {
    if (s["x-component"] === footerNodeName) {
      return s;
    }
    return buf;
  });
  return (
    <>
      {createPortal(
        <Drawer
          width={"50%"}
          title={field.title}
          {...others}
          visible={visible}
          onCancel={() => setVisible(false)}
          className={cls(
            others.className,
            css`
              &.li-action-popup {
                .arco-drawer-content {
                  background: #f0f2f5;
                }
              }
            `
          )}
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
      )}
    </>
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
