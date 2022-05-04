import React, { createContext, useContext } from "react";
import { Form as FormType, ObjectField, IFormFeedback } from "@formily/core";
import {
  useParentForm,
  FormProvider,
  ExpressionScope,
  JSXComponent,
} from "@formily/react";
import { FormLayout, IFormLayoutProps } from "../form-layout";
import { isValid } from "@formily/shared";
export interface FormProps extends IFormLayoutProps {
  form?: FormType;
  component?: JSXComponent;
  onAutoSubmit?: (values: any) => any;
  onAutoSubmitFailed?: (feedbacks: IFormFeedback[]) => void;
  previewTextPlaceholder?: React.ReactNode;
}

const PlaceholderContext = createContext<React.ReactNode>("N/A");

const Placeholder = PlaceholderContext.Provider;

const usePlaceholder = (value?: any) => {
  const placeholder = useContext(PlaceholderContext) || "N/A";
  return isValid(value) && value !== "" ? value : placeholder;
};

export const Form: React.FC<React.PropsWithChildren<FormProps>> = ({
  form,
  component,
  onAutoSubmit,
  onAutoSubmitFailed,
  previewTextPlaceholder,
  ...props
}) => {
  const top = useParentForm();
  const renderContent = (form: FormType | ObjectField) => (
    <ExpressionScope value={{ $$form: form }}>
      <Placeholder value={previewTextPlaceholder}>
        <FormLayout {...props}>
          {React.createElement(
            component ?? "form",
            {
              onSubmit(e: React.FormEvent) {
                e?.stopPropagation?.();
                e?.preventDefault?.();
                form.submit(onAutoSubmit).catch(onAutoSubmitFailed);
              },
            },
            props.children
          )}
        </FormLayout>
      </Placeholder>
    </ExpressionScope>
  );
  if (form)
    return <FormProvider form={form}>{renderContent(form)}</FormProvider>;
  if (!top) throw new Error("must pass form instance by createForm");
  return renderContent(top);
};

Form.defaultProps = {
  component: "form",
};

export default Form;
