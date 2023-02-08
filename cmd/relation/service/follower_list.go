package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *douyinrelation.FollowerListRequest) ([]*douyinrelation.User, error) {
	var followerUsers []*douyinrelation.User
	userIds, err := db.FollowerList(s.ctx, req.UserId)
	if err != nil {
		return followerUsers, err
	}
	if len(userIds) == 0 {
		return followerUsers, nil
	}
	usersMap, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: userIds,
		UserId:  req.UserId,
	})
	if err != nil {
		return followerUsers, err
	}
	followerUsers = pack.UUserToRUser(usersMap)
	return followerUsers, nil
}
