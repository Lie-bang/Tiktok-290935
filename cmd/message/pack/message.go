package pack

import (
	"douyin/cmd/message/dal/db"
	"douyin/kitex_gen/douyinmessage"
	"fmt"
)

func Message(m *db.Message) *douyinmessage.Message {
	if m == nil {
		return nil
	}

	createTime := fmt.Sprintf("%v", m.CreatedAt.Format("2006-01-02 15:04:05"))

	return &douyinmessage.Message{
		MsgId:      int64(m.ID),
		ToUserId:   m.ToUserId,
		FromUserId: m.UserId,
		Content:    m.Content,
		CreateTime: createTime,
	}
}

func Messages(ms []*db.Message) []*douyinmessage.Message {
	messages := make([]*douyinmessage.Message, 0)
	for _, m := range ms {
		if n := Message(m); n != nil {
			messages = append(messages, n)
		}
	}
	return messages
}
