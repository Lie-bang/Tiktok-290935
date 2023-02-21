package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/cmd/user/pack"
	"douyin/cmd/user/rpc"
	"douyin/kitex_gen/douyinfavorite"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
	"douyin/kitex_gen/douyinvideo"
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

	favoriteCount, err := rpc.FavoriteCount(s.ctx, &douyinfavorite.DouyinFavoriteCountUserRequest{
		UserId: req.ToUserId,
	})
	if err != nil {
		return nil, err
	}

	workCount, totalFavorite, err := rpc.WorkAndFavoriteCount(s.ctx, &douyinvideo.Douyin_Work_And_Favorite_CountRequest{
		UserId: req.ToUserId,
	})
	if err != nil {
		return nil, err
	}

	return pack.User(modelUser, relationInfo[0], favoriteCount, workCount, totalFavorite), nil
}
