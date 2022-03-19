import { Button, Input, Popover, Select } from "@arco-design/web-react";
import { IconFilter, IconRefresh } from "@arco-design/web-react/icon";
import { observer } from "@formily/react";
import { useContext } from "react";
import { Action } from "../action";
import Form from "../form";
import FormButtonGroup from "../form-button-group";
import FormGrid from "../form-grid";
import Submit from "../submit";
import { ListContext } from "./context";
import { ComposedListAction } from "./types";

export const ListAction: ComposedListAction = observer((props) => {
  return (
    <div
      {...props}
      style={{
        marginBottom: 9,
      }}
    />
  );
});

ListAction.FilterGroup = observer((props) => {
  return (
    <Popover
      trigger="click"
      content={
        <Form>
          <FormGrid>{props.children}</FormGrid>
          <FormButtonGroup align="right">
            <Submit onSubmit={console.log}>查询</Submit>
          </FormButtonGroup>
        </Form>
      }
    >
      <Button icon={<IconFilter />} />
    </Popover>
  );
});

ListAction.FilterSelect = observer((props) => {
  return (
    <Select
      {...props}
      mode="multiple"
      maxTagCount={1}
      placeholder="Quick filter"
      style={{ width: 150 }}
      allowClear
    >
      {["Beijing", "Shanghai", "Guangzhou", "Shenzhen", "Chengdu", "Wuhan"].map(
        (option) => (
          <Select.Option key={option} value={option}>
            {option}
          </Select.Option>
        )
      )}
    </Select>
  );
});

ListAction.Search = observer((props) => {
  return (
    <Input.Search
      allowClear
      placeholder="Search"
      style={{ width: 150 }}
      {...props}
    />
  );
});

ListAction.RowSelection = observer((props) => {
  const ctx = useContext(ListContext);
  return (
    <Action
      {...props}
      disabled={!!!ctx.selectedRowKeys?.length}
      onClick={() => {
        console.log(ctx.selectedRowKeys);
      }}
    />
  );
});

ListAction.Refresh = observer((props) => {
  const ctx = useContext(ListContext);
  return (
    <Button
      icon={<IconRefresh />}
      {...props}
      onClick={() => {
        ctx.result?.run();
      }}
    />
  );
});

ListAction.BulkDelete = observer((props) => {
  const ctx = useContext(ListContext);
  return (
    <Action
      {...props}
      confirm={{
        title: "Confirm deletion",
        content: `Are you sure you want to delete the ${ctx.selectedRowKeys?.length} selected items? Once you press the delete button, the items will be deleted immediately. You can't undo this action.`,
        okButtonProps: { status: "danger" },
      }}
      disabled={!!!ctx.selectedRowKeys?.length}
      onClick={() => {
        console.log(ctx.selectedRowKeys);
      }}
    />
  );
});
