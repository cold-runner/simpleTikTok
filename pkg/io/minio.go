package io

import (
	"context"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
)

const (
	// contentType 上传文件的类型
	contentType = "application/octet-stream"
	// expires 以分钟为单位
	expires = 5
)

var (
	putOptions = minio.PutObjectOptions{
		ContentType: contentType,
	}
)

func PutObject(objectName string, reader io.Reader, size int64) error {
	return PutObjectWithContext(context.Background(), objectName, reader, size)
}

func PutObjectWithContext(ctx context.Context, objectName string, reader io.Reader, size int64) error {
	_, err := client.PutObject(ctx, minioBucket, objectName, reader, size, putOptions)
	if err != nil {
		log.Errorw("upload object error", err)
		return err
	}

	return nil
}

func GetObjectUrl(objectName string) (*url.URL, error) {
	return GetObjectUrlWithContext(context.Background(), objectName)
}

func GetObjectUrlWithContext(ctx context.Context, objectName string) (*url.URL, error) {
	downloadUrl, err := client.PresignedGetObject(ctx, minioBucket, objectName, expires*time.Minute, make(url.Values))
	if err != nil {
		log.Errorw("get downloadUrl of file failed", err, "file name", objectName)
		return nil, err
	}

	return downloadUrl, nil
}
