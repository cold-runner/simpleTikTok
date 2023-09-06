package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
)

func UploadVideo(ctx context.Context, videoName string, videoData []byte) error {
	// 创建一个 bytes.Buffer 实例，它实现了 io.Reader 接口
	videoBuffer := bytes.NewBuffer(videoData)

	// 上传对象到 Minio 的存储桶中
	_, err := minioClient.PutObject(ctx, VideoBucketName, videoName, videoBuffer, int64(videoBuffer.Len()), minio.PutObjectOptions{ContentType: "video/mp4"})
	if err != nil {
		log.Fatalw("failed to upload video", "error", err)
		return err
	}
	return nil
}

func BuildVideoURL(videoName string) string {
	// 构建并返回视频的访问 URL
	hostAddress := viper.GetString("minio.host")
	//localAddress := viper.GetString("minio.localhost")

	port := viper.GetString("minio.port")
	videoURL := fmt.Sprintf("http://%s:%s/%s/%s", hostAddress, port, VideoBucketName, videoName)
	//videoURLLocal := fmt.Sprintf("http://%s:%s/%s/%s", localAddress, port, VideoBucketName, videoName)
	//return videoURL, videoURLLocal
	return videoURL
}

func BuildImageURL(videoName string) string {
	hostAddress := viper.GetString("minio.host")
	//localAddress := viper.GetString("minio.localhost")
	port := viper.GetString("minio.port")
	imageURL := fmt.Sprintf("http://%s:%s/%s/%s", hostAddress, port, ImageBucketName, videoName)
	//imageURLLocal := fmt.Sprintf("http://%s:%s/%s/%s", localAddress, port, ImageBucketName, videoName)
	//return imageURL, imageURLLocal
	return imageURL
}

func UploadImage(ctx context.Context, imageName string,
	imageBuf *bytes.Buffer) error {
	_, err := minioClient.PutObject(ctx, ImageBucketName, imageName, imageBuf, int64(imageBuf.Len()), minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		log.Errorw("failed to upload image", "error", err)
		return err
	}
	return nil
}

func ExtractCoverByVideoURL(url string) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(url).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, io.Discard).
		Run()
	if err != nil {
		log.Errorw("failed to extract cover", "err", err)
		return nil, err
	}
	log.Debugw("extract cover image successfully")
	return buf, nil
}
