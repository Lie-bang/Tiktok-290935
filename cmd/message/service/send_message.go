package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/douyinmessage"
)

type SendMessageService struct {
	ctx context.Context
}

func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

func (s *SendMessageService) SendMessage(req *douyinmessage.SendMessageRequest) error {
	message := &db.Message{
		UserId:   req.UserId,
		ToUserId: req.ToUserId,
		Content:  req.Content,
	}
	return db.SendMessage(s.ctx, message)
}
