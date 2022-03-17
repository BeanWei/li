import { MockMethod } from "vite-plugin-mock";
import getAppConfig from "./getAppConfig";
import getWelcomePageSchema from "./getWelcomePageSchema";
import userGetProfile from "./userGetProfile";

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
      if (operation === "userGetProfile") {
        return {
          data: userGetProfile,
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
