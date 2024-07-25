package oss

import (
	"context"
	"fmt"
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"sync"
)

var ossOnceLock sync.Once

var ossClient *oss.Client

func ConnectOss() {
	ossOnceLock.Do(func() {
		// 连接 OSS
		provider := credentials.NewStaticCredentialsProvider(config.OssConf.AccessKeyId, config.OssConf.AccessKeySecret)
		cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(config.OssConf.Region)

		// 创建 OSS Client
		ossClient = oss.NewClient(cfg)

		// 获取 Bucket 信息
		request := &oss.GetBucketInfoRequest{
			Bucket: oss.Ptr(config.OssConf.Bucket),
		}
		result, err := ossClient.GetBucketInfo(context.Background(), request)
		if err != nil {
			logger.Panic("failed to get bucket info %v", err)
		}

		// Print the result
		logger.Info("get bucket info result:%v\n", result.StatusCode)
	})
}

// PutFile 上传文件
// ctx 上下文
// filePath 上传的本地文件路径
// objectName 上传目标对象的名称（如果要放到子目录下，需要带上子目录名称：temp/test.txt）
func PutFile(ctx context.Context, filePath, objectName string) error {
	// 开启上传文件请求
	result, err := ossClient.PutObjectFromFile(context.TODO(), &oss.PutObjectRequest{
		Bucket: oss.Ptr(config.OssConf.Bucket),
		Key:    oss.Ptr(objectName),
	}, filePath)

	// 上传文件失败
	if err != nil {
		msg := fmt.Sprintf("上传文件失败 %v", err)
		logger.Fatal(msg)
		return err
	}
	logger.Info("上传文件成功: %#v\n", result.Status)
	return nil
}
