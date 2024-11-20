package dynamodb

import (
	"balances/configs"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewSession() (*session.Session, error) {
	c := configs.New()

	return session.NewSessionWithOptions(
		session.Options{
			Config: aws.Config{
				Credentials:      credentials.NewStaticCredentials(c.AwsID, c.AwsSecret, ""),
				Region:           aws.String(c.AwsRegion),
				Endpoint:         aws.String(c.AwsAddress),
				S3ForcePathStyle: aws.Bool(true),
				MaxRetries:       aws.Int(3),
			},
			Profile: c.AwsProfile,
		},
	)
}

func Connect() (*dynamodb.DynamoDB, error) {
	sess, err := NewSession()
	if err != nil {
		log.Printf("connection error: %v", err)
		return nil, err
	}
	return dynamodb.New(sess), nil
}
