package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/kawojue/go-auth/configs"
	"github.com/kawojue/go-initenv"
)

func UploadS3(ctx *gin.Context, filePath string, fileBytes []byte) {
	bucketName := initenv.GetEnv("BUCKET_NAME", "")

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(filePath),
		Body:          bytes.NewReader(fileBytes),
		ContentType:   aws.String(http.DetectContentType(fileBytes)),
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	s3Client := s3.New(configs.Session())

	_, err := s3Client.PutObject(params)

	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Error uploading file")
		return
	}
}

func DeleteS3(ctx *gin.Context, path string) {
	bucketName := initenv.GetEnv("BUCKET_NAME", "")

	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path),
	}

	s3Client := s3.New(configs.Session())

	_, err := s3Client.DeleteObject(params)

	if err != nil {
		SendError(ctx, http.StatusInternalServerError, "Error deleting file")
		return
	}
}

func GetS3(path string) string {
	distDomain := initenv.GetEnv("DIST_DOMAIN", "")

	if distDomain == "" {
		panic(errors.New("distrubution domain is empty"))
	}

	return fmt.Sprintf("%s/%s", distDomain, path)
}
