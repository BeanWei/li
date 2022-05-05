import React, { Fragment } from "react";
import { createPortal } from "react-dom";
import { createRoot } from "react-dom/client";
import { observable } from "@formily/reactive";
import { Observer } from "@formily/react";
export interface IPortalProps {
  id?: string | symbol;
}

const PortalMap = observable(new Map<string | symbol, React.ReactNode>());

export const createPortalProvider = (id: string | symbol) => {
  const Portal = (props: React.PropsWithChildren<IPortalProps>) => {
    if (props.id && !PortalMap.has(props.id)) {
      PortalMap.set(props.id, null);
    }

    return (
      <Fragment>
        {props.children}
        <Observer>
          {
            // @ts-ignore
            () => {
              if (!props.id) return null;
              const portal = PortalMap.get(props.id);
              if (portal) return createPortal(portal, document.body);
              return null;
            }
          }
        </Observer>
      </Fragment>
    );
  };
  Portal.defaultProps = {
    id,
  };
  return Portal;
};

export function createPortalRoot<T extends React.ReactNode>(
  host: HTMLElement,
  id: string
) {
  const root = createRoot(host);

  function render(renderer?: () => T) {
    if (PortalMap.has(id)) {
      PortalMap.set(id, renderer?.());
    } else if (host) {
      root.render(<Fragment>{renderer?.()}</Fragment>);
    }
  }

  async function unmount() {
    if (PortalMap.has(id)) {
      PortalMap.set(id, null);
    }
    if (host) {
      host.parentNode?.removeChild(host);
      // Delay to unmount to avoid React 18 sync warning
      Promise.resolve().then(() => {
        root.unmount();
      });
    }
  }

  return {
    render,
    unmount,
  };
}
