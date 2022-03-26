import { Suspense } from "react";
import { Loading } from "../components";

const Blank = ({ children }: any) => children || null;

export const compose = (...components: any[]) => {
  const Root = [...components, Blank].reduce((parent, child) => {
    const [Parent, parentProps] = Array.isArray(parent) ? parent : [parent];
    const [Child, childProps] = Array.isArray(child) ? child : [child];
    return ({ children }: any) => (
      <Parent {...parentProps}>
        <Child {...childProps}>{children}</Child>
      </Parent>
    );
  });
  return (LastChild?: any) => (props?: any) => {
    return (
      <Suspense fallback={<Loading />}>
        <Root>{LastChild && <LastChild {...props} />}</Root>
      </Suspense>
    );
  };
};
