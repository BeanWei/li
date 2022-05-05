import { Fragment, useEffect } from "react";
import { Form, SelectProps } from "@arco-design/web-react";
import { PropertyApprovers, PropertyCondition, PropertyName } from "./property";
import { eletype } from "../../config";

const PropertyPanel: React.FC<{
  activeNode: any;
  onChange: (id: string, values: Record<string, any>) => void;
  userOptions?: SelectProps["options"];
  roleOptions?: SelectProps["options"];
}> = (props) => {
  const [form] = Form.useForm();

  const renderForm = () => {
    switch (props.activeNode.type) {
      case eletype.sequenceflow:
        return (
          <Fragment>
            <PropertyName />
            <PropertyCondition />
          </Fragment>
        );
      case eletype.usertask:
        return (
          <Fragment>
            <PropertyName />
            <PropertyApprovers
              userOptions={props.userOptions}
              roleOptions={props.roleOptions}
            />
          </Fragment>
        );
      default:
        return <PropertyName />;
    }
  };

  useEffect(() => {
    form.resetFields();
    form.setFieldsValue(props.activeNode.properties);
  }, [props.activeNode.id]);

  return (
    <Form
      form={form}
      layout="vertical"
      onValuesChange={(_, values) => {
        props.onChange(props.activeNode.id, values);
      }}
    >
      {renderForm()}
    </Form>
  );
};

export default PropertyPanel;
