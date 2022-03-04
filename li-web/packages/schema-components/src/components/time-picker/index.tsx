import { connect, mapReadPretty } from "@formily/react";
import { TimePicker as ArcoTimePicker } from "@arco-design/web-react";
import {
  TimePickerProps as ArcoTimePickerProps,
  TimeRangePickerProps,
} from "@arco-design/web-react";
import { isArr } from "@formily/shared";

type ComposedTimePicker = React.FC<ArcoTimePickerProps> & {
  RangePicker?: React.FC<TimeRangePickerProps>;
};

const ReadPretty: React.FC<TimeRangePickerProps> = (props: any) => {
  const getLabels = () => {
    return isArr(props.value) ? props.value.join("~") : props.value;
  };
  return (
    <div className={props.className} style={props.style}>
      {getLabels()}
    </div>
  );
};

export const TimePicker: ComposedTimePicker = connect(
  ArcoTimePicker,
  mapReadPretty(ReadPretty)
);

TimePicker.RangePicker = connect(
  ArcoTimePicker.RangePicker,
  mapReadPretty(ReadPretty)
);

export default TimePicker;
