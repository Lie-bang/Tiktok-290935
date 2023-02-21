package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/rdb"
	"douyin/cmd/relation/rpc"
	"douyin/kitex_gen/douyinrelation"
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/errno"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

func (s *ActionService) DBAction(req *douyinrelation.ActionRequest) error {
	if req.ToUserId == req.UserId {
		return errno.ParamErr
	}
	actionModel := &db.Relation{
		FollowerId: req.ToUserId,
		UserId:     req.UserId,
		Cancel:     req.ActionType,
	}
	user, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserNameRequest{
		UserIds: []int64{req.ToUserId},
		UserId:  req.UserId,
	})
	if err != nil {
		return err
	}
	if len(user) != 1 {
		return errno.ParamErr
	}
	if req.ActionType == 1 {
		return db.Action(s.ctx, actionModel)
	} else if req.ActionType == 2 {
		return db.DeleteAction(s.ctx, actionModel)
	}
	return errno.ParamErr
}

func (s *ActionService) Action(req *douyinrelation.ActionRequest) error {
	if req.UserId == req.ToUserId {
		return errno.ParamErr
	}
	switch req.ActionType {
	case 1:
		return rdb.FollowAction(s.ctx, req.UserId, req.ToUserId)
	case 2:
		return rdb.DeleteFollowAction(s.ctx, req.UserId, req.ToUserId)
	default:
		return errno.ParamErr
	}
}
