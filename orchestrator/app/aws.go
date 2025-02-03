package app

import (
	"fmt"
	"orchestrator/logger"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func isFileExist(archive string) (bool, error) {
    bucketName := os.Getenv("AWS_S3_BUCKET_NAME")
    prefixKey := os.Getenv("AWS_S3_BUCKET_KEY")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
    })
    if err != nil {
        return false, err
    }

    svc := s3.New(sess)

    input := &s3.ListObjectsV2Input{
        Bucket: aws.String(bucketName),
        Prefix: aws.String(prefixKey),
    }

    result, err := svc.ListObjectsV2(input)
    if err != nil {
        return false, err
    }

    for _, item := range result.Contents {
        if *item.Key == fmt.Sprintf("%s/%s", prefixKey, archive) {
            msg := fmt.Sprintf("Found archive in aws s3 bucket: %s", *item.Key)
            logger.Info(msg)
            return true, nil
        }
    }
    return false, nil
}
