package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/douyincomment"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (c *CommentActionService) CommentAction(req *douyincomment.DouyinCommentActionRequest, uid int64) (*douyincomment.Comment, error) {
	resp, err := db.UpdateComment(c.ctx, req, uid)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = rpc.CommentCountUpdate(context.Background(), &douyinvideo.DouyinCommentCountRequest{
		VideoId:    req.VideoId,
		ActionType: int64(req.ActionType),
	})
	if err != nil {
		return nil, err
	}

	cmt, err := pack.CommentDbToDouyinComment(resp)
	if err != nil {
		return nil, err
	}

	return cmt, nil
}
