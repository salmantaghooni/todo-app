package mq

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"

	"todo-app/pkg/config"
)

// SQSClient interface defines methods for SQS operations
type SQSClient interface {
	SendMessage(message interface{}) error
}

// sqsClient is the implementation of SQSClient
type sqsClient struct {
	client   *sqs.Client
	queueURL string
}

// NewSQSClient initializes a new SQS client
func NewSQSClient(cfg config.MQConfig) (SQSClient, error) {
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
	)

	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(awsCfg)

	return &sqsClient{
		client:   client,
		queueURL: cfg.QueueURL,
	}, nil
}

// SendMessage sends a message to the SQS queue
func (s *sqsClient) SendMessage(message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = s.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    aws.String(s.queueURL),
		MessageBody: aws.String(string(body)),
	})
	return err
}
