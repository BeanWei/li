import { createForm } from "@formily/core";
import { FormProvider, Schema } from "@formily/react";
import { uid } from "@formily/shared";
import React, { useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import { SchemaComponentContext } from "../context";
import { ISchemaComponentProvider } from "../types";
import { SchemaComponentOptions } from "./SchemaComponentOptions";

const randomString = (prefix: string = "") => {
  return `${prefix}${uid()}`;
};

Schema.silent(true);

const Registry = {
  silent: true,
  compile(expression: string, scope = {}) {
    if (Registry.silent) {
      try {
        return new Function("$root", `with($root) { return (${expression}); }`)(
          scope
        );
      } catch {
        return `{{${expression}}}`;
      }
    } else {
      return new Function("$root", `with($root) { return (${expression}); }`)(
        scope
      );
    }
  },
};

Schema.registerCompiler(Registry.compile);

export const SchemaComponentProvider: React.FC<ISchemaComponentProvider> = (
  props
) => {
  const { components, children } = props;
  const [, setUid] = useState(uid());
  const [formId, setFormId] = useState(uid());
  const form = props.form || useMemo(() => createForm(), [formId]);
  const { t } = useTranslation();
  const scope = { ...props.scope, t, randomString };

  return (
    <SchemaComponentContext.Provider
      value={{
        scope,
        components,
        reset: () => setFormId(uid()),
        refresh: () => setUid(uid()),
      }}
    >
      <FormProvider form={form}>
        <SchemaComponentOptions inherit scope={scope} components={components}>
          {children}
        </SchemaComponentOptions>
      </FormProvider>
    </SchemaComponentContext.Provider>
  );
};
