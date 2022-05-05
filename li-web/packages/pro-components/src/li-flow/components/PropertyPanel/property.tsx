import { Fragment } from "react";
import {
  Form,
  Input,
  Radio,
  Select,
  SelectProps,
} from "@arco-design/web-react";

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
  userOptions: SelectProps["options"];
  roleOptions?: SelectProps["options"];
}> = (props) => {
  return (
    <Fragment>
      <Form.Item label="设置审批人" field="approver_getter" initialValue="user">
        <Radio.Group direction="vertical">
          <Radio value="user">指定成员</Radio>
          <Radio value="role">角色</Radio>
          <Radio value="field">表单中获取</Radio>
        </Radio.Group>
      </Form.Item>
      <Form.Item shouldUpdate noStyle>
        {
          // @ts-ignore
          (values: Record<string, any>) => {
            switch (values.approver_getter) {
              case "user":
                return (
                  <Form.Item
                    label="选择审批人"
                    field="approval_users"
                    shouldUpdate
                  >
                    <Select options={props.userOptions} mode="multiple" />
                  </Form.Item>
                );
              case "role":
                return (
                  <Form.Item label="选择角色" field="approval_roles">
                    <Select options={props.roleOptions} mode="multiple" />
                  </Form.Item>
                );
              case "field":
                return (
                  <Form.Item label="表单字段" field="approval_field">
                    <Input />
                  </Form.Item>
                );
              default:
                return null;
            }
          }
        }
      </Form.Item>
      <Form.Item
        label="多人审批时采用的审批方式"
        field="approval_method"
        initialValue="orsign"
      >
        <Radio.Group direction="vertical">
          <Radio value="orsign">或签</Radio>
          <Radio value="countersign">会签</Radio>
          {/* <Radio value="sequence">依次审批</Radio> */}
        </Radio.Group>
      </Form.Item>
    </Fragment>
  );
};
