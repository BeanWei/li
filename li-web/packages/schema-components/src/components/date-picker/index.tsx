import React from "react";
import {
  DatePicker as ArcoDatePicker,
  MonthPickerProps,
  QuarterPickerProps,
  WeekPickerProps,
  YearPickerProps,
} from "@arco-design/web-react";
import type { DatePickerProps, RangePickerProps } from "@arco-design/web-react";
import { connect, mapReadPretty } from "@formily/react";
import { isArr } from "@formily/shared";
import cls from "classnames";
import { usePrefixCls } from "../__builtins__";
import "@arco-design/web-react/lib/DatePicker/style/index";

type ComposedDatePicker = React.FC<DatePickerProps> & {
  WeekPicker?: React.FC<WeekPickerProps>;
  MonthPicker?: React.FC<MonthPickerProps>;
  YearPicker?: React.FC<YearPickerProps>;
  QuarterPicker?: React.FC<QuarterPickerProps>;
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

export const DatePicker: ComposedDatePicker = connect(
  ArcoDatePicker,
  mapReadPretty(ReadPretty)
);

DatePicker.WeekPicker = connect(
  ArcoDatePicker.WeekPicker,
  mapReadPretty(ReadPretty)
);

DatePicker.MonthPicker = connect(
  ArcoDatePicker.MonthPicker,
  mapReadPretty(ReadPretty)
);

DatePicker.YearPicker = connect(
  ArcoDatePicker.YearPicker,
  mapReadPretty(ReadPretty)
);

DatePicker.QuarterPicker = connect(
  ArcoDatePicker.QuarterPicker,
  mapReadPretty(ReadPretty)
);

DatePicker.RangePicker = connect(
  ArcoDatePicker.RangePicker,
  mapReadPretty(ReadPretty)
);

export default DatePicker;
