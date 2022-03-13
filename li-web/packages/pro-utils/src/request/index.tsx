import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import useAhookRequest, { UseRequestProvider } from "@ahooksjs/use-request";
import { Message, Notification } from "@arco-design/web-react";

import {
  BaseOptions,
  BasePaginatedOptions,
  BaseResult,
  CombineService,
  LoadMoreFormatReturn,
  LoadMoreOptions,
  LoadMoreOptionsWithFormat,
  LoadMoreParams,
  LoadMoreResult,
  OptionsWithFormat,
  PaginatedFormatReturn,
  PaginatedOptionsWithFormat,
  PaginatedParams,
  PaginatedResult,
} from "@ahooksjs/use-request/es/types";

type ResultWithData<T = any> = { data?: T; [key: string]: any };

function useRequest<
  R = any,
  P extends any[] = any,
  U = any,
  UU extends U = any
>(
  service: CombineService<R, P>,
  options: OptionsWithFormat<R, P, U, UU>
): BaseResult<U, P>;
function useRequest<R extends ResultWithData = any, P extends any[] = any>(
  service: CombineService<R, P>,
  options?: BaseOptions<R["data"], P>
): BaseResult<R["data"], P>;
function useRequest<R extends LoadMoreFormatReturn = any, RR = any>(
  service: CombineService<RR, LoadMoreParams<R>>,
  options: LoadMoreOptionsWithFormat<R, RR>
): LoadMoreResult<R>;
function useRequest<
  R extends ResultWithData<LoadMoreFormatReturn | any> = any,
  RR extends R = any
>(
  service: CombineService<R, LoadMoreParams<R["data"]>>,
  options: LoadMoreOptions<RR["data"]>
): LoadMoreResult<R["data"]>;

function useRequest<R = any, Item = any, U extends Item = any>(
  service: CombineService<R, PaginatedParams>,
  options: PaginatedOptionsWithFormat<R, Item, U>
): PaginatedResult<Item>;
function useRequest<Item = any, U extends Item = any>(
  service: CombineService<
    ResultWithData<PaginatedFormatReturn<Item>>,
    PaginatedParams
  >,
  options: BasePaginatedOptions<U>
): PaginatedResult<Item>;
function useRequest(service: any, options: any = {}) {
  return useAhookRequest(service, {
    formatResult: (result) => result?.data,
    requestMethod: (requestOptions: any) => {
      console.log("[Request/RequestOptions]: ", requestOptions);
      if (typeof requestOptions === "string") {
        return request(requestOptions, {}, {});
      }
      if (typeof requestOptions === "object") {
        const { operation, variables, ...rest } = requestOptions;
        return request(operation, variables, rest);
      }
      throw new Error("request options error");
    },
    ...options,
  });
}

export interface RequestConfig extends AxiosRequestConfig {
  errorConfig?: {
    errorPage?: string;
    adaptor?: IAdaptor; // adaptor 用以用户将不满足接口的后端数据修改成 errorInfo
    errorHandler?: IErrorHandler;
    defaultNoneResponseErrorMessage?: string;
    defaultRequestErrorMessage?: string;
  };
  formatResultAdaptor?: IFormatResultAdaptor;
}

export enum ErrorShowType {
  SILENT = 0,
  WARN_MESSAGE = 1,
  ERROR_MESSAGE = 2,
  NOTIFICATION = 3,
  REDIRECT = 9,
}

export interface IErrorInfo {
  success: boolean;
  data?: any;
  errorCode?: string;
  errorMessage?: string;
  showType?: ErrorShowType;
  traceId?: string;
  host?: string;
  [key: string]: any;
}
// resData 其实就是 response.data, response 则是 axios 的响应对象
interface IAdaptor {
  (resData: any, response: AxiosResponse): IErrorInfo;
}

export interface RequestError extends Error {
  request?: any;
  response?: any;
  data?: any;
  info?: IErrorInfo;
}

