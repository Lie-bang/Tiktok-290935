package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/dal/rdb"
	"douyin/cmd/message/pack"
	"douyin/kitex_gen/douyinmessage"
	"strconv"
)

type ChatRecordService struct {
	ctx context.Context
}

func NewChatRecordService(ctx context.Context) *ChatRecordService {
	return &ChatRecordService{ctx: ctx}
}

func (s *ChatRecordService) ChatRecord(req *douyinmessage.ChatRecordRequest) ([]*douyinmessage.Message, error) {
	res, err := rdb.GetLastTime(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return nil, err
	}

	lastTime, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return nil, err
	}
	messageModels, err := db.ChatRecord(s.ctx, req.UserId, req.ToUserId, lastTime)
	if err != nil {
		return nil, err
	}

	lens := len(messageModels)
	if lens == 0 {
		return nil, nil
	}

	err = rdb.UpdateLastTime(s.ctx, req.UserId, req.ToUserId, messageModels[lens-1].CreateTime)
	messages := pack.Messages(messageModels)
	return messages, nil
}
