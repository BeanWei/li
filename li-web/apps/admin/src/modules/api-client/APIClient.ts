import { observable } from "@formily/reactive";
import { Result } from "ahooks/lib/useRequest/src/types";
import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import Cookies from "js-cookie";
import qs from "qs";

export interface ActionParams {
  filterByTk?: any;
  [key: string]: any;
}

type ResourceActionOptions<P = any> = {
  resource?: string;
  resourceOf?: any;
  action?: string;
  params?: P;
};

export interface IResource {
  list?: (params?: ActionParams) => Promise<any>;
  get?: (params?: ActionParams) => Promise<any>;
  create?: (params?: ActionParams) => Promise<any>;
  update?: (params?: ActionParams) => Promise<any>;
  destroy?: (params?: ActionParams) => Promise<any>;
  [key: string]: (params?: ActionParams) => Promise<any>;
}

export class APIClient {
  axios: AxiosInstance;

  services: Record<string, Result<any, any>>;

  tokenKey = "nocobaseToken";

  constructor(instance?: AxiosInstance | AxiosRequestConfig) {
    this.services = observable({});
    if (typeof instance === "function") {
      this.axios = instance;
    } else {
      this.axios = axios.create(instance);
    }
    this.qsMiddleware();
    this.authMiddleware();
  }

  qsMiddleware() {
    this.axios.interceptors.request.use((config) => {
      config.paramsSerializer = (params) => {
        return qs.stringify(params, {
          strictNullHandling: true,
          arrayFormat: "brackets",
        });
      };
      return config;
    });
  }

  // TODO
  authMiddleware() {
    this.axios.interceptors.request.use((config) => {
      const token = localStorage.getItem(this.tokenKey);
      if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
      }
      const currentRoleName = Cookies.get("currentRoleName");
      if (currentRoleName) {
        config.headers["X-Role"] = currentRoleName;
      }
      return config;
    });
  }

  // TODO
  setBearerToken(token: any) {
    localStorage.setItem(this.tokenKey, token || "");
    Cookies.remove("currentRoleName");
  }

  service(uid: string): Result<any, any> {
    return this.services[uid];
  }

  request<T = any, R = AxiosResponse<T>, D = any>(
    config: AxiosRequestConfig<D> | ResourceActionOptions
  ): Promise<R> {
    const { resource, resourceOf, action, params } = config as any;
    if (resource) {
      return this.resource(resource, resourceOf)[action](params);
    }
    return this.axios.request<T, R, D>(config);
  }

  resource(name: string, of?: any): IResource {
    const target = {};
    const handler = {
      get: (_: any, actionName: string) => {
        let url = name.split(".").join(`/${of || "_"}/`);
        url += `:${actionName}`;
        const config: AxiosRequestConfig = { url };
        if (["get", "list"].includes(actionName)) {
          config["method"] = "get";
        } else {
          config["method"] = "post";
        }
        return async (params?: ActionParams) => {
          const { values, filter, ...others } = params || {};
          config["params"] = others;
          if (filter) {
            if (typeof filter === "string") {
              config["params"]["filter"] = filter;
            } else {
              config["params"]["filter"] = JSON.stringify(filter);
            }
          }
          if (config.method !== "get") {
            config["data"] = values || {};
          }
          return await this.request(config);
        };
      },
    };
    return new Proxy(target, handler);
  }
}
