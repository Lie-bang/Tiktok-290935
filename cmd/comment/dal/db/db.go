package db

import (
	"context"
	"douyin/kitex_gen/douyincomment"
	"douyin/pkg/consts"
	"log"
	"time"
)

type CommentDb struct {
	CommentId  int64
	UserId     int64
	VideoId    int64
	Content    string
	CreateDate string
	PubTime    int64
}

func (C *CommentDb) TableName() string {
	return consts.CommentTableName
}

func UpdateComment(ctx context.Context, request *douyincomment.DouyinCommentActionRequest, uid int64) (*CommentDb, error) {
	var cmt CommentDb

	cmt.UserId = uid
	cmt.VideoId = request.VideoId

	if request.ActionType == 1 {
		cmt.PubTime = time.Now().Unix()
		t := time.Now().Format("2006-01-02 15:04:05")
		cmt.CreateDate = t[5:10]
		cmt.Content = *request.CommentText
		res := DB.WithContext(ctx).Create(&cmt)
		if res.Error != nil {
			return nil, res.Error
		}

	} else if request.ActionType == 2 {
		cmt.CommentId = *request.CommentId
		res := DB.WithContext(ctx).Where("comment_id = ? AND video_id = ?", request.CommentId, cmt.VideoId).Delete(&cmt)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		log.Print("wrong action type")
		return nil, nil
	}
	return &cmt, nil
}

func QueryCommentList(ctx context.Context, request *douyincomment.DouyinCommentListRequest) (*[]CommentDb, error) {
	var cmt *[]CommentDb
	conn := DB.WithContext(ctx).Where("video_id = ?", request.VideoId).Order("pub_time desc").Find(&cmt)
	if conn.Error != nil {
		log.Print(conn.Error)
		return nil, conn.Error
	}
	return cmt, nil
}
