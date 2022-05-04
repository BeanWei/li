import { Avatar as ArcoAvatar, AvatarProps } from "@arco-design/web-react";
import React from "react";

export const Avatar: React.FC<
  React.PropsWithChildren<{
    shape?: AvatarProps["shape"];
    size?: AvatarProps["size"];
    alt?: string;
    src?: string;
  }>
> = (props) => {
  if (props.src) {
    return (
      <ArcoAvatar shape={props.shape || "circle"} size={props.size} {...props}>
        <img alt={props.alt} src={props.src} />
      </ArcoAvatar>
    );
  }
  return <ArcoAvatar {...props}>{props.children}</ArcoAvatar>;
};

export default Avatar;
