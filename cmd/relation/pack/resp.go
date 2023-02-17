package pack

import (
	"douyin/kitex_gen/douyinrelation"
	"douyin/pkg/errno"
	"errors"
	"time"
)

func BuildBaseResp(err error) *douyinrelation.BaseResp {
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

func baseResp(err errno.ErrNo) *douyinrelation.BaseResp {
	return &douyinrelation.BaseResp{StatusCode: err.ErrCode, StatusMessage: &err.ErrMsg, ServiceTime: time.Now().Unix()}
}
