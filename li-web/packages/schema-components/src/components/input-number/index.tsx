import { InputNumber as ArcoInputNumber } from "@arco-design/web-react";
import { connect, mapReadPretty } from "@formily/react";
import { isValid } from "@formily/shared";
import { toFixed } from "./utils/MiniDecimal";
import { getNumberPrecision } from "./utils/numberUtil";

export const InputNumber = connect(
  ArcoInputNumber,
  mapReadPretty((props) => {
    const { step, value, addonBefore, addonAfter } = props;
    if (!isValid(props.value)) {
      return <div></div>;
    }
    const precision = Math.max(
      getNumberPrecision(String(value)),
      getNumberPrecision(step)
    );
    return (
      <div>
        {addonBefore}
        {toFixed(String(value), ".", precision)}
        {addonAfter}
      </div>
    );
  })
);

export default InputNumber;
