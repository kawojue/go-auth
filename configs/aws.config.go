package configs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kawojue/go-initenv"
)

func Session() *session.Session {
	awsAccessKeyId, bucketRegion, awsSecretAcessKey := initenv.GetEnv("AWS_SECRET_ID", ""), initenv.GetEnv("BUCKET_REGION", ""), initenv.GetEnv("AWS_ACCESS_SECRET", "")

	newSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAcessKey, ""),
	})

	sess := session.Must(newSession, err)

	return sess
}
