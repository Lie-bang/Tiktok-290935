package db

import (
	"context"
	"douyin/pkg/consts"
	"errors"
	"gorm.io/gorm"
)

type Message struct {
	ID         int64 `gorm:"primaryKey"`
	UserId     int64
	ToUserId   int64
	Content    string
	CreateTime int64
}

func (m *Message) TableName() string {
	return consts.MessageTableName
}

func ChatRecord(ctx context.Context, userId, toUserId, lastTime int64) ([]*Message, error) {
	var res []*Message
	err := DB.WithContext(ctx).
		Where("user_id = ? and to_user_id = ? and create_time > ?", userId, toUserId, lastTime).
		Or("user_id = ? and to_user_id = ? and create_time > ?", toUserId, userId, lastTime).
		Order("create_time").
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SendMessage(ctx context.Context, msg *Message) error {
	return DB.WithContext(ctx).Create(msg).Error
}

func GetFirstMessage(ctx context.Context, userId, toUserId int64) (*Message, error) {
	var res *Message
	err := DB.WithContext(ctx).Where("user_id = ? and to_user_id = ?", userId, toUserId).
		Or("user_id = ? and to_user_id = ?", toUserId, userId).Last(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
