package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/dal/rdb"
	"douyin/kitex_gen/douyinmessage"
	"time"
)

type SendMessageService struct {
	ctx context.Context
}

func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

func (s *SendMessageService) SendMessage(req *douyinmessage.SendMessageRequest) error {
	message := &db.Message{
		UserId:      req.UserId,
		ToUserId:    req.ToUserId,
		Content:     req.Content,
		CreatedTime: time.Now().Unix(),
	}
	err := rdb.UpdateLastTime(s.ctx, req.UserId, req.ToUserId, message.CreatedTime)
	if err != nil {
		return err
	}
	return db.SendMessage(s.ctx, message)
}
