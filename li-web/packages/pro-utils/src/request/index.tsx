import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import { useRequest as useAhookRequest } from "ahooks";
import { Message, Notification } from "@arco-design/web-react";
import { Options, Result } from "ahooks/lib/useRequest/src/types";

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
  code?: string;
  message?: string;
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
    error: any,
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
    if (error.name === "Error") {
      const adaptor: IAdaptor =
        errorConfig?.adaptor || ((errorData) => errorData);
      errorInfo = adaptor(error.response.data, error.response);
      error.info = errorInfo;
      error.data = error.response.data;
    }
    errorInfo = error.info;
    if (errorInfo) {
      const { message = "(ノ﹏ヽ)", code } = errorInfo;
      switch (errorInfo.showType) {
        case ErrorShowType.SILENT:
          // do nothong
          break;
        case ErrorShowType.WARN_MESSAGE:
          Message.warning(message);
          break;
        case ErrorShowType.ERROR_MESSAGE:
          Message.error(message);
          break;
        case ErrorShowType.NOTIFICATION:
          Notification.error({ content: message, title: code });
          break;
        case ErrorShowType.REDIRECT:
          // TODO: redirect
          break;
        default:
          Message.error(message);
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

let requestInstance: AxiosInstance;

const defaultConfig: RequestConfig = {
  baseURL: window.location.origin + "/api/liql",
  method: "POST",
};

const getRequestInstance = (): AxiosInstance => {
  if (requestInstance) return requestInstance;
  requestInstance = axios.create(defaultConfig);
  return requestInstance;
};

const request: IRequest = (operation, variables = {}, opts = {}) => {
  const requestInstance = getRequestInstance();
  console.log(`${operation} => `, variables);
  return new Promise((resolve, reject) => {
    requestInstance
      .request({ ...opts, data: { operation, variables } })
      .then((res) => {
        if (res.data.code !== 0) {
          try {
            const handler =
              defaultConfig.errorConfig?.errorHandler || defaultErrorHandler;
            handler(
              {
                response: res,
                info: res.data,
              },
              opts,
              defaultConfig
            );
          } catch (e) {
            reject(e);
          }
        } else {
          const formatResultAdaptor =
            defaultConfig?.formatResultAdaptor || ((res) => res.data?.data);
          resolve(formatResultAdaptor(res));
        }
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

const useRequest = (
  operation: string,
  variables?: Record<string, any>,
  options?: Options<any, any>
): Result<any, any> => {
  return useAhookRequest((params) => {
    return request(operation, params || variables);
  }, options);
};

export { useRequest, request };

export type { Result, AxiosRequestConfig, AxiosInstance, AxiosResponse };
