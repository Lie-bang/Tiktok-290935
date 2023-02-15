package service

import (
	"context"
	"crypto/md5"
	"douyin/cmd/user/dal/db"
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/errno"
	"fmt"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *douyinuser.CreateUserRequest) error {
	user, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(user) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: password,
	}})
}
