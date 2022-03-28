package file

import (
	"context"
	"io"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
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
	dir := s.Dir
	if input.BucketName != "" {
		dir = gfile.Join(dir, input.BucketName)
	}
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
	return input.output(key), nil
}

func (s *storageLocalClient) DeleteObject(ctx context.Context, input *DeleteObjectInput) error {
	err := gfile.Remove(gfile.Join(s.Dir, input.BucketName, input.FileName))
	if err != nil {
		return gerror.Wrapf(err, `StorageLocalClient.DeleteObject.Remove failed`)
	}
	return nil
}

var _ Storage = (*storageLocalClient)(nil)
