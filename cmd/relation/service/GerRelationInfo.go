package service

import (
	"context"
	"douyin/cmd/relation/dal/rdb"
	"douyin/kitex_gen/douyinrelation"
	"strconv"
)

type GetRelationInfoService struct {
	ctx context.Context
}

func NewGetRelationInfoService(ctx context.Context) *GetRelationInfoService {
	return &GetRelationInfoService{ctx: ctx}
}

func (s *GetRelationInfoService) GetRelationInfo(req *douyinrelation.GetRelationInfoRequest) ([]*douyinrelation.User, error) {
	followUsers := make([]*douyinrelation.User, 0)

	for _, toUserId := range req.ToUserIds {
		toIdStr := strconv.FormatInt(toUserId, 10)
		followCount, followerCount, err := rdb.CountFollowAndFollower(s.ctx, toIdStr)
		if err != nil {
			return nil, err
		}
		isFollow, err := rdb.IsFollow(s.ctx, req.UserId, toUserId)
		if err != nil {
			return nil, err
		}
		followUsers = append(followUsers, &douyinrelation.User{
			UserId:        toUserId,
			Username:      "",
			FollowCount:   followCount,
			FollowerCount: followerCount,
			IsFollow:      isFollow,
			Avatar:        "test",
		})
	}
	return followUsers, nil
}
