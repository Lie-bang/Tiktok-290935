package minio_op

import (
	"douyin/pkg/consts"
	"github.com/minio/minio-go"
	"log"
)

const (
	endpoint        string = "127.0.0.1:9000"
	accessKeyID     string = "minioadmin"
	secretAccessKey string = "minioadmin"
	useSSL          bool   = false
)

var (
	client *minio.Client
	err    error
)

func InitMinio() {
	client, err := minio.New(consts.Endpoint, consts.AccessKeyID, consts.SecretAccessKey, consts.UseSSL)
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
	err = client.MakeBucket(consts.BucketName, "cn-south-1")
	if err != nil {
		log.Println("创建bucket错误: ", err)
		exists, _ := client.BucketExists(consts.BucketName)
		if exists {
			log.Printf("bucket: %s已经存在", consts.BucketName)
		}
	} else {
		log.Printf("Successfully created %s\n", consts.BucketName)
	}

	//bucinfo, _ := client.ListBuckets()
	//for _, bucket := range bucinfo {
	//	fmt.Println(bucket)
	//}
	//
	////bucketName := "mymusic"
	//objectName := "audit.log"
	//filePath := "./audit.txt"
	//contextType := "application/text"
	//
	//object, err := client.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contextType})
	//if err != nil {
	//	log.Println("上传失败：", err)
	//}
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, object)
	//
	//err = client.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
	//if err != nil {
	//	log.Println("下载失败", err)
	//}
	//log.Printf("下载成功")
}
