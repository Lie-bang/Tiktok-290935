package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/rpc"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *douyinuser.GetUserRequest) (*douyinuser.User, error) {
	modelUser, err := db.GetUser(s.ctx, req.ToUserId)
	if err != nil {
		return nil, err
	}
	followCount, err := rpc.CountFollow(s.ctx, &douyinrelation.CountFollowRequest{UserId: req.ToUserId})
	if err != nil {
		return nil, err
	}
	followerCount, err := rpc.CountFollower(s.ctx, &douyinrelation.CountFollowerRequest{UserId: req.ToUserId})
	if err != nil {
		return nil, err
	}
	isFollow, err := rpc.IsFollow(s.ctx, &douyinrelation.IsFollowRequest{
		UserId:   req.UserId,
		ToUserId: req.ToUserId,
	})

	return pack.User(modelUser, followCount, followerCount, isFollow), nil
}
