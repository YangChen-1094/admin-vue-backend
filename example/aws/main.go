package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main(){

	accessKey := "AKIAOWWDRSMP3YGLMHSA"
	secretKey := "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"
	end_point := "http://s3.cn-north-1.amazonaws.com.cn"

	s3Session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint: aws.String(end_point),
		Region: aws.String("cn-north-1"),
		DisableSSL: aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})
	if err != nil {
		fmt.Printf("Unable to list buckets, %v", err)
		return
	}
	bucket := "game-mhjy-dev"
	svc := s3.New(s3Session)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}

	res, _ := svc.ListObjects(params)
	for _, item := range res.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}
