import { UploadProps } from "@arco-design/web-react";
import React from "react";

export type UploadAttachmentProps = Omit<UploadProps, "onChange"> & {
  value?: string[];
  onChange?: (value?: (string | undefined)[]) => void;
};

export type UploadAvatarProps = Omit<UploadProps, "onChange"> & {
  value?: string;
  onChange?: (value?: string) => void;
};

export type ComposedUpload = React.FC & {
  Attachment?: React.FC<UploadAttachmentProps>;
  Avatar?: React.FC<UploadAvatarProps>;
};
