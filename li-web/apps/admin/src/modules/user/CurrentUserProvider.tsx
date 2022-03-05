import { createContext, useContext } from "react";
import { Spin } from "@arco-design/web-react";
import { Redirect } from "react-router-dom";
import { useRequest } from "pro-utils";

export const CurrentUserContext = createContext<any>({});

export const useCurrentUserContext = () => {
  return useContext(CurrentUserContext);
};

export const CurrentUserProvider = (props: any) => {
  const result = useRequest({
    url: "getSessionUser",
  });
  if (result.loading) {
    return <Spin />;
  }
  if (result.error) {
    return <Redirect to={"/signin"} />;
  }
  return (
    <CurrentUserContext.Provider value={result}>
      {props.children}
    </CurrentUserContext.Provider>
  );
};
