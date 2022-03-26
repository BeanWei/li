import { InputNumber, InputNumberProps } from "@arco-design/web-react";
import { connect, mapReadPretty } from "@formily/react";
import { isValid } from "@formily/shared";

const getFormatter = (currency?: string): Intl.NumberFormat => {
  if (currency === "USD") {
    return new Intl.NumberFormat("en-US", {
      style: "currency",
      currency: "USD",
    });
  }
  return new Intl.NumberFormat("zh-CN", { style: "currency", currency: "CNY" });
};

const getPrefix = (currency?: string): string => {
  if (currency === "USD") {
    return "$";
  }
  return "¥";
};

const InputMoney: React.FC<InputNumberProps & { currency?: string }> = (
  props
) => {
  const { currency, ...rest } = props;
  return (
    <InputNumber
      {...rest}
      min={0}
      prefix={getPrefix(currency)}
      formatter={(value: any) => {
        if (!value) return "";
        return `${
          value.toString().split(".").length === 2
            ? Number(value).toFixed(2)
            : Number(value)
        }`.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
      }}
      parser={(value: string) => value.replace(/,/g, "")}
    />
  );
};

export const Money = connect(
  InputMoney,
  mapReadPretty((props) => {
    const { value, currency } = props;
    if (!isValid(props.value)) {
      return <div></div>;
    }
    const formatter = getFormatter(currency);
    return <div>{formatter.format(value)}</div>;
  })
);

export default Money;
