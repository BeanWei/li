import { toArr } from "@formily/shared";

export const getCurrentOptions = (values: any, dataSource: any) => {
  values = toArr(values).map((val) =>
    typeof val === "object" ? val.value : val
  );
  const findOptions = (options: any[]): any => {
    let current = [];
    for (const option of options) {
      if (values.includes(option.value)) {
        current.push(option);
      }
      const children = option.options;
      if (Array.isArray(children)) {
        current.push(...findOptions(children));
      }
    }
    return current;
  };
  return findOptions(dataSource);
};
