import { Button } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormModal from "../form-modal";
import { SchemaComponent, UiSchemaComponentProvider } from "../..";
import { ActionFormModalProps } from "./types";

export const ActionFormModal: ActionFormModalProps = observer((props) => {
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
        if (props.forOpen) {
          const result = await request(props.forOpen, props.forOpenVariables);
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
      .forConfirm(async (payload, next) => {
        if (props.forSubmit) {
          await request(props.forSubmit, payload);
          next(payload);
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
});

export default ActionFormModal;