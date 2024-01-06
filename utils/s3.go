package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	AWS_REGION            = os.Getenv("AWS_REGION")
	AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWS_BUCKETNAME        = os.Getenv("AWS_BUCKETNAME")
)

// Config s3 client
func CreateS3Client() (*s3.Client, error) {
	// Load aws config with region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_REGION))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	// Crate a client with accelatetion enabled
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UseAccelerate = true
	})
	if client == nil {
		return nil, errors.New("S3 client is nil")
	}

	// Confirm functioning of S3 client
	_, err = client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client: %v", err)
	}

	log.Println("S3 client created successfully")
	return client, nil
}
