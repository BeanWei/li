import { Message } from "@arco-design/web-react";

export const loading = async (
  title: React.ReactNode = "Loading...",
  processor: () => Promise<any>
) => {
  let hide: any = null;
  let loading = setTimeout(() => {
    hide = Message.loading({
      content: title,
    });
  }, 100);
  try {
    return await processor();
  } finally {
    hide?.();
    clearTimeout(loading);
  }
};
