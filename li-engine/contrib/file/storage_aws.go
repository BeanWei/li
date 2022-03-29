package file

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

type storageAwsClient struct {
	Client *s3.S3
}

type AwsClientOption struct {
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	Region          string
	DisableSSL      bool
	ForcePathStyle  bool
}

// NewStorageAwsClient .
func NewStorageAwsClient(opt *AwsClientOption) (*storageAwsClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(opt.AccessKeyID, opt.SecretAccessKey, ""),
		Endpoint:         aws.String(opt.Endpoint),
		Region:           aws.String(opt.Region),
		DisableSSL:       aws.Bool(opt.DisableSSL),
		S3ForcePathStyle: aws.Bool(opt.ForcePathStyle),
	})
	if err != nil {
		return nil, err
	}
	return &storageAwsClient{
		Client: s3.New(sess),
	}, nil
}

func (s *storageAwsClient) PutObject(ctx context.Context, input *PutObjectInput) (*PutObjectOutput, error) {
	if input.BucketName == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, `parameter "input.BucketName" is required`)
	}
	file, err := input.File.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	key := input.filename()
	_, err = s.Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(input.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return nil, err
	}
	return &PutObjectOutput{
		FileName: key,
		FileUrl:  "/" + input.BucketName + "/" + key,
	}, nil
}

func (s *storageAwsClient) DeleteObject(ctx context.Context, input *DeleteObjectInput) error {
	_, err := s.Client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(input.BucketName),
		Key:    aws.String(input.FileName),
	})
	return err
}

func (s *storageAwsClient) ServeFile(r *ghttp.Request, input *ProxyInput) {
	objuri := "/" + input.BucketName + "/" + input.FileName
	remote, err := url.Parse(s.Client.Endpoint + objuri)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Host = s.Client.Endpoint
		req.URL.Host = req.Host
		req.URL.Path = objuri
	}
	// 浏览器缓存 30 天
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Set("Cache-Control", "max-age=2592000")
		return nil
	}
	proxy.ServeHTTP(r.Response.Writer, r.Request)
}

var _ Storage = (*storageAwsClient)(nil)
