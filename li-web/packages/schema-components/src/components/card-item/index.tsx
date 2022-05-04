import { Card } from "@arco-design/web-react";

export const CardItem: React.FC<React.PropsWithChildren<{}>> = (props) => {
  return (
    <Card style={{ marginBottom: 24 }} bordered={false} {...props}>
      {props.children}
    </Card>
  );
};

export default CardItem;
