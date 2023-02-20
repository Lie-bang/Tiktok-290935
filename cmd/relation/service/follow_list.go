package service

import (
	"context"
	"douyin/cmd/relation/dal/rdb"
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
	followUsers := make([]*douyinrelation.User, 0)
	userRModels, err := rdb.FollowList(s.ctx, req.UserId, req.ToUserId)
	if len(userRModels) == 0 {
		return followUsers, nil
	}
	if err != nil {
		return nil, err
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
		return followUsers, err
	}
	followUsers = pack.ToRUser(userRModels, usersMap)
	return followUsers, nil
}
