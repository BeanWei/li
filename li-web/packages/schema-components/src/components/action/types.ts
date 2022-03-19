import {
  ButtonProps,
  DrawerProps,
  ModalProps,
  PopoverProps,
} from "@arco-design/web-react";
import { ConfirmProps } from "@arco-design/web-react/es/Modal/confirm";

export type ActionProps = ButtonProps & {
  component?: any;
  confirm?: ConfirmProps;
  useAction?: () => {
    run(): Promise<void>;
  };
  [key: string]: any;
};

export type ComposedAction = React.FC<ActionProps> & {
  Page?: any;
  Container?: any;
  Drawer?: ComposedActionDrawer;
  Modal?: ComposedActionModal;
  Popover?: ComposedActionPopover;
  Link?: any;
  Cancel?: React.FC<ButtonProps>;
  [key: string]: any;
};

export type ComposedActionDrawer = React.FC<
  DrawerProps & { footerNodeName?: string }
> & {
  Footer?: React.FC;
};

export type ComposedActionModal = React.FC<
  ModalProps & { footerNodeName?: string }
> & {
  Footer?: React.FC;
};

export type ComposedActionPopover = React.FC<
  PopoverProps & { footerNodeName?: string }
> & {
  Footer?: React.FC;
};
