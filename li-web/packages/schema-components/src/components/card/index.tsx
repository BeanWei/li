import { Card as ArcoCard } from "@arco-design/web-react";

export const Card: React.FC = (props) => {
  return (
    <ArcoCard style={{ marginBottom: 24 }} bordered={false} {...props}>
      {props.children}
    </ArcoCard>
  );
};
