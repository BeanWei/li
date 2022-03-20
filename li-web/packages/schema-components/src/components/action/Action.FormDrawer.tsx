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

export const ActionFormDrawer: React.FC<ActionFormDrawerProps> = observer(
  (props) => {
    const {
      initialValues,
      forInit,
      forInitVariables,
      forSubmit,
      forSubmitSuccess,
      isMenuItem,
      drawerProps,
      layoutProps,
      ...buttonProps
    } = props;

    const schema = useFieldSchema();
    const field = useField();
    const { locale } = useContext(ConfigProvider.ConfigContext);

    const handleClick = () => {
      const drawer = FormDrawer(drawerProps || field.title, () => {
        return (
          <FormLayout {...layoutProps}>
            <UiSchemaComponentProvider>
              <SchemaComponent schema={schema} onlyRenderProperties />
            </UiSchemaComponentProvider>
            <FormDrawer.Footer>
              <FormButtonGroup align="right">
                <Button
                  {...drawerProps?.cancelButtonProps}
                  onClick={() => {
                    drawer.close();
                  }}
                >
                  {drawerProps?.cancelText || locale?.Drawer.cancelText}
                </Button>
                <Submit
                  {...drawerProps?.okButtonProps}
                  forSubmit={forSubmit}
                  forSubmitSuccess={(paylod) => {
                    drawer.close();
                    forSubmitSuccess?.(paylod);
                  }}
                >
                  {drawerProps?.okText || locale?.Drawer.okText}
                </Submit>
              </FormButtonGroup>
            </FormDrawer.Footer>
          </FormLayout>
        );
      }).forOpen(async (paylod, next) => {
        if (forInit) {
          const result = await request(forInit, forInitVariables);
          next({
            initialValues: {
              ...initialValues,
              ...result.data.data,
            },
          });
        } else {
          next({
            initialValues: initialValues,
          });
        }
      });
      drawer.open();
    };

    return isMenuItem ? (
      <div onClick={handleClick}>{field.title}</div>
    ) : (
      <Button {...buttonProps} onClick={handleClick}>
        {field.title}
      </Button>
    );
  }
);

export default ActionFormDrawer;
