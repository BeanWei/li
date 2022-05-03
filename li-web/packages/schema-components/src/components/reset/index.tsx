import React from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import { IFieldResetOptions, IFormFeedback } from "@formily/core";
import { useFieldSchema, useParentForm } from "@formily/react";
import { useTranslation } from "react-i18next";

export interface IResetProps extends IFieldResetOptions, ButtonProps {
  onClick?: (e: Event) => any;
  onResetValidateSuccess?: (payload: any) => void;
  onResetValidateFailed?: (feedbacks: IFormFeedback[]) => void;
}

export const Reset: React.FC<IResetProps> = ({
  forceClear,
  validate,
  onResetValidateSuccess,
  onResetValidateFailed,
  ...props
}) => {
  const form = useParentForm();
  const fieldSchema = useFieldSchema();
  const { t } = useTranslation();
  return (
    <Button
      {...props}
      onClick={(e) => {
        if (props.onClick) {
          if (props.onClick(e) === false) return;
        }
        form
          .reset("*", {
            forceClear,
            validate,
          })
          .then(onResetValidateSuccess)
          .catch(onResetValidateFailed);
      }}
    >
      {t(fieldSchema?.["x-content"] || fieldSchema?.title)}
    </Button>
  );
};

export default Reset;
