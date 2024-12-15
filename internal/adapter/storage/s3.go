package storage

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	s3manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"

	"todo-app/pkg/config"
)

// S3Client interface defines methods for S3 operations
type S3Client interface {
	UploadFile(file multipart.File, filename string) (string, error)
}

// s3Client is the implementation of S3Client
type s3Client struct {
	client   *s3.Client
	uploader *s3manager.Uploader
	bucket   string
}

// NewS3Client initializes a new S3 client
func NewS3Client(cfg config.StorageConfig) (S3Client, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithEndpointResolver(
			aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           cfg.Endpoint,
					SigningRegion: cfg.Region,
				}, nil
			}),
		),
		config.WithCredentialsProvider(
			aws.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg)
	uploader := s3manager.NewUploader(client)

	return &s3Client{
		client:   client,
		uploader: uploader,
		bucket:   cfg.Bucket,
	}, nil
}

// UploadFile uploads a file to S3 and returns the file ID
func (s *s3Client) UploadFile(file multipart.File, filename string) (string, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	key := fmt.Sprintf("%s-%s", uuid.New().String(), filename)

	_, err := s.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buf.Bytes()),
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
