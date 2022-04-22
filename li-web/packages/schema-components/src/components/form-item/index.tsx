import { useEffect, useRef, useState } from "react";
import cls from "classnames";
import { Grid, Popover, Tooltip } from "@arco-design/web-react";
import {
  IconCheckCircle,
  IconCloseCircle,
  IconExclamationCircle,
  IconLoading,
  IconQuestionCircle,
} from "@arco-design/web-react/icon";
import { connect, mapProps } from "@formily/react";
import { isVoidField } from "@formily/core";
import { FormLayoutShallowContext, useFormLayout } from "../form-layout";
import { getPrefixCls, pickDataProps } from "../__builtins__";
import "./index.less";
import { Trans } from "react-i18next";

export interface IFormItemProps {
  className?: string;
  style?: React.CSSProperties;
  label?: React.ReactNode;
  colon?: boolean;
  tooltip?: React.ReactNode;
  tooltipIcon?: React.ReactNode;
  layout?: "vertical" | "horizontal" | "inline";
  tooltipLayout?: "icon" | "text";
  labelStyle?: React.CSSProperties;
  labelAlign?: "left" | "right";
  labelWrap?: boolean;
  labelWidth?: number | string;
  wrapperWidth?: number | string;
  labelCol?: number;
  wrapperCol?: number;
  wrapperAlign?: "left" | "right";
  wrapperWrap?: boolean;
  wrapperStyle?: React.CSSProperties;
  fullness?: boolean;
  addonBefore?: React.ReactNode;
  addonAfter?: React.ReactNode;
  size?: "small" | "default" | "large";
  extra?: React.ReactNode;
  feedbackText?: React.ReactNode;
  feedbackLayout?: "loose" | "terse" | "popover" | "none" | (string & {});
  feedbackStatus?: "error" | "warning" | "success" | "pending";
  feedbackIcon?: React.ReactNode;
  getPopupContainer?: (node: HTMLElement) => HTMLElement;
  asterisk?: boolean;
  gridSpan?: number;
}

type ComposeFormItem = React.FC<IFormItemProps> & {
  BaseItem?: React.FC<IFormItemProps>;
};

const useFormItemLayout = (props: IFormItemProps) => {
  const layout = useFormLayout();
  return {
    ...props,
    layout: props.layout ?? layout.layout ?? "horizontal",
    colon: props.colon ?? layout.colon,
    labelAlign:
      layout.layout === "vertical"
        ? props.labelAlign ?? layout.labelAlign ?? "left"
        : props.labelAlign ?? layout.labelAlign ?? "right",
    labelWrap: props.labelWrap ?? layout.labelWrap,
    labelWidth: props.labelWidth ?? layout.labelWidth,
    wrapperWidth: props.wrapperWidth ?? layout.wrapperWidth,
    labelCol: props.labelCol ?? layout.labelCol,
    wrapperCol: props.wrapperCol ?? layout.wrapperCol,
    wrapperAlign: props.wrapperAlign ?? layout.wrapperAlign,
    wrapperWrap: props.wrapperWrap ?? layout.wrapperWrap,
    fullness: props.fullness ?? layout.fullness,
    size: props.size ?? layout.size,
    asterisk: props.asterisk,
    feedbackIcon: props.feedbackIcon,
    feedbackLayout: props.feedbackLayout ?? layout.feedbackLayout ?? "loose",
    tooltipLayout: props.tooltipLayout ?? layout.tooltipLayout ?? "icon",
    tooltipIcon: props.tooltipIcon ?? layout.tooltipIcon ?? (
      <IconQuestionCircle />
    ),
  };
};

function useOverflow<
  Container extends HTMLElement,
  Content extends HTMLElement
>() {
  const [overflow, setOverflow] = useState(false);
  const containerRef = useRef<Container | any>();
  const contentRef = useRef<Content | any>();
  const layout = useFormLayout();
  const labelCol = JSON.stringify(layout.labelCol);

  useEffect(() => {
    requestAnimationFrame(() => {
      if (containerRef.current && contentRef.current) {
        const contentWidth = contentRef.current.getBoundingClientRect().width;
        const containerWidth =
          containerRef.current.getBoundingClientRect().width;
        if (contentWidth && containerWidth && containerWidth < contentWidth) {
          if (!overflow) setOverflow(true);
        } else {
          if (overflow) setOverflow(false);
        }
      }
    });
  }, [labelCol]);

  return {
    overflow,
    containerRef,
    contentRef,
  };
}

const ICON_MAP = {
  error: <IconCloseCircle />,
  success: <IconCheckCircle />,
  warning: <IconExclamationCircle />,
  pending: <IconLoading />,
};

