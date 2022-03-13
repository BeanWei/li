import { Card } from "@arco-design/web-react";

const ListCard: React.FC = (props) => {
  return <Card>{props.children}</Card>;
};

export default ListCard;
