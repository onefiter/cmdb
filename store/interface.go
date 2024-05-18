package store

// Uploader 接口
// 为了屏蔽多个云厂商oss操作的差异，
// 抽象出一个store组件，用于解决文件的上传和下载问题，
// 定义一个Uploader接口
type Uploader interface {
	// 上传文件到后端存储
	Upload(bucketName, objectKey, localFilePath string) error
}
