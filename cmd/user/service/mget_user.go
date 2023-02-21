package service

import (
	"context"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/douyinuser"
)

type MGetUserNameService struct {
	ctx context.Context
}

// NewMGetUserService new MGetUserService
func NewMGetUserNameService(ctx context.Context) *MGetUserNameService {
	return &MGetUserNameService{ctx: ctx}
}

// MGetUser multiple get list of user info
func (s *MGetUserNameService) MGetUserName(req *douyinuser.MGetUserNameRequest) (map[int64]string, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	res := make(map[int64]string, 0)
	for _, m := range modelUsers {
		res[int64(m.ID)] = m.Username
	}
	return res, nil
}
