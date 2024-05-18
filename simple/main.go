package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	// 程序内置
	endpoint     = "oss-cn-beijing.aliyuncs.com"
	accessKey    = "xxxx"
	accessSecret = "xxxx"

	// 默认配置
	bucketName = "devcloud-station"

	// 用户需要传递的参数
	// 期望用户自己输入
	uploadFile = ""

	help = false
)

// 打印使用说明

func usage() {
	// 1.打印一些描述信息
	fmt.Fprint(os.Stderr, `cloud-station version: 0.0.1
Usage: cloud-station [-h] -f <upload_file_path>
Option:
`)

	// 2.打印有哪些参数可以使用
	flag.PrintDefaults()
}

func upload(filePath string) error {
	// 1.实例化client
	client, err := oss.New(endpoint, accessKey, accessSecret)

	if err != nil {
		return err
	}

	// 2.获取bucket对象
	bucket, err := client.Bucket(bucketName)

	if err != nil {
		return err
	}

	// 3.上传文件到bucket
	if err := bucket.PutObjectFromFile(filePath, filePath); err != nil {
		return err
	}

	// 4.打印下载链接
	downloadURL, err := bucket.SignURL(filePath, oss.HTTPGet, 60*60*24)

	if err != nil {
		return err
	}

	fmt.Printf("文件下载URL：%s", downloadURL)
	fmt.Println("请在1天之内下载.")
	return nil
}

func validate() error {
	if endpoint == "" || accessKey == "" || accessSecret == "" {
		return errors.New("endpoint, access_key, access_secret is empty")
	}

	if uploadFile == "" {
		return errors.New("upload file path required")
	}

	return nil
}

func loadParams() {
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传文件的名称")
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}
}

func main() {
	// 参数加载
	loadParams()

	// 参数校验
	if err := validate(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
		usage()
		os.Exit(1)
	}

	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("文件：%s 上传完成\n", uploadFile)
}
