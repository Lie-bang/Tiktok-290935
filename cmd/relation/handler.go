package main

import (
	"context"
	"douyin/cmd/relation/pack"
	"douyin/cmd/relation/service"
	"douyin/kitex_gen/douyinrelation"
	"douyin/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Action implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Action(ctx context.Context, req *douyinrelation.ActionRequest) (resp *douyinrelation.ActionResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.ActionResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewActionService(ctx).Action(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowList(ctx context.Context, req *douyinrelation.FollowListRequest) (resp *douyinrelation.FollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.FollowListResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followUsers, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = followUsers
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowerList(ctx context.Context, req *douyinrelation.FollowerListRequest) (resp *douyinrelation.FollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.FollowerListResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followerUsers, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = followerUsers
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FriendList(ctx context.Context, req *douyinrelation.FriendListRequest) (resp *douyinrelation.FriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.FriendListResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	friendUsers, err := service.NewFriendListService(ctx).FriendList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = friendUsers
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetRelationInfo implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetRelationInfo(ctx context.Context, req *douyinrelation.GetRelationInfoRequest) (resp *douyinrelation.GetRelationInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.GetRelationInfoResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	usersInfo, err := service.NewGetRelationInfoService(ctx).GetRelationInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserList = usersInfo
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
