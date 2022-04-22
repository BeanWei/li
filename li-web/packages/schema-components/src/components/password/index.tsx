import React from "react";
import { connect, mapReadPretty } from "@formily/react";
import { Input, InputProps } from "@arco-design/web-react";
import { useTranslation } from "react-i18next";
import { PasswordStrength } from "./PasswordStrength";

export interface IPasswordProps extends InputProps {
  checkStrength: boolean;
  className?: string;
}

export const Password = connect(
  (props: IPasswordProps) => {
    const { value, className, checkStrength, ...rest } = props;
    const blockStyle: React.CSSProperties = {
      position: "absolute",
      zIndex: 1,
      height: 8,
      top: 0,
      background: "#fff",
      width: 1,
      transform: "translate(-50%, 0)",
    };
    const { t } = useTranslation();
    return (
      <span className={className || ""}>
        <Input.Password
          {...rest}
          placeholder={rest.placeholder ? t(rest.placeholder) : undefined}
          value={value}
        />
        {checkStrength && (
          <PasswordStrength value={String(value)}>
            {(score) => {
              return (
                <div
                  style={{
                    background: "#e0e0e0",
                    marginBottom: 3,
                    position: "relative",
                  }}
                >
                  <div style={{ ...blockStyle, left: "20%" }} />
                  <div style={{ ...blockStyle, left: "40%" }} />
                  <div style={{ ...blockStyle, left: "60%" }} />
                  <div style={{ ...blockStyle, left: "80%" }} />
                  <div
                    style={{
                      position: "relative",
                      backgroundImage:
                        "-webkit-linear-gradient(left, #ff5500, #ff9300)",
                      transition: "all 0.35s ease-in-out",
                      height: 8,
                      width: "100%",
                      marginTop: 5,
                      clipPath: `polygon(0 0,${score}% 0,${score}% 100%,0 100%)`,
                    }}
                  />
                </div>
              );
            }}
          </PasswordStrength>
        )}
      </span>
    );
  },
  mapReadPretty((props) => {
    if (!props.value) {
      return <div>-</div>;
    }
    return <div>********</div>;
  })
);

export default Password;
