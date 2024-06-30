package store

import (
	"context"
	"net/http"
	"strings"

	"github.com/quincy0/qpro/qLog"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"go.uber.org/zap"
)

var (
	FilePrefix = "https://aeyes.tos-cn-beijing.volces.com/audio/"
	tosClient  *tos.ClientV2
	accessKey  = "AKLTZGY5MTBmZGNmMmU5NDJjNmFkOThkY2IwOWMwOTJiYzQ"
	secretKey  = "TTJNNU1qaG1NakJtTm1RMk5HUTRPVGd4TkdRNE5UZzJNV1kzTm1Rek5XRQ=="
	// Bucket 对应的 Endpoint，以华北2（北京）为例：https://tos-cn-beijing.volces.com
	endpoint = "https://tos-cn-beijing.volces.com"
	region   = "cn-beijing"
	// 填写 BucketName
	bucketName = "aeyes"

	// 将文件上传到 example_dir 目录下的 example.txt 文件
	objectDir = "audio/"
)

func init() {
	var err error
	// 初始化客户端
	tosClient, err = tos.NewClientV2(
		endpoint,
		tos.WithRegion(region),
		tos.WithCredentials(tos.NewStaticCredentials(accessKey, secretKey)),
		tos.WithMaxConnections(10),
	)
	if err != nil {
		panic(err)
	}
	//defer tosClient.Close()
}

func Upload(ctx context.Context, name string, data string) error {
	body := strings.NewReader(data)
	_, err := tosClient.PutObjectV2(ctx, &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: bucketName,
			Key:    objectDir + name,
		},
		Content: body,
	})
	return err
}

// name xxx.wav
func IsExist(ctx context.Context, name string) bool {
	_, err := tosClient.HeadObjectV2(ctx, &tos.HeadObjectV2Input{
		Bucket: bucketName,
		Key:    objectDir + name,
	})
	if err != nil {
		if serverErr, ok := err.(*tos.TosServerError); ok {
			// 判断对象是否存在
			if serverErr.StatusCode == http.StatusNotFound {
				//zlog.TraceInfo(ctx, "Object not found.", zap.String("name", name))
			} else {
				qLog.TraceError(ctx, "check failed", zap.Error(err))
			}
		} else {
			qLog.TraceError(ctx, "check failed", zap.Error(err))
		}
		return false
	}
	return true
}
