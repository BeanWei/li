import { Form, Input } from "@arco-design/web-react";

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
