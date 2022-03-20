import { Button } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormModal from "../form-modal";
import { SchemaComponent, UiSchemaComponentProvider } from "../..";
import { ActionFormModalProps } from "./types";

export const ActionFormModal: React.FC<ActionFormModalProps> = observer(
  (props) => {
    const schema = useFieldSchema();
    const field = useField();

    const handleClick = () => {
      FormModal(props.modalProps || field.title, () => {
        return (
          <FormLayout {...props.layoutProps}>
            <UiSchemaComponentProvider>
              <SchemaComponent schema={schema} onlyRenderProperties />
            </UiSchemaComponentProvider>
          </FormLayout>
        );
      })
        .forOpen(async (paylod, next) => {
          if (props.forInit) {
            const result = await request(props.forInit, props.forInitVariables);
            next({
              initialValues: {
                ...props.initialValues,
                ...result.data.data,
              },
            });
          } else {
            next({ initialValues: props.initialValues });
          }
        })
        .forConfirm((payload, next) => {
          if (props.forSubmit) {
            request(props.forSubmit, payload).then(() => {
              next(payload);
              props.forSubmitSuccess?.(payload);
            });
          } else {
            next(payload);
          }
        })
        .open();
    };

    return (
      <FormModal.Portal>
        {props.isMenuItem ? (
          <div onClick={handleClick}>{field.title}</div>
        ) : (
          <Button {...props.buttonProps} onClick={handleClick}>
            {field.title}
          </Button>
        )}
      </FormModal.Portal>
    );
  }
);

export default ActionFormModal;
