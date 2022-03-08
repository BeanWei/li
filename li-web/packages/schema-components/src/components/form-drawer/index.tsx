import React, { Fragment, useLayoutEffect, useRef, useState } from "react";
import { createPortal } from "react-dom";
import {
  createForm,
  IFormProps,
  Form,
  onFormSubmitSuccess,
} from "@formily/core";
import { toJS } from "@formily/reactive";
import { FormProvider, observer } from "@formily/react";
import {
  isNum,
  isStr,
  isBool,
  isFn,
  applyMiddleware,
  IMiddleware,
} from "@formily/shared";
import { Drawer, DrawerProps } from "@arco-design/web-react";
import {
  usePrefixCls,
  createPortalProvider,
  createPortalRoot,
  loading,
} from "../__builtins__";

type FormDrawerRenderer =
  | React.ReactElement
  | ((form: Form) => React.ReactElement);

type DrawerTitle = string | number | React.ReactElement;

type EventType =
  | React.KeyboardEvent<HTMLDivElement>
  | React.MouseEvent<HTMLDivElement | HTMLButtonElement>;

const isDrawerTitle = (props: any): props is DrawerTitle => {
  return (
    isNum(props) || isStr(props) || isBool(props) || React.isValidElement(props)
  );
};

const getDrawerProps = (props: any): IDrawerProps => {
  if (isDrawerTitle(props)) {
    return {
      title: props,
    };
  } else {
    return props;
  }
};

export interface IFormDrawer {
  forOpen(middleware: IMiddleware<IFormProps>): IFormDrawer;
  open(props?: IFormProps): Promise<any>;
  close(): void;
}

export interface IDrawerProps extends DrawerProps {
  onCancel?: () => void | boolean;
  loadingText?: React.ReactNode;
}

export function FormDrawer(
  title: IDrawerProps,
  id: string,
  renderer: FormDrawerRenderer
): IFormDrawer;
export function FormDrawer(
  title: IDrawerProps,
  id: FormDrawerRenderer
): IFormDrawer;
export function FormDrawer(
  title: DrawerTitle,
  id: string,
  renderer: FormDrawerRenderer
): IFormDrawer;
export function FormDrawer(
  title: DrawerTitle,
  id: FormDrawerRenderer
): IFormDrawer;
export function FormDrawer(title: any, id: any, renderer?: any): IFormDrawer {
  if (isFn(id) || React.isValidElement(id)) {
    renderer = id;
    id = "form-drawer";
  }
  const env: any = {
    host: document.createElement("div"),
    openMiddlewares: [],
    form: null,
    promise: null,
  };
  const root = createPortalRoot(env.host, id);
  const props = getDrawerProps(title);
  const drawer: IDrawerProps = {
    width: "40%",
    ...props,
    onCancel: () => {
      if (props?.onCancel?.() !== false) {
        formDrawer.close();
      }
    },
    afterClose: () => {
      props?.afterClose?.();
      root.unmount();
    },
  };
  const DrawerContent = observer(() => {
    return (
      <Fragment>{isFn(renderer) ? renderer(env.form) : renderer}</Fragment>
    );
  });
  const renderDrawer = (visible = true) => {
    return (
      <Drawer autoFocus={false} {...drawer} visible={visible}>
        <FormProvider form={env.form}>
          <DrawerContent />
        </FormProvider>
      </Drawer>
    );
  };

  document.body.appendChild(env.host);
  const formDrawer = {
    forOpen: (middleware: IMiddleware<IFormProps>) => {
      if (isFn(middleware)) {
        env.openMiddlewares.push(middleware);
      }
      return formDrawer;
    },
    open: (props: IFormProps) => {
      if (env.promise) return env.promise;
      env.promise = new Promise(async (resolve, reject) => {
        try {
          props = await loading(drawer.loadingText, () =>
            applyMiddleware(props, env.openMiddlewares)
          );
          env.form =
            env.form ||
            createForm({
              ...props,
              effects(form) {
                onFormSubmitSuccess(() => {
                  resolve(toJS(form.values));
                  formDrawer.close();
                });
                props?.effects?.(form);
              },
            });
        } catch (e) {
          reject(e);
        }
        root.render(() => renderDrawer(false));
        setTimeout(() => {
          root.render(() => renderDrawer(true));
        }, 16);
      });
      return env.promise;
    },
    close: () => {
      if (!env.host) return;
      root.render(() => renderDrawer(false));
    },
  };
  return formDrawer;
}

const DrawerFooter: React.FC = (props) => {
  const ref = useRef<HTMLDivElement | null>(null);
  const [footer, setFooter] = useState<HTMLDivElement | null>(null);
  const footerRef = useRef<HTMLDivElement | null>(null);
  const prefixCls = usePrefixCls("drawer");
  useLayoutEffect(() => {
    const content = ref.current?.closest(`.${prefixCls}-wrapper-body`);
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

FormDrawer.Footer = DrawerFooter;

FormDrawer.Portal = createPortalProvider("form-drawer");

export default FormDrawer;
