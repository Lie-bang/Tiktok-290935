package rdb

import (
	"context"
	"douyin/pkg/consts"
	"fmt"
)

func GetLastTime(ctx context.Context, userId, toUserId int64) (string, error) {
	res, err := RDB.Get(ctx, fmt.Sprintf("%s%d%s%d", consts.RedisChatRecord, userId, "to", toUserId)).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func UpdateLastTime(ctx context.Context, userId, toUserId, lastTime int64) error {
	err := RDB.Set(ctx,
		fmt.Sprintf("%s%d%s%d", consts.RedisChatRecord, userId, "to", toUserId),
		lastTime,
		0,
	).Err()
	if err != nil {
		return err
	}
	return nil
}
