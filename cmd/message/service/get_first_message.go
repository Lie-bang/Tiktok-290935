package service

import (
	"context"
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/dal/rdb"
	"douyin/cmd/message/pack"
	"douyin/kitex_gen/douyinmessage"
)

type GetFirstMessagesService struct {
	ctx context.Context
}

func NewGetFirstMessagesService(ctx context.Context) *GetFirstMessagesService {
	return &GetFirstMessagesService{ctx: ctx}
}

func (s *GetFirstMessagesService) GetFirstMessages(req *douyinmessage.GetFirstMessagesRequest) ([]*douyinmessage.Message, error) {
	messages := make([]*douyinmessage.Message, 0)
	for _, toUserId := range req.ToUserIds {
		messageModel, err := db.GetFirstMessage(s.ctx, req.UserId, toUserId)
		if err != nil {
			return nil, err
		}
		message := pack.Message(messageModel)
		messages = append(messages, message)
		err = rdb.UpdateLastTime(s.ctx, req.UserId, toUserId, 0)
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
}
