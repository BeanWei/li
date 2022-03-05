import { css } from "@emotion/css";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { Modal } from "@arco-design/web-react";
import { createPortal } from "react-dom";
import { useActionContext } from ".";
import { ComposedActionDrawer } from "./types";

export const ActionModal: ComposedActionDrawer = observer((props) => {
  const { footerNodeName = "Action.Modal.Footer", ...rest } = props;
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
        // @ts-ignore
        <Modal
          title={schema.title}
          {...rest}
          visible={visible}
          onCancel={() => setVisible(false)}
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
        </Modal>,
        document.body
      )}
    </>
  );
});

ActionModal.Footer = observer(() => {
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

export default ActionModal;
