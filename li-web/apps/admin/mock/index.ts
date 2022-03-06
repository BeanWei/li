import { MockMethod } from "vite-plugin-mock";
import getAppMenuSchema from "./getAppMenuSchema";
import userGetProfile from "./userGetProfile";

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
      if (operation === "userGetProfile") {
        return {
          data: userGetProfile,
        };
      }
      return {
        code: 0,
        data: {},
      };
    },
  },
] as MockMethod[];
