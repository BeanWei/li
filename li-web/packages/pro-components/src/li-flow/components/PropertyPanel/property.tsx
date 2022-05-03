import { Form, Input, Select, SelectProps } from "@arco-design/web-react";

export const PropertyName: React.FC = () => {
  return (
    <Form.Item label="标题" field="name">
      <Input />
    </Form.Item>
  );
};

export const PropertyCondition: React.FC = () => {
  return (
    <Form.Item label="条件表达式" field="condition">
      <Input.TextArea />
    </Form.Item>
  );
};

export const PropertyApprovers: React.FC<{
  options: SelectProps["options"];
}> = (props) => {
  return (
    <Form.Item label="审批人" field="approvers">
      <Select options={props.options} />
    </Form.Item>
  );
};
