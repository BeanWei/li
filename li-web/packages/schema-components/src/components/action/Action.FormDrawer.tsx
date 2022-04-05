import { useContext } from "react";
import { Button, ConfigProvider } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormDrawer from "../form-drawer";
import { ActionFormDrawerProps } from "./types";
import FormButtonGroup from "../form-button-group";
import Submit from "../submit";
import { SchemaComponent, UiSchemaComponentProvider } from "../..";
import { Icon } from "../__builtins__";

export const ActionFormDrawer: React.FC<ActionFormDrawerProps> = observer(
  (props) => {
    const {
      initialValues,
      forInit,
      forInitVariables,
      forSubmitSuccess,
      isMenuItem,
      actionText,
      drawerProps,
      layoutProps,
      ...buttonProps
    } = props;

    const schema = useFieldSchema();
    const field = useField();
    const { locale } = useContext(ConfigProvider.ConfigContext);

    const handleClick = () => {
      const drawer = FormDrawer(
        drawerProps || actionText || field.title,
        () => {
          return (
            <FormLayout {...layoutProps}>
              <UiSchemaComponentProvider>
                <SchemaComponent
                  schema={schema.items as any}
                  onlyRenderProperties
                />
              </UiSchemaComponentProvider>
              {schema.properties && (
                <FormDrawer.Footer>
                  <FormButtonGroup align="right">
                    {schema.reduceProperties((b: React.ReactNode[], s) => {
                      if (s["x-component"] === "Action.FormDrawer.Cancel") {
                        return b.concat([
                          <Button
                            {...drawerProps?.cancelButtonProps}
                            {...s["x-component-props"]}
                            key={s.name}
                            onClick={() => {
                              drawer.close();
                            }}
                          >
                            {s.title || locale?.Drawer.cancelText}
                          </Button>,
                        ]);
                      }
                      if (s["x-component"] === "Action.FormDrawer.Submit") {
                        return b.concat([
                          <Submit
                            {...drawerProps?.cancelButtonProps}
                            {...s["x-component-props"]}
                            key={s.name}
                            forSubmitSuccess={(paylod) => {
                              drawer.close();
                              forSubmitSuccess?.(paylod);
                            }}
                          >
                            {s.title || locale?.Drawer.okText}
                          </Submit>,
                        ]);
                      }
                      return b;
                    }, [])}
                  </FormButtonGroup>
                </FormDrawer.Footer>
              )}
            </FormLayout>
          );
        }
      ).forOpen((paylod, next) => {
        if (forInit) {
          request(forInit, forInitVariables).then((data) => {
            next({
              initialValues: {
                ...initialValues,
                ...data,
              },
              editable:
                schema.reduceProperties((b: boolean[], s) => {
                  if (s["x-component"] === "Action.FormDrawer.Submit") {
                    return b.concat([true]);
                  }
                  return b;
                }, []).length > 0,
            });
          });
        } else {
          next({
            initialValues: initialValues,
            editable:
              schema.reduceProperties((b: boolean[], s) => {
                if (s["x-component"] === "Action.FormDrawer.Submit") {
                  return b.concat([true]);
                }
                return b;
              }, []).length > 0,
          });
        }
      });
      drawer.open();
    };

    return isMenuItem ? (
      <div onClick={handleClick}>{actionText || field.title}</div>
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
        {actionText || field.title}
      </Button>
    );
  }
);

export default ActionFormDrawer;
