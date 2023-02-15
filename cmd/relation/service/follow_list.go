package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

func (s *FollowListService) FollowList(req *douyinrelation.FollowListRequest) ([]*douyinrelation.User, error) {
	var followUsers []*douyinrelation.User
	userIds, err := db.FollowList(s.ctx, req.UserId)
	if len(userIds) == 0 {
		return followUsers, nil
	}
	usersMap, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: userIds,
		UserId: req.UserId,
	})
	if err != nil {
		return followUsers, err
	}
	followUsers = pack.UUserToRUser(usersMap)
	return followUsers, nil
}
