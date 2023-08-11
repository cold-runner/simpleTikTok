package io

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// descp 以后这些放在全局配置中
const (
	minioEndpoint  = "localhost:9001"
	minioAccessKey = "minio"
	minioSecretKey = "minio12345"
	minioBucket    = "tiktok"
	minioLocation  = "ap-southeast-1"
)

var client *minio.Client

func init() {
	var err error
	client, err = minio.New(
		minioEndpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(minioAccessKey, minioSecretKey, ""),
		},
	)
	if err != nil {
		log.Errorw("connect minio server fail", err)
		return
	}

	if err = newBucket(context.Background(), client, minioBucket); err != nil {
		log.Errorw("minio client init bucket failed", err)
		return
	}
}

func newBucket(ctx context.Context, minioClient *minio.Client, bucket string) error {
	if len(bucket) == 0 {
		log.Errorw("bucket name is empty")
	}

	exist, err := minioClient.BucketExists(ctx, bucket)
	if err == nil && exist {
		return nil
	}

	err = minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: minioLocation})
	if err != nil {
		log.Errorw("minio client init bucket failed", err)
	}
	return err
}
