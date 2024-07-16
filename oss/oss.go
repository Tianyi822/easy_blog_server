package oss

import (
	"context"
	"github.com/Tianyi822/easy_blog_server/config"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"sync"
)

var ossOnceLock sync.Once

func ConnectOss() {
	ossOnceLock.Do(func() {
		// 连接 OSS
		provider := credentials.NewStaticCredentialsProvider(config.OssConf.AccessKeyId, config.OssConf.AccessKeySecret)
		cfg := oss.LoadDefaultConfig().WithCredentialsProvider(provider).WithRegion(config.OssConf.Region)

		client := oss.NewClient(cfg)

		request := &oss.GetBucketInfoRequest{
			Bucket: oss.Ptr(config.OssConf.Bucket),
		}
		result, err := client.GetBucketInfo(context.TODO(), request)

		if err != nil {
			logger.Panic("failed to get bucket info %v", err)
		}
		// Print the result
		logger.Info("get bucket info result:%v\n", result)
	})
}
