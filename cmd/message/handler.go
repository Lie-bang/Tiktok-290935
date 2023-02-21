package main

import (
	"context"
	"douyin/cmd/message/pack"
	"douyin/cmd/message/service"
	douyinmessage "douyin/kitex_gen/douyinmessage"
	"douyin/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// ChatRecord implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) ChatRecord(ctx context.Context, req *douyinmessage.ChatRecordRequest) (resp *douyinmessage.ChatRecordResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinmessage.ChatRecordResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	messages, err := service.NewChatRecordService(ctx).ChatRecord(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.MsgList = messages
	return resp, nil
}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *douyinmessage.SendMessageRequest) (resp *douyinmessage.SendMessageResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinmessage.SendMessageResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	err = service.NewSendMessageService(ctx).SendMessage(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFirstMessages implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetFirstMessages(ctx context.Context, req *douyinmessage.GetFirstMessagesRequest) (resp *douyinmessage.GetFirstMessagesResponse, err error) {
	// TODO: Your code here...
	resp = new(douyinmessage.GetFirstMessagesResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, err
	}

	message, err := service.NewGetFirstMessagesService(ctx).GetFirstMessages(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Messages = message
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
