import { Card } from "@arco-design/web-react";
import React from "react";

const ListCard: React.FC<React.PropsWithChildren<{}>> = (props) => {
  return <Card>{props.children}</Card>;
};

export default ListCard;
