import { Divider as ArcoDivider } from "@arco-design/web-react";

export const Divider: React.FC = (props) => {
  return (
    <ArcoDivider style={{ marginBottom: 24 }} {...props}>
      {props.children}
    </ArcoDivider>
  );
};

export default Divider;
