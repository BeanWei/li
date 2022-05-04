import React, { createContext, useContext, useEffect, useState } from "react";
import { Helmet } from "react-helmet";

interface DocumentTitleContextProps {
  title?: any;
  setTitle?: (title?: any) => void;
}

export const DocumentTitleContext = createContext<DocumentTitleContextProps>({
  title: null,
  setTitle() {},
});

export const DocumentTitleProvider: React.FC<
  React.PropsWithChildren<{
    addonBefore?: string;
    addonAfter?: string;
  }>
> = (props) => {
  const { addonBefore, addonAfter } = props;
  const [title, setTitle] = useState("");
  const documentTitle = title
    ? `${addonBefore ? `${addonBefore} | ` : ""}${title || ""}${
        addonAfter ? ` | ${addonAfter}` : ""
      }`
    : "Li Admin";
  return (
    <DocumentTitleContext.Provider
      value={{
        title,
        setTitle,
      }}
    >
      <Helmet>
        <title>{documentTitle}</title>
      </Helmet>
      {props.children}
    </DocumentTitleContext.Provider>
  );
};

export const useDocumentTitle = () => {
  return useContext(DocumentTitleContext);
};

export const useCurrentDocumentTitle = (title: string) => {
  const { setTitle } = useDocumentTitle();
  useEffect(() => {
    setTitle?.(title);
  }, []);
};
