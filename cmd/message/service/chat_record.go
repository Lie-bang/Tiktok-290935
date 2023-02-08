package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/pack"
	"douyin/kitex_gen/douyinmessage"
)

type ChatRecordService struct {
	ctx context.Context
}

func NewChatRecordService(ctx context.Context) *ChatRecordService {
	return &ChatRecordService{ctx: ctx}
}

func (s *ChatRecordService) ChatRecord(req *douyinmessage.ChatRecordRequest) ([]*douyinmessage.Message, error) {
	messageModels, err := db.ChatRecord(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return nil, err
	}

	messages := pack.Messages(messageModels)
	return messages, nil
}
