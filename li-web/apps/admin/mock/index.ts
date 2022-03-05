import { MockMethod } from "vite-plugin-mock";

export default [
  {
    url: "/api/liql",
    method: "post",
    response: ({ body }) => {
      const { operation, variables } = body;
      if (operation === "getAppMenuSchema") {
        return {
          data: getAppMenuSchema,
        };
      }
      return {
        code: 0,
        data: {},
      };
    },
  },
] as MockMethod[];
