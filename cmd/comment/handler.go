package main

import (
	"context"
	"douyin/cmd/comment/service"
	"douyin/kitex_gen/douyincomment"
	"log"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, request *douyincomment.DouyinCommentActionRequest, uid int64) (resp *douyincomment.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	result, err := service.NewCommentActionService(ctx).CommentAction(request, uid)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyincomment.DouyinCommentActionResponse)
	resp.Comment = result
	resp.StatusCode = 0
	StatusMsg := "Comment Action Success"
	resp.StatusMsg = &StatusMsg
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, request *douyincomment.DouyinCommentListRequest) (resp *douyincomment.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	result, err := service.NewCommentListService(ctx).CommentList(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	resp = new(douyincomment.DouyinCommentListResponse)
	resp.CommentList = result
	resp.StatusCode = 0
	StatusMsg := "Get Comment List Success"
	resp.StatusMsg = &StatusMsg
	return resp, err
}
