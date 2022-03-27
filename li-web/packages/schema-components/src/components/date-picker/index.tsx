import React from "react";
import { DatePicker as ArcoDatePicker } from "@arco-design/web-react";
import type { DatePickerProps, RangePickerProps } from "@arco-design/web-react";
import { connect, mapReadPretty } from "@formily/react";
import { isArr } from "@formily/shared";
import cls from "classnames";
import { usePrefixCls } from "../__builtins__";

type ComposedDatePicker = React.FC<DatePickerProps> & {
  RangePicker?: React.FC<RangePickerProps>;
};

const ReadPretty: React.FC = (props: any) => {
  if (!props.value) {
    return <div></div>;
  }
  const prefixCls = usePrefixCls("description-date-picker", props);
  const getLabels = () => {
    return isArr(props.value) ? props.value.join("~") : props.value;
  };
  return <div className={cls(prefixCls, props.className)}>{getLabels()}</div>;
};

const BaseDatePicker: React.FC<any> = (props) => {
  const { mode, ...rest } = props;
  switch (mode) {
    case "date":
      return <ArcoDatePicker {...rest} />;
    case "week":
      return <ArcoDatePicker.WeekPicker {...rest} />;
    case "month":
      return <ArcoDatePicker.MonthPicker {...rest} />;
    case "year":
      return <ArcoDatePicker.YearPicker {...rest} />;
    case "quarter":
      return <ArcoDatePicker.QuarterPicker {...rest} />;
    default:
      return <ArcoDatePicker {...rest} />;
  }
};

export const DatePicker: ComposedDatePicker = connect(
  BaseDatePicker,
  mapReadPretty(ReadPretty)
);

export default DatePicker;
