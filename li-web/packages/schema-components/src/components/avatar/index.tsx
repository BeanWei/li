import { Avatar as ArcoAvatar, AvatarProps } from "@arco-design/web-react";

export const Avatar: React.FC<{
  shape?: AvatarProps["shape"];
  size?: AvatarProps["size"];
  alt?: string;
  src?: string;
}> = (props) => {
  if (props.src) {
    return (
      <ArcoAvatar shape={props.shape} size={props.size} {...props}>
        <img alt={props.alt} src={props.src} />
      </ArcoAvatar>
    );
  }
  return <ArcoAvatar {...props}>{props.children}</ArcoAvatar>;
};

export default Avatar;