interface IRequest {
  (
    operation: string,
    variables?: Record<string, any>,
    opts?: AxiosRequestConfig & { skipErrorHandler?: boolean }
  ): Promise<AxiosResponse<any, any>>;
}

interface IErrorHandler {
  (
    error: RequestError,
    opts: AxiosRequestConfig & { skipErrorHandler?: boolean },
    config: RequestConfig
  ): void;
}

interface IFormatResultAdaptor {
  (res: AxiosResponse): any;
}

const defaultErrorHandler: IErrorHandler = (error, opts, config) => {
  if (opts?.skipErrorHandler) throw error;
  const { errorConfig } = config;
  if (error.response) {
    // 请求成功发出且服务器也响应了状态码，但状态代码超出了 2xx 的范围 或者 成功响应，success字段为false 由我们抛出的错误
    let errorInfo: IErrorInfo | undefined;
    // 不是我们的错误
    if (error.name === "ResponseError") {
      const adaptor: IAdaptor =
        errorConfig?.adaptor || ((errorData) => errorData);
      errorInfo = adaptor(error.response.data, error.response);
      error.info = errorInfo;
      error.data = error.response.data;
    }
    errorInfo = error.info;
    if (errorInfo) {
      const { errorMessage = "(ノ﹏ヽ)", errorCode } = errorInfo;
      switch (errorInfo.showType) {
        case ErrorShowType.SILENT:
          // do nothong
          break;
        case ErrorShowType.WARN_MESSAGE:
          Message.warning(errorMessage);
          break;
        case ErrorShowType.ERROR_MESSAGE:
          Message.error(errorMessage);
          break;
        case ErrorShowType.NOTIFICATION:
          Notification.error({ content: errorMessage, title: errorCode });
          break;
        case ErrorShowType.REDIRECT:
          // TODO: redirect
          break;
        default:
          Message.error(errorMessage);
      }
    }
  } else if (error.request) {
    // 请求已经成功发起，但没有收到响应
    // \`error.request\` 在浏览器中是 XMLHttpRequest 的实例，
    // 而在node.js中是 http.ClientRequest 的实例
    Message.error(
      errorConfig?.defaultNoneResponseErrorMessage ||
        "None response! Please retry."
    );
  } else {
    // 发送请求时出了点问题
    Message.error(
      errorConfig?.defaultRequestErrorMessage || "Request error, please retry."
    );
  }
  throw error;
};

const defaultConfig: RequestConfig = {
  baseURL: window.location.origin + "/api/liql",
  method: "POST",
};

let requestInstance: AxiosInstance;
const getRequestInstance = (): AxiosInstance => {
  if (requestInstance) return requestInstance;
  requestInstance = axios.create(defaultConfig);

  // 当响应的数据 success 是 false 的时候，抛出 error 以供 errorHandler 处理。
  requestInstance.interceptors.response.use((response) => {
    const { data } = response;
    const adaptor =
      defaultConfig?.errorConfig?.adaptor || ((resData) => resData);
    const errorInfo = adaptor(data, response);
    if (errorInfo.success === false) {
      const error: RequestError = new Error(errorInfo.errorMessage);
      error.name = "BizError";
      error.data = data;
      error.info = errorInfo;
      error.response = response;
      throw error;
    }
    return response;
  });
  return requestInstance;
};

const request: IRequest = (operation, variables = {}, opts = {}) => {
  const requestInstance = getRequestInstance();
  return new Promise((resolve, reject) => {
    requestInstance
      .request({ ...opts, data: { operation, variables } })
      .then((res) => {
        const formatResultAdaptor =
          defaultConfig?.formatResultAdaptor || ((res) => res.data);
        resolve(formatResultAdaptor(res));
      })
      .catch((error) => {
        try {
          const handler =
            defaultConfig.errorConfig?.errorHandler || defaultErrorHandler;
          handler(error, opts, defaultConfig);
        } catch (e) {
          reject(e);
        }
      });
  });
};

export { useRequest, UseRequestProvider, request };

export type { AxiosInstance, AxiosRequestConfig, AxiosResponse };
