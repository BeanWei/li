import React, { useEffect, useState } from "react";
import { connect, mapProps, mapReadPretty } from "@formily/react";
import {
  Input as ArcoInput,
  InputProps,
  Popover,
  TextAreaProps,
} from "@arco-design/web-react";
import { IconLoading } from "@arco-design/web-react/icon";
import cls from "classnames";
import { usePrefixCls } from "../__builtins__";
import { useCompile } from "../../hooks";
import "./index.less";

type ComposedInput = React.FC<InputProps> & {
  TextArea?: React.FC<TextAreaProps>;
};

export const Input: ComposedInput = connect(
  ArcoInput,
  mapProps((props, field: any) => {
    return {
      ...props,
      suffix: (
        <span>
          {field?.["loading"] || field?.["validating"] ? (
            <IconLoading />
          ) : (
            props.suffix
          )}
        </span>
      ),
    };
  }),
  mapReadPretty((props) => {
    const prefixCls = usePrefixCls("description-input", props);
    const domRef = React.useRef<HTMLInputElement>(null);
    const compile = useCompile();
    const [ellipsis, setEllipsis] = useState(false);
    const content = compile(props.value);
    const ellipsisContent = (
      <Popover content={content} style={{ width: 100 }}>
        <span className={"input-ellipsis"}>{content}</span>
      </Popover>
    );
    useEffect(() => {
      if (
        domRef.current?.scrollWidth &&
        domRef.current?.scrollWidth > domRef.current?.clientWidth
      ) {
        setEllipsis(true);
      }
    }, []);

    return (
      <div className={cls(prefixCls, props.className)} style={props.style}>
        {props.addonBefore}
        {props.prefix}
        <span ref={domRef}>{ellipsis ? ellipsisContent : content}</span>
        {props.suffix}
        {props.addonAfter}
      </div>
    );
  })
);

Input.TextArea = connect(
  ArcoInput.TextArea,
  mapReadPretty((props) => {
    const prefixCls = usePrefixCls("description-textarea", props);
    const domRef = React.useRef<HTMLInputElement>(null);
    const [ellipsis, setEllipsis] = useState(false);
    const ellipsisProp = props.ellipsis === true ? {} : props.ellipsis;
    const ellipsisContent = (
      <Popover content={props.value}>
        <span
          className={"input-ellipsis"}
          style={{
            ...ellipsisProp,
          }}
        >
          {props.text || props.value}
        </span>
      </Popover>
    );
    useEffect(() => {
      if (
        domRef.current?.scrollWidth &&
        domRef.current?.scrollWidth > domRef.current?.clientWidth
      ) {
        setEllipsis(true);
      }
    }, []);
    return (
      <div className={cls(prefixCls, props.className)} style={props.style}>
        {props.addonBefore}
        {props.prefix}
        <span ref={domRef}>{ellipsis ? ellipsisContent : props.value}</span>
        {props.suffix}
        {props.addonAfter}
      </div>
    );
  })
);

export default Input;
