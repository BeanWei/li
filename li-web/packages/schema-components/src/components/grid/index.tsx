import React from "react";
import { Grid as ArcoGrid, ColProps, RowProps } from "@arco-design/web-react";

type ComposedGrid = React.FC<{}> & {
  Row?: React.FC<RowProps>;
  Col?: React.FC<ColProps>;
};

export const Grid: ComposedGrid = () => {
  return <></>;
};

Grid.Row = (props) => {
  return <ArcoGrid.Row {...props} />;
};

Grid.Col = (props) => {
  return <ArcoGrid.Col {...props} />;
};

export default Grid;
