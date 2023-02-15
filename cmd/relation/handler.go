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

// CountFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollow(ctx context.Context, req *douyinrelation.CountFollowRequest) (resp *douyinrelation.CountFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.CountFollowResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followNums, err := service.NewCountFollowService(ctx).CountFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.FollowCount = followNums
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CountFollower implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) CountFollower(ctx context.Context, req *douyinrelation.CountFollowerRequest) (resp *douyinrelation.CountFollowerResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.CountFollowerResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	followerNums, err := service.NewCountFollowerService(ctx).CountFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.FollowerCount = followerNums
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// IsFollow implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) IsFollow(ctx context.Context, req *douyinrelation.IsFollowRequest) (resp *douyinrelation.IsFollowResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinrelation.IsFollowResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	isFollow, err := service.NewIsFollowService(ctx).IsFollow(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.IsFollow = isFollow
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
