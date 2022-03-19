import { Button } from "@arco-design/web-react";
import { observer, useField, useFieldSchema } from "@formily/react";
import { request } from "pro-utils";
import FormLayout from "../form-layout";
import FormDrawer from "../form-drawer";
import { ActionFormDrawerProps } from "./types";
import FormButtonGroup from "../form-button-group";
import Submit from "../submit";
import { SchemaComponent, UiSchemaComponentProvider } from "../..";

export const ActionFormDrawer: ActionFormDrawerProps = observer((props) => {
  const schema = useFieldSchema();
  const field = useField();

  const handleClick = () => {
    const drawer = FormDrawer(props.drawerProps || field.title, () => {
      return (
        <FormLayout {...props.layoutProps}>
          <UiSchemaComponentProvider>
            <SchemaComponent schema={schema} onlyRenderProperties />
          </UiSchemaComponentProvider>
          <FormDrawer.Footer>
            <FormButtonGroup align="right">
              <Button
                {...props.drawerProps?.cancelButtonProps}
                onClick={() => {
                  drawer.close();
                }}
              >
                {props.drawerProps?.cancelText || "Cancel"}
              </Button>
              <Submit
                {...props.drawerProps?.okButtonProps}
                onSubmit={async (values) => {
                  if (props.forSubmit) {
                    request(props.forSubmit, values).then(() => {
                      drawer.close();
                    });
                  }
                }}
              >
                {props.drawerProps?.okText || "Submit"}
              </Submit>
            </FormButtonGroup>
          </FormDrawer.Footer>
        </FormLayout>
      );
    }).forOpen(async (paylod, next) => {
      if (props.forOpen) {
        const result = await request(props.forOpen, props.forOpenVariables);
        next({
          initialValues: {
            ...props.initialValues,
            ...result.data.data,
          },
        });
      } else {
        next({
          initialValues: props.initialValues,
        });
      }
    });
    drawer.open();
  };

  return props.isMenuItem ? (
    <div onClick={handleClick}>{field.title}</div>
  ) : (
    <Button {...props.buttonProps} onClick={handleClick}>
      {field.title}
    </Button>
  );
});

export default ActionFormDrawer;
