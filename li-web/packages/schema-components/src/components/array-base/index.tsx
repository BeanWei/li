import React, { createContext, useContext } from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import {
  IconDelete,
  IconDown,
  IconUp,
  IconPlus,
  IconMenu,
  IconProps,
} from "@arco-design/web-react/icon";
import { ArrayField } from "@formily/core";
import {
  useField,
  useFieldSchema,
  Schema,
  JSXComponent,
  ExpressionScope,
} from "@formily/react";
import { isValid, clone, isBool } from "@formily/shared";
import { SortableHandle } from "react-sortable-hoc";
import cls from "classnames";
import { useTranslation } from "react-i18next";
import { usePrefixCls } from "../__builtins__";
import "./index.less";

export interface IArrayBaseAdditionProps extends ButtonProps {
  title?: string;
  method?: "push" | "unshift";
  defaultValue?: any;
}

export interface IArrayBaseContext {
  props: IArrayBaseProps;
  field: ArrayField;
  schema: Schema;
}

export interface IArrayBaseItemProps {
  index: number;
  record: any;
}

export type ArrayBaseMixins = {
  Addition?: React.FC<React.PropsWithChildren<IArrayBaseAdditionProps>>;
  Remove?: React.FC<React.PropsWithChildren<IconProps & { index?: number }>>;
  MoveUp?: React.FC<React.PropsWithChildren<IconProps & { index?: number }>>;
  MoveDown?: React.FC<React.PropsWithChildren<IconProps & { index?: number }>>;
  SortHandle?: React.FC<
    React.PropsWithChildren<IconProps & { index?: number }>
  >;
  Index?: React.FC;
  useArray?: () => IArrayBaseContext | null;
  useIndex?: () => number | undefined;
  useRecord?: () => any;
};

export interface IArrayBaseProps {
  disabled?: boolean;
  onAdd?: (index: number) => void;
  onRemove?: (index: number) => void;
  onMoveDown?: (index: number) => void;
  onMoveUp?: (index: number) => void;
}

type ComposedArrayBase = React.FC<React.PropsWithChildren<IArrayBaseProps>> &
  ArrayBaseMixins & {
    Item: React.FC<React.PropsWithChildren<IArrayBaseItemProps>>;
    mixin: <T extends JSXComponent>(target: T) => T & ArrayBaseMixins;
  };

const ArrayBaseContext = createContext<IArrayBaseContext | null>(null);

const ItemContext = createContext<IArrayBaseItemProps | null>(null);

const useArray = () => {
  return useContext(ArrayBaseContext);
};

const useIndex = (index?: number) => {
  const ctx = useContext(ItemContext);
  return ctx ? ctx.index : index;
};

const useRecord = (record?: number) => {
  const ctx = useContext(ItemContext);
  return ctx ? ctx.record : record;
};

const getDefaultValue = (defaultValue: any, schema: Schema): any => {
  if (isValid(defaultValue)) return clone(defaultValue);
  if (Array.isArray(schema?.items))
    return getDefaultValue(defaultValue, schema.items[0]);
  if (schema?.items?.type === "array") return [];
  if (schema?.items?.type === "boolean") return true;
  if (schema?.items?.type === "date") return "";
  if (schema?.items?.type === "datetime") return "";
  if (schema?.items?.type === "number") return 0;
  if (schema?.items?.type === "object") return {};
  if (schema?.items?.type === "string") return "";
  return null;
};

export const ArrayBase: ComposedArrayBase = (props) => {
  const field = useField<ArrayField>();
  const schema = useFieldSchema();
  return (
    <ArrayBaseContext.Provider value={{ field, schema, props }}>
      {props.children}
    </ArrayBaseContext.Provider>
  );
};

ArrayBase.Item = ({ children, ...props }) => {
  return (
    <ItemContext.Provider value={props}>
      <ExpressionScope value={{ $record: props.record, $index: props.index }}>
        {children}
      </ExpressionScope>
    </ItemContext.Provider>
  );
};

const SortHandle = SortableHandle((props: any) => {
  const prefixCls = usePrefixCls("formily-array-base");
  return (
    <IconMenu
      {...props}
      className={cls(`${prefixCls}-sort-handle`, props.className)}
      style={{ ...props.style }}
    />
  );
}) as any;

