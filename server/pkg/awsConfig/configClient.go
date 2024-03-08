package awsConfig

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Session() (*session.Session, error) {

	var region string = os.Getenv("AWS_REGION")
	var aws_access_key_id string = os.Getenv("aws_access_key_id")
	var aws_secret_access_key string = os.Getenv("aws_secret_access_key")
	var _ string = os.Getenv("aws_secret_access_key")
	// var aws_token string = os.Getenv("aws_secret_access_key")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		// Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("aws connection failed %v", err)
	}
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	return sess, nil
}
