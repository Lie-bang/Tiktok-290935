package service

import (
	"context"
	"douyin/cmd/relation/dal/rdb"
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
	followerUsers := make([]*douyinrelation.User, 0)
	userRModels, err := rdb.FollowerList(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return followerUsers, err
	}
	if len(userRModels) == 0 {
		return followerUsers, nil
	}

	userIds := make([]int64, 0)
	for _, rm := range userRModels {
		userIds = append(userIds, rm.ID)
	}
	usersMap, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserNameRequest{
		UserId:  req.UserId,
		UserIds: userIds,
	})
	if err != nil {
		return followerUsers, err
	}
	followerUsers = pack.ToRUser(userRModels, usersMap)

	return followerUsers, nil
}
