import React, { Fragment, useRef, useLayoutEffect, useState } from "react";
import { createPortal } from "react-dom";
import { createForm, IFormProps, Form } from "@formily/core";
import { toJS } from "@formily/reactive";
import { FormProvider, Observer, observer } from "@formily/react";
import {
  isNum,
  isStr,
  isBool,
  isFn,
  applyMiddleware,
  IMiddleware,
} from "@formily/shared";
import { Modal, ModalProps } from "@arco-design/web-react";
import {
  usePrefixCls,
  loading,
  createPortalProvider,
  createPortalRoot,
} from "../__builtins__";

type FormModalRenderer =
  | React.ReactElement
  | ((form: Form) => React.ReactElement);

type ModalTitle = string | number | React.ReactElement;

const isModalTitle = (props: any): props is ModalTitle => {
  return (
    isNum(props) || isStr(props) || isBool(props) || React.isValidElement(props)
  );
};

const getModelProps = (props: any): IModalProps => {
  if (isModalTitle(props)) {
    return {
      title: props,
    };
  } else {
    return props;
  }
};

export interface IFormModal {
  forOpen(middleware: IMiddleware<IFormProps>): IFormModal;
  forConfirm(middleware: IMiddleware<Form>): IFormModal;
  forCancel(middleware: IMiddleware<Form>): IFormModal;
  open(props?: IFormProps): Promise<any>;
  close(): void;
}

// @ts-ignore
export interface IModalProps extends ModalProps {
  onOk?: (e?: MouseEvent) => void | boolean;
  onCancel?: () => void | boolean;
  loadingText?: React.ReactNode;
}

export function FormModal(
  title: IModalProps,
  id: string,
  renderer: FormModalRenderer
): IFormModal;
export function FormModal(
  title: IModalProps,
  renderer: FormModalRenderer
): IFormModal;
export function FormModal(
  title: ModalTitle,
  id: string,
  renderer: FormModalRenderer
): IFormModal;
export function FormModal(
  title: ModalTitle,
  renderer: FormModalRenderer
): IFormModal;
export function FormModal(title: any, id: any, renderer?: any): IFormModal {
  if (isFn(id) || React.isValidElement(id)) {
    renderer = id;
    id = "form-modal";
  }
  const env: any = {
    host: document.createElement("div"),
    form: null,
    promise: null,
    openMiddlewares: [],
    confirmMiddlewares: [],
    cancelMiddlewares: [],
  };
  const root = createPortalRoot(env.host, id);
  const props = getModelProps(title);
  const modal = {
    ...props,
    afterClose: () => {
      props?.afterClose?.();
      root.unmount();
    },
  };
  const ModalContent = observer(() => {
    return (
      <Fragment>{isFn(renderer) ? renderer(env.form) : renderer}</Fragment>
    );
  });
  const renderModal = (
    visible = true,
    resolve?: () => any,
    reject?: () => any
  ) => {
    return (
      <Observer>
        {() => (
          <Modal
            {...modal}
            visible={visible}
            confirmLoading={env.form?.submitting}
            onCancel={() => {
              if (modal?.onCancel?.() !== false) {
                reject?.();
              }
            }}
            onOk={async (e) => {
              if (modal?.onOk?.(e) !== false) {
                resolve?.();
              }
            }}
          >
            <FormProvider form={env.form}>
              <ModalContent />
            </FormProvider>
          </Modal>
        )}
      </Observer>
    );
  };

  document.body.appendChild(env.host);
  const formModal = {
    forOpen: (middleware: IMiddleware<IFormProps>) => {
      if (isFn(middleware)) {
        env.openMiddlewares.push(middleware);
      }
      return formModal;
    },
    forConfirm: (middleware: IMiddleware<Form>) => {
      if (isFn(middleware)) {
        env.confirmMiddlewares.push(middleware);
      }
      return formModal;
    },
    forCancel: (middleware: IMiddleware<Form>) => {
      if (isFn(middleware)) {
        env.cancelMiddlewares.push(middleware);
      }
      return formModal;
    },
    open: async (props: IFormProps) => {
      if (env.promise) return env.promise;
      env.promise = new Promise(async (resolve, reject) => {
        try {
          props = await loading(modal.loadingText, () =>
            applyMiddleware(props, env.openMiddlewares)
          );
          env.form = env.form || createForm(props);
        } catch (e) {
          reject(e);
        }
        root.render(() =>
          renderModal(
            true,
            () => {
              env.form
                .submit(async () => {
                  await applyMiddleware(env.form, env.confirmMiddlewares);
                  resolve(toJS(env.form.values));
                  formModal.close();
                })
                .catch(() => {});
            },
            async () => {
              await loading(modal.loadingText, () =>
                applyMiddleware(env.form, env.cancelMiddlewares)
              );
              formModal.close();
            }
          )
        );
      });
      return env.promise;
    },
    close: () => {
      if (!env.host) return;
      root.render(() => renderModal(false));
    },
  };
  return formModal;
}

const ModalFooter: React.FC = (props) => {
  const ref = useRef<HTMLDivElement | null>(null);
  const [footer, setFooter] = useState<HTMLDivElement | null>(null);
  const footerRef = useRef<HTMLDivElement | null>(null);
  const prefixCls = usePrefixCls("modal");
  useLayoutEffect(() => {
    const content = ref.current?.closest(`.${prefixCls}-content`);
    if (content) {
      if (!footerRef.current) {
        footerRef.current = content.querySelector(`.${prefixCls}-footer`);
        if (!footerRef.current) {
          footerRef.current = document.createElement("div");
          footerRef.current.classList.add(`${prefixCls}-footer`);
          content.appendChild(footerRef.current);
        }
      }
      setFooter(footerRef.current);
    }
  });

  footerRef.current = footer;

  return (
    <div ref={ref} style={{ display: "none" }}>
      {footer && createPortal(props.children, footer)}
    </div>
  );
};

FormModal.Footer = ModalFooter;

FormModal.Portal = createPortalProvider("form-modal");

export default FormModal;
