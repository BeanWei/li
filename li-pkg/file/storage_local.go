package file

import (
	"context"
	"io"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

type storageLocalClient struct {
	Dir string
}

// NewStorageAwsClient .
func NewStorageLocalClient(dir string) (*storageLocalClient, error) {
	if !gfile.IsDir(dir) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `parameter "dir" should be a directory path`)
	}
	return &storageLocalClient{
		Dir: dir,
	}, nil
}

func (s *storageLocalClient) PutObject(ctx context.Context, input *PutObjectInput) (*PutObjectOutput, error) {
	if input.BucketName == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `parameter "input.BucketName" is required`)
	}
	dir := gfile.Join(s.Dir, input.BucketName)
	if !gfile.Exists(dir) {
		if err := gfile.Mkdir(dir); err != nil {
			return nil, gerror.Wrapf(err, `StorageLocalClient.PutObject.Mkdir failed`)
		}
	}

	file, err := input.File.Open()
	if err != nil {
		return nil, gerror.Wrapf(err, `StorageLocalClient.PutObject.Open failed`)
	}
	defer file.Close()

	key := input.filename()
	filePath := gfile.Join(dir, key)
	newFile, err := gfile.Create(filePath)
	if err != nil {
		return nil, gerror.Wrapf(err, `StorageLocalClient.PutObject.CreateFilePath failed`)
	}
	defer newFile.Close()
	if _, err = io.Copy(newFile, file); err != nil {
		return nil, gerror.Wrapf(err, `io.Copy failed from "%s" to "%s"`, input.File.Filename, filePath)
	}
	return &PutObjectOutput{
		FileName: key,
		FileUrl:  "/" + input.BucketName + "/" + key,
	}, nil
}

func (s *storageLocalClient) DeleteObject(ctx context.Context, input *DeleteObjectInput) error {
	err := gfile.Remove(gfile.Join(s.Dir, input.BucketName, input.FileName))
	if err != nil {
		return gerror.Wrapf(err, `StorageLocalClient.DeleteObject.Remove failed`)
	}
	return nil
}

func (s *storageLocalClient) Proxy(r *ghttp.Request, input *ProxyInput) {
	file := gres.Get(gfile.Join(s.Dir, input.BucketName, input.FileName))
	if file != nil {
		info := file.FileInfo()
		http.ServeContent(r.Response.Writer.RawWriter(), r.Request, info.Name(), info.ModTime(), file)
	}
}

var _ Storage = (*storageLocalClient)(nil)
