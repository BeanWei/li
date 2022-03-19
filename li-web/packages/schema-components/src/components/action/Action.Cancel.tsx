import { Button, ButtonProps } from "@arco-design/web-react";
import { useField } from "@formily/react";
import { observer } from "@formily/reactive-react";
import { useContext } from "react";
import { ActionContext } from "./context";

export const ActionCancel: React.FC<ButtonProps> = observer((props) => {
  const field = useField();
  const ctx = useContext(ActionContext);
  return (
    <Button
      {...props}
      onClick={() => {
        ctx.setVisible?.(false);
      }}
    >
      {field.title}
    </Button>
  );
});

export default ActionCancel;
