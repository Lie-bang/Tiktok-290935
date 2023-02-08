package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/rpc"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserService) MGetUser(req *douyinuser.MGetUserRequest) ([]*douyinuser.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	var followCounts, followerCounts []int64
	var isFollows []bool
	for _, user := range modelUsers {
		userId := user.ID
		followCount, err := rpc.CountFollow(s.ctx, &douyinrelation.CountFollowRequest{UserId: int64(userId)})
		if err != nil {
			return nil, err
		}
		followCounts = append(followCounts, followCount)
		followerCount, err := rpc.CountFollower(s.ctx, &douyinrelation.CountFollowerRequest{UserId: int64(userId)})
		if err != nil {
			return nil, err
		}
		followerCounts = append(followerCounts, followerCount)
		isFollow, err := rpc.IsFollow(s.ctx, &douyinrelation.IsFollowRequest{
			UserId:   req.UserId,
			ToUserId: int64(userId),
		})
		isFollows = append(isFollows, isFollow)
	}
	return pack.Users(modelUsers, followCounts, followerCounts, isFollows), nil
}
