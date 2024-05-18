package aliyun

import (
	"errors"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/onefier/cmdb/store"
)

var (
	// 对象是否实现了接口的约束
	_ store.Uploader = &AliOssStore{}
)

type Options struct {
	Endpoint    string
	AccessKey   string
	AcessSecret string
}

func (o *Options) Validate() error {
	if o.Endpoint == "" || o.AccessKey == "" || o.AcessSecret == "" {
		return errors.New("endpoint, access_key, access_secret is empty")
	}
	return nil
}

func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:    os.Getenv("ALI_OSS_ENDPOINT"),
		AccessKey:   os.Getenv("ALI_AK"),
		AcessSecret: os.Getenv("ALI_SK"),
	})
}

func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	// 校验参数
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	c, err := oss.New(opts.Endpoint, opts.AccessKey, opts.AcessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, err
}

type AliOssStore struct {
	client *oss.Client
}

func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	bucket, err := s.client.Bucket(bucketName)

	if err != nil {
		return err
	}

	// 3.上传文件到bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return err
	}

	// 4.打印下载链接
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)

	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL：%s", downloadURL)
	fmt.Println("请在1天之内下载.")
	return nil
}
