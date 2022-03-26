import { Button } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormModal from "../form-modal";
import { SchemaComponent, UiSchemaComponentProvider } from "../..";
import { ActionFormModalProps } from "./types";

export const ActionFormModal: React.FC<ActionFormModalProps> = observer(
  (props) => {
    const {
      initialValues,
      forInit,
      forInitVariables,
      forSubmit,
      forSubmitSuccess,
      isMenuItem,
      modalProps,
      layoutProps,
      ...buttonProps
    } = props;

    const schema = useFieldSchema();
    const field = useField();

    const handleClick = () => {
      FormModal(modalProps || field.title, () => {
        return (
          <FormLayout {...layoutProps}>
            <UiSchemaComponentProvider>
              <SchemaComponent schema={schema} onlyRenderProperties />
            </UiSchemaComponentProvider>
          </FormLayout>
        );
      })
        .forOpen((paylod, next) => {
          if (forInit) {
            request(forInit, forInitVariables).then((data) => {
              next({
                initialValues: {
                  ...initialValues,
                  ...data,
                },
              });
            });
          } else {
            next({ initialValues: initialValues });
          }
        })
        .forConfirm((payload, next) => {
          if (forSubmit) {
            request(forSubmit, payload).then(() => {
              next(payload);
              forSubmitSuccess?.(payload);
            });
          } else {
            next(payload);
          }
        })
        .open();
    };

    return (
      <FormModal.Portal>
        {isMenuItem ? (
          <div onClick={handleClick}>{field.title}</div>
        ) : (
          <Button {...buttonProps} onClick={handleClick}>
            {field.title}
          </Button>
        )}
      </FormModal.Portal>
    );
  }
);

export default ActionFormModal;
