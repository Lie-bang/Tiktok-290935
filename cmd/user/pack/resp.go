package pack

import (
	"douyin/kitex_gen/douyinuser"
	"douyin/pkg/errno"
	"errors"
	"time"
)

func BuildBaseResp(err error) *douyinuser.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *douyinuser.BaseResp {
	return &douyinuser.BaseResp{StatusCode: err.ErrCode, StatusMessage: &err.ErrMsg, ServiceTime: time.Now().Unix()}
}