export const BaseItem: React.FC<IFormItemProps> = ({ children, ...props }) => {
  const formLayout = useFormItemLayout(props);
  const { containerRef, contentRef, overflow } = useOverflow<
    HTMLDivElement,
    HTMLSpanElement
  >();
  const {
    label,
    style,
    layout,
    colon,
    addonBefore,
    addonAfter,
    asterisk,
    feedbackStatus,
    extra,
    feedbackText,
    fullness,
    feedbackLayout,
    feedbackIcon,
    getPopupContainer,
    labelWidth,
    wrapperWidth,
    labelCol,
    wrapperCol,
    labelAlign,
    wrapperAlign = "left",
    size,
    labelWrap,
    wrapperWrap,
    tooltipLayout,
    tooltip,
    tooltipIcon,
  } = formLayout;
  const labelStyle = { ...formLayout.labelStyle };
  const wrapperStyle = { ...formLayout.wrapperStyle };
  // 固定宽度
  let enableCol = false;
  if (labelWidth || wrapperWidth) {
    if (labelWidth) {
      labelStyle.width = labelWidth === "auto" ? undefined : labelWidth;
      labelStyle.maxWidth = labelWidth === "auto" ? undefined : labelWidth;
    }
    if (wrapperWidth) {
      wrapperStyle.width = wrapperWidth === "auto" ? undefined : wrapperWidth;
      wrapperStyle.maxWidth =
        wrapperWidth === "auto" ? undefined : wrapperWidth;
    }
    // 栅格模式
  }
  if (labelCol || wrapperCol) {
    if (!labelStyle.width && !wrapperStyle.width) {
      enableCol = true;
    }
  }

  const prefixCls = getPrefixCls();
  const formatChildren =
    feedbackLayout === "popover" && feedbackStatus ? (
      <Popover
        position="top"
        content={
          <div
            className={cls({
              [`${prefixCls}-form-item-${feedbackStatus}-help`]:
                !!feedbackStatus,
              [`${prefixCls}-form-item-help`]: true,
            })}
          >
            {ICON_MAP[feedbackStatus]} {feedbackText}
          </div>
        }
        popupVisible={!!feedbackText}
        getPopupContainer={getPopupContainer}
      >
        {children}
      </Popover>
    ) : (
      children
    );

  const gridStyles: React.CSSProperties = {};

  const getOverflowTooltip = () => {
    if (overflow) {
      return (
        <div>
          <div>{label}</div>
          <div>{tooltip}</div>
        </div>
      );
    }
    return tooltip;
  };

  const renderLabelText = () => {
    const labelChildren = (
      <label ref={contentRef}>
        {asterisk && (
          <strong className={`${prefixCls}-form-item-symbol`}>
            <svg
              fill="currentColor"
              viewBox="0 0 1024 1024"
              width="1em"
              height="1em"
            >
              <path d="M583.338667 17.066667c18.773333 0 34.133333 15.36 34.133333 34.133333v349.013333l313.344-101.888a34.133333 34.133333 0 0 1 43.008 22.016l42.154667 129.706667a34.133333 34.133333 0 0 1-21.845334 43.178667l-315.733333 102.4 208.896 287.744a34.133333 34.133333 0 0 1-7.509333 47.786666l-110.421334 80.213334a34.133333 34.133333 0 0 1-47.786666-7.509334L505.685333 706.218667 288.426667 1005.226667a34.133333 34.133333 0 0 1-47.786667 7.509333l-110.421333-80.213333a34.133333 34.133333 0 0 1-7.509334-47.786667l214.186667-295.253333L29.013333 489.813333a34.133333 34.133333 0 0 1-22.016-43.008l42.154667-129.877333a34.133333 34.133333 0 0 1 43.008-22.016l320.512 104.106667L412.672 51.2c0-18.773333 15.36-34.133333 34.133333-34.133333h136.533334z" />
            </svg>
          </strong>
        )}
        <Trans>{label}</Trans>
      </label>
    );

    if ((tooltipLayout === "text" && tooltip) || overflow) {
      return (
        <Tooltip position="top" content={getOverflowTooltip()}>
          {labelChildren}
        </Tooltip>
      );
    }
    return labelChildren;
  };

  const renderTooltipIcon = () => {
    if (tooltip && tooltipLayout === "icon" && !overflow) {
      return (
        <span className={`${prefixCls}-form-label-item-tooltip-icon`}>
          <Tooltip position="top" content={tooltip}>
            {tooltipIcon}
          </Tooltip>
        </span>
      );
    }
  };

  const renderLabel = () => {
    if (!label) return null;
    return (
      <Grid.Col
        ref={containerRef}
        span={enableCol && !!labelCol ? labelCol : undefined}
        className={cls(
          `${prefixCls}-form-item-label`,
          `${prefixCls}-form-label-item`,
          {
            [`${prefixCls}-form-label-item-tooltip`]:
              (tooltip && tooltipLayout === "text") || overflow,
            [`${prefixCls}-form-label-item-flex`]: !enableCol || !!!labelCol,
          }
        )}
        style={labelStyle}
      >
        {renderLabelText()}
        {renderTooltipIcon()}
        <span className={`${prefixCls}-form-item-colon`}>
          {colon ? ":" : ""}
        </span>
      </Grid.Col>
    );
  };

  return (
    <Grid.Row
      {...pickDataProps(props)}
      div={layout !== "horizontal"}
      data-grid-span={props.gridSpan}
      className={cls({
        [`${prefixCls}-form-item`]: true,
        [`${prefixCls}-form-item-layout-${layout}`]: true,
        [`${prefixCls}-form-item-${feedbackStatus}`]: !!feedbackStatus,
        [`${prefixCls}-form-item-feedback-has-text`]: !!feedbackText,
        [`${prefixCls}-form-item-size-${size}`]: !!size,
        [`${prefixCls}-form-item-feedback-layout-${feedbackLayout}`]:
          !!feedbackLayout,
        [`${prefixCls}-form-item-fullness`]: !!fullness || !!feedbackIcon,
        [`${prefixCls}-form-item-label-align-${labelAlign}`]: true,
        [`${prefixCls}-form-item-control-align-${wrapperAlign}`]: true,
        [`${prefixCls}-form-item-label-wrap`]: !!labelWrap,
        [`${prefixCls}-form-item-control-wrap`]: !!wrapperWrap,
        [props.className || ""]: !!props.className,
      })}
      style={{
        ...style,
        ...gridStyles,
      }}
    >
      {renderLabel()}
      <Grid.Col
        className={cls(`${prefixCls}-form-item-wrapper`, {
          [`${prefixCls}-item-wrapper-flex`]:
            !enableCol || !!!wrapperCol || !label,
        })}
        span={enableCol && !!wrapperCol ? wrapperCol : undefined}
      >
        <div className={cls(`${prefixCls}-form-item-control-content`)}>
          {addonBefore && (
            <div className={cls(`${prefixCls}-form-item-addon-before`)}>
              {addonBefore}
            </div>
          )}
          <div
            style={wrapperStyle}
            className={cls({
              [`${prefixCls}-form-item-control-content-component`]: true,
              [`${prefixCls}-form-item-control-content-component-has-feedback-icon`]:
                !!feedbackIcon,
            })}
          >
            <FormLayoutShallowContext.Provider value={undefined}>
              {formatChildren}
            </FormLayoutShallowContext.Provider>
            {feedbackIcon && (
              <div className={cls(`${prefixCls}-form-item-feedback-icon`)}>
                {feedbackIcon}
              </div>
            )}
          </div>
          {addonAfter && (
            <div className={cls(`${prefixCls}-form-item-addon-after`)}>
              {addonAfter}
            </div>
          )}
        </div>
        {!!feedbackText &&
          feedbackLayout !== "popover" &&
          feedbackLayout !== "none" && (
            <div
              className={cls({
                [`${prefixCls}-form-message`]: true,
                [`${prefixCls}-form-item-${feedbackStatus}-help`]:
                  !!feedbackStatus,
                [`${prefixCls}-form-item-help`]: true,
                [`${prefixCls}-form-item-help-enter`]: true,
                [`${prefixCls}-form-item-help-enter-active`]: true,
              })}
            >
              {feedbackText}
            </div>
          )}
        {extra && (
          <div className={cls(`${prefixCls}-form-item-extra`)}>{extra}</div>
        )}
      </Grid.Col>
    </Grid.Row>
  );
};

