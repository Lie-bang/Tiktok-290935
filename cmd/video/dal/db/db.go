package db

import (
	"context"
	"douyin/kitex_gen/douyinvideo"
	"douyin/pkg/consts"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Video_db struct {
	Id             int
	UserId         int
	PlayUrl        string
	CoverUrl       string
	FavouriteCount int
	CommentCount   int
	IsFavourite    int
	Title          string
	PubTime        int64
}

func (v *Video_db) TableName() string {

	return consts.VideoTableName
}

// 放置操作数据库的函数
func QueryVideo(ctx context.Context, LatestTime int64) (*[]Video_db, error) {
	//以这种方式返回多条记录
	var video *[]Video_db
	fmt.Println("get in db.go QueryVideo")
	fmt.Println("db time:", LatestTime)
	conn := DB.WithContext(ctx).
		Where("pub_time < ?", LatestTime).
		Limit(30).
		Order("pub_time desc").Find(&video)

	if err := conn.Find(&video).Error; err != nil {
		fmt.Println("数据库c查询错w无")
		return nil, err
	}

	fmt.Println("get out db.go QueryVideo")

	return video, nil
}

func InsertVideo(ctx context.Context, video *douyinvideo.Video, userId int, pubTime int64) error {
	var videoDb = Video_db{
		Id:             int(video.Id),
		UserId:         userId,
		PlayUrl:        video.PlayUrl,
		CoverUrl:       video.CoverUrl,
		FavouriteCount: 0,
		CommentCount:   0,
		//IsFavourite:    0,
		Title:   video.Title,
		PubTime: pubTime,
	}
	result := DB.WithContext(ctx).Omit("Id").
		Select(
			"UserId",
			"PlayUrl",
			"CoverUrl",
			"FavouriteCount",
			"CommentCount",
			//"IsFavourite",
			"Title",
			"PubTime").
		Create(&videoDb)
	if result.Error != nil {
		log.Print("insert video failed: ", result.Error)
		return result.Error
	}
	return nil
}

func QueryListVideo(ctx context.Context, userid int) (*[]Video_db, error) {
	var video *[]Video_db
	fmt.Println("user_id now:", userid)
	conn := DB.WithContext(ctx).
		Where("user_id = ?", userid).
		Order("pub_time desc").Find(&video)
	fmt.Println(video)
	if conn.Error != nil {
		log.Print(conn.Error)
		return nil, conn.Error
	}
	return video, nil
}

func QueryFavoriteVideoList(ctx context.Context, videoIds []int64) (*[]Video_db, error) {
	//var finalVideo *[]Video_db
	fmt.Println("要遍历的video id:", videoIds)
	var video *[]Video_db
	var videolist []Video_db
	for _, v := range videoIds {
		conn := DB.WithContext(ctx).Where("id = ?", v).Find(&video)
		if conn.Error != nil {
			log.Print(conn.Error)
			return nil, conn.Error
		}
		videolist = append(videolist, (*video)[0])
		fmt.Println("**************")
		fmt.Println(video)
		//此处没想好返回的n内容怎么存储比较合适
	}
	//fmt.Println("########")
	//fmt.Println(*video)
	return &videolist, nil
}

func UpdateCommentCount(ctx context.Context, req *douyinvideo.DouyinCommentCountRequest) error {
	var video Video_db
	if req.ActionType == 1 {

		result := DB.WithContext(ctx).Model(&video).Where("id = ?", req.VideoId).
			Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if result.Error != nil {
			return result.Error
		}

	} else if req.ActionType == 2 {

		result := DB.WithContext(ctx).Model(&video).Where("id = ?", req.VideoId).
			Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if result.Error != nil {
			return result.Error
		}

	} else {
		log.Print("wrong action type for update comment count")
		return nil
	}
	return nil
}

func QueryWorkAndFavoriteCount(ctx context.Context, req *douyinvideo.Douyin_Work_And_Favorite_CountRequest) (int64, int64, error) {
	var video []Video_db
	result := DB.WithContext(ctx).Where("user_id = ?", req.UserId).Find(&video)
	if result.Error != nil {
		return 0, 0, result.Error
	}

	var work_count int64
	work_count = result.RowsAffected
	var fav_count int64
	fav_count = 0

	for _, v := range video {
		fav_count = fav_count + int64(v.FavouriteCount)
	}
	return work_count, fav_count, nil
}
