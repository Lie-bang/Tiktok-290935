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
	"fmt"
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
	fmt.Println("ready get into rpc.FavoriteCount")
	fmt.Println("requ.UserId now: ", req.UserId)
	favoriteCount, err := rpc.FavoriteCount(s.ctx, &douyinfavorite.DouyinFavoriteCountUserRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("favoriteCount:", favoriteCount)

	workCount, totalFavorite, err := rpc.WorkAndFavoriteCount(s.ctx, &douyinvideo.Douyin_Work_And_Favorite_CountRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	//var favoriteCount int64 = 1
	//var workCount int64 = 1
	//var totalFavorite int64 = 1

	return pack.UserG(modelUser, relationInfo[0], favoriteCount, workCount, totalFavorite), nil
}
