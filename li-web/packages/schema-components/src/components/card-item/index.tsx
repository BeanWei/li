import { Card } from "@arco-design/web-react";

export const CardItem: React.FC = (props) => {
  return (
    <Card style={{ marginBottom: 24 }} bordered={false} {...props}>
      {props.children}
    </Card>
  );
};
