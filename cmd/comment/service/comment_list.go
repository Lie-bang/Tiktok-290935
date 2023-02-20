package service

import (
	"context"
	"douyin/cmd/comment/dal/db"
	"douyin/cmd/comment/pack"
	"douyin/kitex_gen/douyincomment"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

func (c *CommentListService) CommentList(req *douyincomment.DouyinCommentListRequest) ([]*douyincomment.Comment, error) {
	resp, err := db.QueryCommentList(c.ctx, req)
	if err != nil {
		return nil, err
	}

	temp := *resp
	var cmtList []*douyincomment.Comment
	for _, v := range temp {
		cmt, err := pack.CommentDbToDouyinComment(&v)
		if err != nil {
			return nil, err
		}
		cmtList = append(cmtList, cmt)
	}

	return cmtList, nil
}
