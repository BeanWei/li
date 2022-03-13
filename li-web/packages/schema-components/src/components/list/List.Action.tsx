import { Button, Popconfirm } from "@arco-design/web-react";
import { useField } from "@formily/react";
import { Fragment, useContext } from "react";
import { ListContext } from "./context";
import { ComposedListAction } from "./types";

export const ListAction: ComposedListAction = () => {
  return <Fragment />;
};

ListAction.BulkDelete = (props) => {
  const field = useField();
  const ctx = useContext(ListContext);

  return (
    <Popconfirm
      title="Are you sure you want to delete?"
      {...props}
      onOk={() => {
        console.log(ctx.selectedRowKeys);
      }}
    >
      <Button {...props} disabled={!!!ctx.selectedRowKeys?.length}>
        {field.title}
      </Button>
    </Popconfirm>
  );
};
