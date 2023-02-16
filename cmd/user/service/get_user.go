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

	relationInfo, err := rpc.GetRelationInfo(s.ctx, &douyinrelation.GetRelationInfoRequest{
		UserId:    req.UserId,
		ToUserIds: []int64{req.ToUserId},
	})
	if err != nil {
		return nil, err
	}

	return pack.User(modelUser, relationInfo[0]), nil
}
