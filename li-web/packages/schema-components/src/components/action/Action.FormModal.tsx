import { useContext } from "react";
import { Button, ConfigProvider } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormModal from "../form-modal";
import {
  FormButtonGroup,
  SchemaComponent,
  Submit,
  UiSchemaComponentProvider,
} from "../..";
import { ActionFormModalProps } from "./types";
import { Icon } from "../__builtins__";

export const ActionFormModal: React.FC<ActionFormModalProps> = observer(
  (props) => {
    const {
      initialValues,
      forInit,
      forInitVariables,
      forSubmitSuccess,
      isMenuItem,
      modalProps,
      layoutProps,
      ...buttonProps
    } = props;

    const schema = useFieldSchema();
    const field = useField();
    const { locale } = useContext(ConfigProvider.ConfigContext);

    const handleClick = () => {
      const modal = FormModal(modalProps || field.title, () => {
        return (
          <FormLayout {...layoutProps}>
            <UiSchemaComponentProvider>
              <SchemaComponent schema={schema} onlyRenderProperties />
            </UiSchemaComponentProvider>
            {schema.properties && (
              <FormModal.Footer>
                <FormButtonGroup align="right">
                  {schema.reduceProperties((b: React.ReactNode[], s) => {
                    if (s["x-component"] === "Action.FormModal.Cancel") {
                      return b.concat([
                        <Button
                          {...modalProps?.cancelButtonProps}
                          {...s["x-component-props"]}
                          key={s.name}
                          onClick={() => {
                            modal.close();
                          }}
                        >
                          {s.title || locale?.Modal.cancelText}
                        </Button>,
                      ]);
                    }
                    if (s["x-component"] === "Action.FormModal.Submit") {
                      return b.concat([
                        <Submit
                          {...modalProps?.cancelButtonProps}
                          {...s["x-component-props"]}
                          key={s.name}
                          forSubmitSuccess={(paylod) => {
                            modal.close();
                            forSubmitSuccess?.(paylod);
                          }}
                        >
                          {s.title || locale?.Modal.okText}
                        </Submit>,
                      ]);
                    }
                    return b;
                  }, [])}
                </FormButtonGroup>
              </FormModal.Footer>
            )}
          </FormLayout>
        );
      }).forOpen((paylod, next) => {
        if (forInit) {
          request(forInit, forInitVariables).then((data) => {
            next({
              initialValues: {
                ...initialValues,
                ...data,
              },
              editable:
                schema.reduceProperties((b: boolean[], s) => {
                  if (s["x-component"] === "Action.FormModal.Submit") {
                    return b.concat([true]);
                  }
                  return b;
                }, []).length > 0,
            });
          });
        } else {
          next({
            initialValues,
            editable:
              schema.reduceProperties((b: boolean[], s) => {
                if (s["x-component"] === "Action.FormModal.Submit") {
                  return b.concat([true]);
                }
                return b;
              }, []).length > 0,
          });
        }
      });
      modal.open();
    };

    return (
      <FormModal.Portal>
        {isMenuItem ? (
          <div onClick={handleClick}>{field.title}</div>
        ) : (
          <Button
            {...buttonProps}
            onClick={handleClick}
            icon={
              buttonProps.icon && typeof buttonProps.icon === "string" ? (
                <Icon type={buttonProps.icon} />
              ) : (
                buttonProps.icon
              )
            }
          >
            {field.title}
          </Button>
        )}
      </FormModal.Portal>
    );
  }
);

export default ActionFormModal;
