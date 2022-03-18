import React from "react";
import { Button, ButtonProps } from "@arco-design/web-react";
import { IFormFeedback } from "@formily/core";
import { useParentForm, observer, useFieldSchema } from "@formily/react";
import { useHistory } from "react-router-dom";
import { request } from "pro-utils";

export interface ISubmitProps extends ButtonProps {
  onClick?: (e: Event) => any;
  onSubmit?: (values: any) => any;
  onSubmitSuccess?: (payload: any) => void;
  onSubmitSuccessTo?: string;
  onSubmitFailed?: (feedbacks: IFormFeedback[]) => void;
}

export const Submit: React.FC<ISubmitProps> = observer(
  ({
    onSubmit,
    onSubmitFailed,
    onSubmitSuccess,
    onSubmitSuccessTo,
    ...props
  }: ISubmitProps) => {
    const form = useParentForm();
    const schema = useFieldSchema();
    const history = useHistory();

    return (
      <Button
        htmlType={onSubmit ? "button" : "submit"}
        type="primary"
        {...props}
        loading={props.loading !== undefined ? props.loading : form.submitting}
        onClick={(e) => {
          if (props.onClick) {
            if (props.onClick(e) === false) return;
          }
          if (onSubmit) {
            form.submit(onSubmit).then(onSubmitSuccess).catch(onSubmitFailed);
          } else if (schema["x-operation"]) {
            form
              .submit((values) => request(schema["x-operation"], values))
              .then(
                onSubmitSuccessTo
                  ? () => {
                      history.push(onSubmitSuccessTo);
                    }
                  : onSubmitSuccess
              )
              .catch(onSubmitFailed);
          }
        }}
      >
        {props.children}
      </Button>
    );
  },
  {
    forwardRef: true,
  }
);

export default Submit;
