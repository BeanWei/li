import { Fragment, useState } from "react";
import {
  Avatar,
  Image,
  Progress,
  Space,
  Upload as ArcoUpload,
} from "@arco-design/web-react";
import { connect, mapReadPretty } from "@formily/react";
import { isValid } from "@formily/shared";
import {
  ComposedUpload,
  UploadAttachmentProps,
  UploadAvatarProps,
} from "./types";
import { UploadItem } from "@arco-design/web-react/es/Upload";
import { IconEdit, IconPlus } from "@arco-design/web-react/icon";
import { getPrefixCls } from "../__builtins__";

const normalizeFileList = (fileList?: string[] | string) => {
  const files = Array.isArray(fileList) ? fileList : fileList ? [fileList] : [];
  if (files && files.length) {
    return files.map((url, index) => {
      return {
        uid: `${index}`,
        url,
      };
    });
  }
  return [];
};

const UploadAttachment: React.FC<UploadAttachmentProps> = connect(
  (props: UploadAttachmentProps) => {
    const { onChange, value, ...rest } = props;
    return (
      <ArcoUpload
        {...rest}
        withCredentials
        defaultFileList={normalizeFileList(value)}
        onChange={(fileList) => {
          onChange?.(fileList.map((file) => file.url));
        }}
      />
    );
  },
  mapReadPretty((props) => {
    const { value } = props;
    if (!isValid(value)) {
      return <div>-</div>;
    }
    return (
      <Space direction="vertical">
        <Image.PreviewGroup infinite>
          <Space>
            {value.map((src: string, index: number) => {
              return <Image key={index} src={src} width={80} />;
            })}
          </Space>
        </Image.PreviewGroup>
      </Space>
    );
  })
);

const UploadAvatar: React.FC<UploadAvatarProps> = connect(
  (props: UploadAvatarProps) => {
    const { onChange, value, ...rest } = props;
    const [file, setFile] = useState<UploadItem>({
      url: value,
      uid: "-1",
    });
    const prefixCls = getPrefixCls();
    const cs = `${prefixCls}-upload-list-item${
      file && file.status === "error" ? " is-error" : ""
    }`;
    return (
      <ArcoUpload
        {...rest}
        withCredentials
        showUploadList={false}
        onChange={(_, currentFile) => {
          if (currentFile.status == "done") {
            // @ts-ignore
            const { data = {} } = currentFile.response;
            onChange?.(data.url);
            setFile({
              ...file,
              url: data.url,
            });
          }
        }}
      >
        <div className={cs} style={{ marginTop: 0 }}>
          {file && file.url ? (
            <div
              className={`${prefixCls}-upload-list-item-picture custom-upload-avatar`}
            >
              <img src={file.url} />
              <div className={`${prefixCls}-upload-list-item-picture-mask`}>
                <IconEdit />
              </div>
              {file.status === "uploading" && (file.percent || 0) < 100 && (
                <Progress
                  percent={file.percent || 0}
                  type="circle"
                  size="mini"
                  style={{
                    position: "absolute",
                    left: "50%",
                    top: "50%",
                    transform: "translateX(-50%) translateY(-50%)",
                  }}
                />
              )}
            </div>
          ) : (
            <div className={`${prefixCls}-upload-trigger-picture`}>
              <div className={`${prefixCls}-upload-trigger-picture-text`}>
                <IconPlus />
              </div>
            </div>
          )}
        </div>
      </ArcoUpload>
    );
  },
  mapReadPretty((props) => {
    const { value } = props;
    if (!isValid(value)) {
      return <div>-</div>;
    }
    return (
      <Avatar>
        <img alt="avatar" src={value} />
      </Avatar>
    );
  })
);

export const Upload: ComposedUpload = (props) => {
  return <Fragment />;
};

Upload.Attachment = UploadAttachment;
Upload.Avatar = UploadAvatar;

export default Upload;
