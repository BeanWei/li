package service

import (
	"context"
	"sync"

	"github.com/BeanWei/li/li-engine/contrib/file"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	fileclent     file.Storage
	fileclentOnce sync.Once
)

func NewFileClient(ctx context.Context) (file.Storage, error) {
	var err error
	fileclentOnce.Do(func() {
		adapter := g.Cfg().MustGet(ctx, "file.default.adapter").String()
		switch adapter {
		case "aws":
			fileclent, err = file.NewStorageAwsClient(&file.AwsClientOption{
				AccessKeyID:     g.Cfg().MustGet(ctx, "file.aws.accessKeyID").String(),
				SecretAccessKey: g.Cfg().MustGet(ctx, "file.aws.secretAccessKey").String(),
				Endpoint:        g.Cfg().MustGet(ctx, "file.aws.endpoint").String(),
				Region:          g.Cfg().MustGet(ctx, "file.aws.region").String(),
				DisableSSL:      g.Cfg().MustGet(ctx, "file.aws.disableSSL").Bool(),
				ForcePathStyle:  g.Cfg().MustGet(ctx, "file.aws.forcePathStyle").Bool(),
			})
		default:
			fileclent, err = file.NewStorageLocalClient(&file.LocalClientOption{
				Dir: g.Cfg().MustGet(ctx, "file.local.dir").String(),
			})
		}
	})
	if err != nil {
		return nil, err
	}
	return fileclent, nil
}
