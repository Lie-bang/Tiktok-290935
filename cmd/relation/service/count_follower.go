package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/kitex_gen/douyinrelation"
)

type CountFollowerService struct {
	ctx context.Context
}

func NewCountFollowerService(ctx context.Context) *CountFollowerService {
	return &CountFollowerService{ctx: ctx}
}

func (s *CountFollowerService) CountFollower(req *douyinrelation.CountFollowerRequest) (int64, error) {
	return db.CountFollower(s.ctx, req.UserId)
}
