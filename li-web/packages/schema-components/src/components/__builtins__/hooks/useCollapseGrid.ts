import { useMemo } from "react";
import FormGrid from "../../form-grid";

export const useCollapseGrid = (maxRows: number) => {
  const grid = useMemo(
    () =>
      FormGrid.createFormGrid({
        maxColumns: 4,
        maxWidth: 240,
        maxRows: maxRows,
        shouldVisible: (node, grid) => {
          if (node.index === grid.childSize - 1) return true;
          if (grid.maxRows === Infinity) return true;
          return (node.shadowRow || 0) < maxRows + 1;
        },
      }),
    []
  );
  const expanded = grid.maxRows === Infinity;
  const realRows = grid.shadowRows;
  const computeRows = grid.fullnessLastColumn
    ? grid.shadowRows - 1
    : grid.shadowRows;

  const toggle = () => {
    if (grid.maxRows === Infinity) {
      grid.maxRows = maxRows;
    } else {
      grid.maxRows = Infinity;
    }
  };
  const takeType = () => {
    if (realRows < maxRows + 1) return "incomplete-wrap";
    if (computeRows > maxRows) return "collapsible";
    return "complete-wrap";
  };
  return {
    grid,
    expanded,
    toggle,
    type: takeType(),
  };
};
