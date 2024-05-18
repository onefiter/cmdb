package aliyun_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/onefier/cmdb/store"
	"github.com/onefier/cmdb/store/aliyun"
)

var (
	uploader store.Uploader
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

// Aliyun Oss Store Upload测试用例
func TestUpload(t *testing.T) {
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_test.go")

	if should.NoError(err) {
		// 没有Error，开启下一个步骤
		t.Log("upload ok")
	}
}

// 通过init 来编写 uploader 实例化逻辑
func init() {
	ali, err := aliyun.NewDefaultAliOssStore()

	if err != nil {
		panic(err)
	}
	uploader = ali
}
