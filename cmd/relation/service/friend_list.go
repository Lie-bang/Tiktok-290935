package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
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
	var followerUsers []*douyinrelation.FriendUser
	userIds, err := db.FriendList(s.ctx, req.UserId)
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

	messages, err := rpc.GetFirstMessages(s.ctx, &douyinmessage.GetFirstMessagesRequest{
		ToUserIds: userIds,
		UserId:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	contents, msgTypes := pack.ToFirstMessage(messages, req.UserId)
	followerUsers = pack.ToFriendUser(usersMap, contents, msgTypes)
	return followerUsers, nil
}
