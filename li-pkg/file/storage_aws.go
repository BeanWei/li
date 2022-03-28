package file

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	return input.output(key), nil
}

func (s *storageAwsClient) DeleteObject(ctx context.Context, input *DeleteObjectInput) error {
	_, err := s.Client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(input.BucketName),
		Key:    aws.String(input.FileName),
	})
	return err
}

var _ Storage = (*storageAwsClient)(nil)
