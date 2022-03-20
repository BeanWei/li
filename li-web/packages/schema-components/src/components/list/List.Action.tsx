import {
  Button,
  Input,
  Popconfirm,
  Popover,
  Select,
  Space,
} from "@arco-design/web-react";
import { IconFilter, IconRefresh } from "@arco-design/web-react/icon";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { request } from "pro-utils";
import { useContext } from "react";
import Form from "../form";
import FormButtonGroup from "../form-button-group";
import FormGrid from "../form-grid";
import Submit from "../submit";
import { ListContext } from "./context";
import { ComposedListAction } from "./types";

export const ListAction: ComposedListAction = observer((props) => {
  const fieldSchema = useFieldSchema();
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        width: "100%",
        marginBottom: 9,
      }}
    >
      <Space>
        {fieldSchema.mapProperties((schema, key) => {
          if (schema["x-component-props"]?.["position"] !== "left") {
            return null;
          }
          return <RecursionField key={key} name={key} schema={schema} />;
        })}
      </Space>
      <Space>
        {fieldSchema.mapProperties((schema, key) => {
          if (schema["x-component-props"]?.["position"] === "left") {
            return null;
          }
          return <RecursionField key={key} name={key} schema={schema} />;
        })}
      </Space>
    </div>
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

ListAction.RowSelection = observer((props) => {
  const { confirmProps, forSubmit, afterReload, ...rest } = props;
  const field = useField();
  const ctx = useContext(ListContext);
  const handleOk = () => {
    if (forSubmit) {
      request(forSubmit, { ids: ctx.selectedRowKeys }).then(() => {
        afterReload && ctx.result?.run();
      });
    }
  };
  if (props.confirmProps) {
    return (
      <Popconfirm {...props.confirmProps} onOk={handleOk}>
        <Button {...rest} disabled={!!!ctx.selectedRowKeys?.length}>
          {field.title}
        </Button>
      </Popconfirm>
    );
  }
  return (
    <Button
      {...rest}
      disabled={!!!ctx.selectedRowKeys?.length}
      onClick={handleOk}
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
