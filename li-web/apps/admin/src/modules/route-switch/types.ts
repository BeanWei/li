export interface RouteProps {
  path: string;
  index?: boolean;
  redirect?: string;
  component?: any;
  // routes?: RouteProps[];
  [key: string]: any;
}

export interface RouteSwitchProviderProps {
  components?: any;
  children?: any;
}

export interface RouteSwitchProps {
  routes?: RouteProps[];
  components?: any;
}
