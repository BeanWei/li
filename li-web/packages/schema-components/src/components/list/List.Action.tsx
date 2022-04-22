import { useContext } from "react";
import { Button, Input, Popconfirm, Space } from "@arco-design/web-react";
import {
  observer,
  RecursionField,
  useField,
  useFieldSchema,
} from "@formily/react";
import { Trans, useTranslation } from "react-i18next";
import { request } from "pro-utils";
import { useRecord } from "../../core";
import ActionFormDrawer from "../action/Action.FormDrawer";
import ActionFormModal from "../action/Action.FormModal";
import { getLocale, Icon } from "../__builtins__";
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
        marginBottom: 16,
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

ListAction.RowSelection = observer((props) => {
  const { confirmProps, forSubmit, afterReload, ...rest } = props;
  const field = useField();
  const ctx = useContext(ListContext);
  const { t } = useTranslation();
  const handleOk = () => {
    if (forSubmit) {
      request(forSubmit, { ids: ctx.selectedKeys }).then(() => {
        afterReload && ctx.result?.run();
      });
    }
  };
  if (props.confirmProps) {
    return (
      <Popconfirm {...props.confirmProps} onOk={handleOk}>
        <Button
          {...rest}
          disabled={!!!ctx.selectedKeys?.length}
          icon={
            props.icon && typeof props.icon === "string" ? (
              <Icon type={props.icon} />
            ) : (
              props.icon
            )
          }
        >
          {t(field.title)}
        </Button>
      </Popconfirm>
    );
  }
  return (
    <Button
      {...rest}
      disabled={!!!ctx.selectedKeys?.length}
      onClick={handleOk}
      icon={
        props.icon && typeof props.icon === "string" ? (
          <Icon type={props.icon} />
        ) : (
          props.icon
        )
      }
    >
      {t(field.title)}
    </Button>
  );
});

ListAction.Reload = observer((props) => {
  const { data, ...rest } = props;
  const field = useField();
  const ctx = useContext(ListContext);
  const { t } = useTranslation();
  return (
    <Button
      {...rest}
      icon={
        props.icon && typeof props.icon === "string" ? (
          <Icon type={props.icon} />
        ) : (
          props.icon
        )
      }
      onClick={() => {
        ctx.reload?.(data);
      }}
    >
      {t(field.title)}
    </Button>
  );
});

ListAction.RecordFormDrawer = observer((props) => {
  const ctx = useContext(ListContext);
  const forInitVariables = useRecord();
  return (
    <ActionFormDrawer
      {...props}
      forInitVariables={forInitVariables}
      forSubmitSuccess={() => {
        ctx.result?.run();
      }}
    />
  );
});

ListAction.RecordFormModal = observer((props) => {
  const ctx = useContext(ListContext);
  const forInitVariables = useRecord();
  return (
    <ActionFormModal
      {...props}
      forInitVariables={forInitVariables}
      forSubmitSuccess={() => {
        ctx.result?.run();
      }}
    />
  );
});

ListAction.RecordDelete = observer((props) => {
  const { confirmProps, forSubmit, ...rest } = props;
  const field = useField();
  const local = getLocale();
  const ctx = useContext(ListContext);
  const variables = useRecord();
  const { t } = useTranslation();
  const handleOk = () => {
    if (forSubmit) {
      request(forSubmit, variables).then(() => {
        ctx.result?.run();
      });
    }
  };
  return (
    <Popconfirm
      {...props.confirmProps}
      title={
        props.confirmProps?.title ? (
          <Trans>{props.confirmProps?.title}</Trans>
        ) : (
          local.List.confirmDelete
        )
      }
      onOk={handleOk}
    >
      <Button
        {...rest}
        icon={
          props.icon && typeof props.icon === "string" ? (
            <Icon type={props.icon} />
          ) : (
            props.icon
          )
        }
      >
        {t(field.title)}
      </Button>
    </Popconfirm>
  );
});

ListAction.Search = (props) => {
  const ctx = useContext(ListContext);
  const local = getLocale();
  const { t } = useTranslation();
  return (
    <Input.Search
      style={props.style}
      placeholder={props.placeholder ? t(props.placeholder) : local.List.search}
      onSearch={(value) => {
        ctx.reload?.({
          query: value,
        });
      }}
    />
  );
};
