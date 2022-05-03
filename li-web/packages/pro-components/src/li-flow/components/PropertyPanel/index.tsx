import { Fragment, useEffect } from "react";
import { Form, SelectProps } from "@arco-design/web-react";
import { PropertyApprovers, PropertyCondition, PropertyName } from "./property";

const PropertyPanel: React.FC<{
  activeNode: any;
  onChange: (id: string, values: Record<string, any>) => void;
  userOptions?: SelectProps["options"];
}> = (props) => {
  const [form] = Form.useForm();

  const renderForm = () => {
    switch (props.activeNode.type) {
      case "SequenceFlow:SequenceFlow":
        return (
          <Fragment>
            <PropertyName />
            <PropertyCondition />
          </Fragment>
        );
      case "UserTask:UserTask":
        return (
          <Fragment>
            <PropertyName />
            <PropertyApprovers options={props.userOptions} />
          </Fragment>
        );
      default:
        return <PropertyName />;
    }
  };

  useEffect(() => {
    form.clearFields();
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
