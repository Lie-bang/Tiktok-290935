package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/douyinrelation"
)

type CountFollowService struct {
	ctx context.Context
}

func NewCountFollowService(ctx context.Context) *CountFollowService {
	return &CountFollowService{ctx: ctx}
}

func (s *CountFollowService) CountFollow(req *douyinrelation.CountFollowRequest) (int64, error) {
	return db.CountFollow(s.ctx, req.UserId)
}
