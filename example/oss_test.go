package example

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets: ", bucket.Name)
	}

}

// TestUploadFile 测试阿里云OssSDK PutObjectFromFile接口
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		t.Log(err)
	}

	// 上传文件到bucket中
	// 云商OssServer会根据你的 key自动创建目录
	err = bucket.PutObjectFromFile("mydir/oss_test.go", "oss_test.go")

	if err != nil {
		t.Log(err)
	}

}

func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)

	fmt.Println(OssEndpoint)
	if err != nil {
		panic(err)
	}

	client = c

}