// 适配
export const FormItem: ComposeFormItem = connect(
  BaseItem,
  mapProps((props, field) => {
    if (isVoidField(field))
      return {
        label: field.title || props.label,
        asterisk: props.asterisk,
        extra: props.extra || field.description,
      };
    if (!field) return props;
    const takeFeedbackStatus = () => {
      if (field.validating) return "pending";
      return field.decoratorProps.feedbackStatus || field.validateStatus;
    };
    const takeMessage = () => {
      const split = (messages: any[]) => {
        return messages.reduce((buf, text, index) => {
          if (!text) return buf;
          return index < messages.length - 1
            ? buf.concat([text, ", "])
            : buf.concat([text]);
        }, []);
      };
      if (field.validating) return;
      if (props.feedbackText) return props.feedbackText;
      if (field.selfErrors.length) return split(field.selfErrors);
      if (field.selfWarnings.length) return split(field.selfWarnings);
      if (field.selfSuccesses.length) return split(field.selfSuccesses);
    };
    const takeAsterisk = () => {
      if (field.required && field.pattern !== "readPretty") {
        return true;
      }
      if ("asterisk" in props) {
        return props.asterisk;
      }
      return false;
    };
    return {
      label: props.label || field.title,
      feedbackStatus: takeFeedbackStatus(),
      feedbackText: takeMessage(),
      asterisk: takeAsterisk(),
      extra: props.extra || field.description,
    };
  })
);

FormItem.BaseItem = BaseItem;

export default FormItem;
