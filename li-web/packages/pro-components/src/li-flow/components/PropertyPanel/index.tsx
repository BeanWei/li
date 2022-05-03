import { Fragment, useEffect } from "react";
import { Form } from "@arco-design/web-react";
import { PropertyCondition, PropertyName } from "./property";

const PropertyPanel: React.FC<{
  activeNode: any;
  onChange: (id: string, values: Record<string, any>) => void;
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
