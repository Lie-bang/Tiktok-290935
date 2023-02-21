package Service

import (
	"bytes"
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/douyinvideo"

	"douyin/pkg/consts"
	"fmt"
	"github.com/minio/minio-go"
	"log"
	"os"
	"strconv"
	"time"
)

type PublishVideoService struct {
	ctx context.Context
}

func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (p *PublishVideoService) PublishVideo(req *douyinvideo.DouyinPublishActionRequest) error {
	//首先将传过来的数据写入到minio
	//将minio的结果写入到数据库
	client, err := minio.New(
		consts.Endpoint,
		consts.AccessKeyID,
		consts.SecretAccessKey,
		consts.UseSSL,
	)
	if err != nil {
		log.Print(err)
		return err
	}
	//视频的命名方式为: user_id + title
	n, err := client.PutObjectWithContext(
		p.ctx,
		consts.BucketName,
		req.Token+"_"+req.Title+".mp4",
		bytes.NewBuffer(req.Data),
		int64(len(req.Data)),
		minio.PutObjectOptions{ContentType: "video/mp4"},
	)
	if err != nil {
		log.Print("上传失败")
		log.Print(err)
		return err
	}
	//记录上传到服务器的时间
	t := time.Now().Unix()
	log.Printf("upload to minio %s of size %d", req.Token+req.Title, n)
	fmt.Println("object name:", req.Token+"_"+req.Title+".mp4")
	err = client.FGetObjectWithContext(
		p.ctx,
		consts.BucketName,
		req.Token+"_"+req.Title+".mp4",
		consts.TempVideoFilePath+req.Token+"_"+req.Title+".mp4",
		minio.GetObjectOptions{},
	)
	if err != nil {
		log.Print("下载失败：", err)
		return err
	}

	imgPath, err := ffmpegPicture(consts.TempVideoFilePath + req.Token + "_" + req.Title)
	if err != nil {
		log.Print(err)
		return err
	}
	fmt.Println("截图成功: ", imgPath)

	picObject, err := client.FPutObjectWithContext(
		p.ctx,
		consts.BucketName,
		req.Token+"_"+req.Title+".png",
		imgPath,
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		log.Print("上传截图失败：", err)
		return err
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", req.Token+req.Title+".png", picObject)

	err = os.Remove(consts.TempVideoFilePath + req.Token + "_" + req.Title + ".mp4")
	if err != nil {
		log.Print("删除本地视频文件失败:", err)
		return err
	}
	err = os.Remove(consts.TempVideoFilePath + req.Token + "_" + req.Title + ".png")
	if err != nil {
		log.Print("删除本地视频截图失败：", err)
		return err
	}

	//开始准备写入数据库
	var video = douyinvideo.Video{
		Id:            0,
		Author:        nil,
		PlayUrl:       consts.HttpTemplate + consts.LocalIp + consts.MinioPort + "/" + consts.BucketName + "/" + req.Token + "_" + req.Title + ".mp4",
		CoverUrl:      consts.HttpTemplate + consts.LocalIp + consts.MinioPort + "/" + consts.BucketName + "/" + req.Token + "_" + req.Title + ".png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         req.Title,
	}
	token, _ := strconv.Atoi(req.Token)

	err = db.InsertVideo(p.ctx, &video, token, t)
	if err != nil {
		log.Print("insert video failed: ", err)
		return err
	}

	return nil
}
