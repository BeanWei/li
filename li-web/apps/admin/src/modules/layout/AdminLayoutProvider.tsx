import { Spin } from "@arco-design/web-react";
import { useRequest } from "pro-utils";
import { createContext, useContext } from "react";
import { Redirect } from "react-router";

export const AdminLayoutContext = createContext<{
  lang?: string;
  setLang?: (value: string) => void;
  theme?: string;
  setTheme?: (value: string) => void;
  profile?: Record<string, any>;
  setProfile?: (value: Record<string, any>) => void;
}>({});

export const useAdminLayoutContext = () => {
  return useContext(AdminLayoutContext);
};

export const AdminLayoutProvider = (props: any) => {
  const result = useRequest("userGetProfile");
  if (result.loading) {
    return <Spin />;
  }
  if (result.error) {
    return <Redirect to={"/signin"} />;
  }
  return (
    <AdminLayoutContext.Provider
      value={{
        lang: result.data.language,
        profile: result.data,
      }}
    >
      {props.children}
    </AdminLayoutContext.Provider>
  );
};
