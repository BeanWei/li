import { MockMethod } from "vite-plugin-mock";
import getAppConfig from "./getAppConfig";
import getWelcomePageSchema from "./getWelcomePageSchema";
import getCurrentUser from "./getCurrentUser";

export default [
  {
    url: "/api/liql",
    method: "post",
    response: ({ body }) => {
      const { operation, variables } = body;
      if (operation === "@getAppConfig") {
        return {
          data: getAppConfig,
        };
      }
      if (operation === "@getCurrentUser") {
        return {
          data: getCurrentUser,
        };
      }
      if (operation == "@getAppView") {
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
