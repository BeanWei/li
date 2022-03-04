import { Card as ArcoCard } from "@arco-design/web-react";
import '@arco-design/web-react/lib/Card/style/index';

export const Card: React.FC = (props) => {
  return (
    <ArcoCard style={{ marginBottom: 24 }} bordered={false} {...props}>
      {props.children}
    </ArcoCard>
  );
};
