import {
  ButtonProps,
  DrawerProps,
  ModalProps,
  PopoverProps,
} from "@arco-design/web-react";

export type ActionProps = ButtonProps & {
  component?: any;
  useAction?: () => {
    run(): Promise<void>;
  };
};

export type ComposedAction = React.FC<ActionProps> & {
  Page?: any;
  Container?: any;
  Drawer?: ComposedActionDrawer;
  Modal?: ComposedActionModal;
  Popover?: ComposedActionPopover;
  Link?: any;
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
