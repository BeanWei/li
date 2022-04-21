import { Navigate, Route, Routes } from "react-router-dom";
import { RouteContext } from "./context";
import { useRouteComponent } from "./hooks";
import { RouteSwitchProps } from "./types";

export function RouteSwitch(props: RouteSwitchProps) {
  const { routes = [] } = props;
  if (!routes.length) {
    return null;
  }
  return (
    <Routes>
      {routes.map((route, index) => {
        return (
          <Route
            key={index}
            path={route.path}
            index={route.index}
            element={
              route.redirect ? (
                <Navigate to={route.redirect} />
              ) : (
                <RouteContext.Provider value={route}>
                  <ComponentRenderer component={route.component} />
                </RouteContext.Provider>
              )
            }
          />
        );
      })}
    </Routes>
  );
}

function ComponentRenderer(props: { component: string }) {
  const Component = useRouteComponent(props.component);
  return <Component />;
}
