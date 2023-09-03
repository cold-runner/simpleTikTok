package minio

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var minioClient *minio.Client
var VideoBucketName string
var ImageBucketName string

func InitializeMinIO() {
	client, err := minio.New(viper.GetString("minio.url"), &minio.Options{
		Creds: credentials.NewStaticV4(viper.GetString("minio.access-key"),
			viper.GetString("minio.secret-key"), ""),
	})
	if err != nil {
		log.Fatalw("failed to connect to minio", "error", err)
	}
	log.Debugw("minio client init successfully")
	VideoBucketName = viper.GetString("minio.bucket.video")
	ImageBucketName = viper.GetString("minio.bucket.image")
	minioClient = client
	createBucketIfNotExist(minioClient, VideoBucketName)
	createBucketIfNotExist(minioClient, ImageBucketName)
}

func createBucketIfNotExist(client *minio.Client, bucketName string) {
	exists, err := client.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalw("failed to check if bucket exists", "error", err)
	}
	if !exists {
		err = client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalw("failed to create bucket", "error", err)
		}
		log.Debugw("bucket created successfully")
	} else {
		log.Debugw("bucket already exists")
	}
	err = client.SetBucketPolicy(context.Background(), bucketName, getPolicy(bucketName))
	if err != nil {
		log.Errorw("failed to set bucket policy", "error", err)
	}
	log.Debugw("bucket policy set successfully", "bucket", bucketName)

}

func getPolicy(bucketName string) string {
	return `{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Sid": "",
			"Effect": "Allow",
			"Principal": {
				"AWS": "*"
			},
			"Action": [
				"s3:GetObject"
			],
			"Resource": [
				"arn:aws:s3:::` + bucketName + `/*"
			]
		}
	]
}`
}

func GetMinioClient() *minio.Client {
	return minioClient
}
