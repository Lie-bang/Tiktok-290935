package service

import (
	"context"
	"douyin/cmd/relation/dal/db"
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

func (s *ActionService) Action(req *douyinrelation.ActionRequest) error {
	if req.ToUserId == req.UserId {
		return errno.ParamErr
	}
	actionModel := &db.Relation{
		FollowerId: req.ToUserId,
		UserId:     req.UserId,
		Cancel:     req.ActionType,
	}
	user,err:=rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{
		UserIds: []int64{req.ToUserId},
		UserId:  req.UserId,
	})
	if err!=nil{
		return err
	}
	if len(user)!=1{
		return errno.ParamErr
	}
	if req.ActionType == 1 {
		return db.Action(s.ctx, actionModel)
	} else if req.ActionType == 2 {
		return db.DeleteAction(s.ctx, actionModel)
	}
	return errno.ParamErr
}