ArrayBase.SortHandle = (props) => {
  const array = useArray();
  if (!array) return null;
  if (array.field?.pattern !== "editable") return null;
  return <SortHandle {...props} />;
};

ArrayBase.Index = (props) => {
  const index = useIndex() || 0;
  const prefixCls = usePrefixCls("formily-array-base");
  return (
    <span {...props} className={`${prefixCls}-index`}>
      #{index + 1}.
    </span>
  );
};

ArrayBase.Addition = (props) => {
  const self = useField();
  const array = useArray();
  const { t } = useTranslation();
  const prefixCls = usePrefixCls("formily-array-base");
  if (!array) return null;
  if (
    array.field?.pattern !== "editable" &&
    array.field?.pattern !== "disabled"
  )
    return null;
  return (
    <Button
      type="dashed"
      long
      {...props}
      disabled={isBool(self?.disabled) ? self?.disabled : array.field?.disabled}
      className={cls(`${prefixCls}-addition`, props.className)}
      onClick={(e) => {
        if (array.props?.disabled) return;
        const defaultValue = getDefaultValue(props.defaultValue, array.schema);
        if (props.method === "unshift") {
          array.field?.unshift?.(defaultValue);
          array.props?.onAdd?.(0);
        } else {
          array.field?.push?.(defaultValue);
          array.props?.onAdd?.(array?.field?.value?.length - 1);
        }
        if (props.onClick) {
          props.onClick(e);
        }
      }}
      icon={<IconPlus />}
    >
      {t(props.title || self.title)}
    </Button>
  );
};

ArrayBase.Remove = React.forwardRef((props, ref) => {
  const index = useIndex(props.index) || 0;
  const array = useArray();
  const prefixCls = usePrefixCls("formily-array-base");
  if (!array) return null;
  if (array.field?.pattern !== "editable") return null;
  return (
    <IconDelete
      {...props}
      className={cls(`${prefixCls}-remove`, props.className)}
      ref={ref}
      onClick={(e) => {
        if (array.props?.disabled) return;
        e.stopPropagation();
        array.field?.remove?.(index);
        array.props?.onRemove?.(index);
        if (props.onClick) {
          props.onClick(e);
        }
      }}
    />
  );
});

ArrayBase.MoveDown = React.forwardRef((props, ref) => {
  const index = useIndex(props.index) || 0;
  const array = useArray();
  const prefixCls = usePrefixCls("formily-array-base");
  if (!array) return null;
  if (array.field?.pattern !== "editable") return null;
  return (
    <IconDown
      {...props}
      className={cls(`${prefixCls}-move-down`, props.className)}
      ref={ref}
      onClick={(e) => {
        if (array.props?.disabled) return;
        e.stopPropagation();
        array.field?.moveDown?.(index);
        array.props?.onMoveDown?.(index);
        if (props.onClick) {
          props.onClick(e);
        }
      }}
    />
  );
});

ArrayBase.MoveUp = React.forwardRef((props, ref) => {
  const index = useIndex(props.index) || 0;
  const array = useArray();
  const prefixCls = usePrefixCls("formily-array-base");
  if (!array) return null;
  if (array.field?.pattern !== "editable") return null;
  return (
    <IconUp
      {...props}
      className={cls(`${prefixCls}-move-up`, props.className)}
      ref={ref}
      onClick={(e) => {
        if (array.props?.disabled) return;
        e.stopPropagation();
        array?.field?.moveUp(index);
        array?.props?.onMoveUp?.(index);
        if (props.onClick) {
          props.onClick(e);
        }
      }}
    />
  );
});

ArrayBase.useArray = useArray;
ArrayBase.useIndex = useIndex;
ArrayBase.useRecord = useRecord;
ArrayBase.mixin = (target: any) => {
  target.Index = ArrayBase.Index;
  target.SortHandle = ArrayBase.SortHandle;
  target.Addition = ArrayBase.Addition;
  target.Remove = ArrayBase.Remove;
  target.MoveDown = ArrayBase.MoveDown;
  target.MoveUp = ArrayBase.MoveUp;
  target.useArray = ArrayBase.useArray;
  target.useIndex = ArrayBase.useIndex;
  target.useRecord = ArrayBase.useRecord;
  return target;
};

export default ArrayBase;
