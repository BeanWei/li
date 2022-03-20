import { MockMethod } from "vite-plugin-mock";
import getAppConfig from "./getAppConfig";
import getCurrentUser from "./getCurrentUser";
import getSystemUserView from "./getSystemUserView";
import getWelcomeView from "./getWelcomeView";
import listUser from "./listUser";

const views = {
  Welcome: getWelcomeView,
  SystemUser: getSystemUserView,
};

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
        return {
          data: views[variables.key] || {},
        };
      }
      if (operation == "listUser") {
        return {
          data: listUser,
          total: listUser.length,
        };
      }
      return {
        code: 0,
        data: {},
      };
    },
  },
] as MockMethod[];
