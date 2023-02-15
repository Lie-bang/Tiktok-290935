package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/douyinrelation"
)

type IsFollowService struct {
	ctx context.Context
}

func NewIsFollowService(ctx context.Context) *IsFollowService {
	return &IsFollowService{ctx: ctx}
}

func (s *IsFollowService) IsFollow(req *douyinrelation.IsFollowRequest) (bool, error) {
	return db.IsFollow(s.ctx, req.UserId, req.ToUserId)
}
