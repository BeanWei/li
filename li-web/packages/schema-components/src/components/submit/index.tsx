import React from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import { IFormFeedback } from "@formily/core";
import { useParentForm, observer, useFieldSchema } from "@formily/react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { request } from "pro-utils";
import { useTranslation } from "react-i18next";

export interface ISubmitProps extends ButtonProps {
  onClick?: (e: Event) => any;
  forSubmit?: string;
  forSubmitSuccess?: (payload: any) => void;
  forSubmitSuccessTo?: string;
  forSubmitFailed?: (feedbacks: IFormFeedback[]) => void;
}

export const Submit: React.FC<ISubmitProps> = observer(
  ({
    forSubmit,
    forSubmitFailed,
    forSubmitSuccess,
    forSubmitSuccessTo,
    ...props
  }: ISubmitProps) => {
    const form = useParentForm();
    const fieldSchema = useFieldSchema();
    const { t } = useTranslation();
    let navigate: NavigateFunction | ((uri: string) => void);
    try {
      navigate = useNavigate();
    } catch {
      navigate = (uri: string) => {
        // @ts-ignore
        window.location = uri;
      };
    }

    return (
      <Button
        htmlType={forSubmit ? "button" : "submit"}
        type="primary"
        {...props}
        loading={props.loading !== undefined ? props.loading : form.submitting}
        onClick={(e) => {
          if (props.onClick) {
            if (props.onClick(e) === false) return;
          }
          form
            .submit((values) => {
              if (forSubmit) {
                return request(forSubmit, values);
              }
            })
            .then(
              forSubmitSuccessTo
                ? () => {
                    navigate(forSubmitSuccessTo);
                  }
                : forSubmitSuccess
            )
            .catch(forSubmitFailed);
        }}
      >
        {t(fieldSchema["x-content"] || fieldSchema.title)}
      </Button>
    );
  },
  {
    forwardRef: true,
  }
);

export default Submit;
