package service

import (
	"context"
	"douyin/cmd/relation/dal/rdb"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/douyinmessage"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *douyinrelation.FriendListRequest) ([]*douyinrelation.FriendUser, error) {
	followerUsers := make([]*douyinrelation.FriendUser, 0)
	userRModels, err := rdb.FriendList(s.ctx, req.UserId, req.ToUserId)
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
		UserIds: userIds,
		UserId:  req.UserId,
	})
	if err != nil {
		return followerUsers, err
	}

	messages, err := rpc.GetFirstMessages(s.ctx, &douyinmessage.GetFirstMessagesRequest{
		ToUserIds: userIds,
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	contents, msgTypes := pack.ToFirstMessage(messages, req.UserId)
	followerUsers = pack.ToFriendUser(userRModels, usersMap, contents, msgTypes)
	return followerUsers, nil
}
