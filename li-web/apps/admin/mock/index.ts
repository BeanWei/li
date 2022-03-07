import { MockMethod } from "vite-plugin-mock";
import getAppMenuSchema from "./getAppMenuSchema";
import getWelcomePageSchema from "./getWelcomePageSchema";
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
      if (operation == "getPageSchema") {
        if (variables.uid === "WelcomePage") {
          return {
            data: getWelcomePageSchema,
          };
        }
        return {
          code: 0,
          data: {},
        };
      }
      return {
        code: 0,
        data: {},
      };
    },
  },
] as MockMethod[];
