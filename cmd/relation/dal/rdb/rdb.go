package rdb

import (
	"context"
	"douyin/pkg/consts"
	"douyin/pkg/errno"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type User struct {
	ID            int64
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

func FollowAction(ctx context.Context, UserId, FollowerId int64) error {
	_, err := RDB.ZRank(
		ctx,
		fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
		strconv.FormatInt(FollowerId, 10)).Result()
	if err != nil {
		if err != redis.Nil {
			return errno.RedisServiceErr
		}
	} else {
		return nil
	}

	now := time.Now()

	cmds, err := RDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.ZAdd(
			ctx,
			fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
			redis.Z{
				Score:  float64(now.Unix()),
				Member: FollowerId,
			},
		)

		pipe.ZAdd(
			ctx,
			fmt.Sprintf("%s%d", consts.UserFollowerList, FollowerId),
			redis.Z{
				Score:  float64(now.Unix()),
				Member: UserId,
			},
		)

		return nil
	})
	if err != nil {
		return errno.RedisServiceErr
	}
	if cmds[0].(*redis.IntCmd).Err() != nil {
		return errno.RedisServiceErr
	}
	if cmds[1].(*redis.IntCmd).Err() != nil {
		return errno.RedisServiceErr
	}
	return nil
}

func DeleteFollowAction(ctx context.Context, UserId, FollowerId int64) error {
	_, err := RDB.ZRank(
		ctx,
		fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
		strconv.FormatInt(FollowerId, 10)).Result()
	if err != nil {
		if err != redis.Nil {
			return errno.RedisServiceErr
		}
	} else {
		return nil
	}

	_, err = RDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.ZRem(
			ctx,
			fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
			FollowerId,
		)

		pipe.ZRem(
			ctx,
			fmt.Sprintf("%s%d", consts.UserFollowerList, FollowerId),
			UserId,
		)

		return nil
	})
	if err != nil {
		return errno.RedisServiceErr
	}
	return nil
}

func FollowList(ctx context.Context, UserId, FollowId int64) ([]*User, error) {
	ids, err := RDB.ZRange(ctx, fmt.Sprintf("%s%d", consts.UserFollowList, FollowId), 0, -1).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	interIds, err := RDB.ZInter(ctx, &redis.ZStore{
		Keys: []string{
			fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
			fmt.Sprintf("%s%d", consts.UserFollowList, FollowId),
		},
	}).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	if len(interIds) == 0 {
		interIds = append(interIds, "")
	}
	size := len(ids)

	users := make([]*User, size)

	interIndex := 0
	for i := 0; i < size; i++ {
		i := i
		id := ids[i]

		var isFollow bool

		if ids[i] == interIds[interIndex] {
			isFollow = true
			interIndex++
		}

		followCnt, followerCnt, err := CountFollowAndFollower(ctx, id)
		if err != nil {
			return nil, err
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		users[i] = &User{
			ID:            idInt,
			FollowCount:   followCnt,
			FollowerCount: followerCnt,
			IsFollow:      isFollow,
		}
	}

	return users, nil
}

func FollowerList(ctx context.Context, UserId, FollowerId int64) ([]*User, error) {
	ids, err := RDB.ZRange(ctx, fmt.Sprintf("%s%d", consts.UserFollowerList, FollowerId), 0, -1).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	interIds, err := RDB.ZInter(ctx, &redis.ZStore{
		Keys: []string{
			fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
			fmt.Sprintf("%s%d", consts.UserFollowerList, FollowerId),
		},
	}).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	if len(interIds) == 0 {
		interIds = append(interIds, "")
	}

	users := make([]*User, len(ids))

	size := len(ids)

	interIndex := 0
	for i := 0; i < size; i++ {
		i := i
		id := ids[i]

		var isFollow bool

		if ids[i] == interIds[interIndex] {
			isFollow = true
			interIndex++
		}

		followCnt, followerCnt, err := CountFollowAndFollower(ctx, id)
		if err != nil {
			return nil, err
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		users[i] = &User{
			ID:            idInt,
			FollowCount:   followCnt,
			FollowerCount: followerCnt,
			IsFollow:      isFollow,
		}
	}

	return users, nil
}

func FriendList(ctx context.Context, UserId, FriendId int64) ([]*User, error) {
	ids, err := RDB.ZInter(ctx, &redis.ZStore{
		Keys: []string{
			fmt.Sprintf("%s%d", consts.UserFollowList, FriendId),
			fmt.Sprintf("%s%d", consts.UserFollowerList, FriendId),
		},
	}).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	interIds, err := RDB.ZInter(ctx, &redis.ZStore{
		Keys: []string{
			fmt.Sprintf("%s%d", consts.UserFollowList, FriendId),
			fmt.Sprintf("%s%d", consts.UserFollowerList, FriendId),
			fmt.Sprintf("%s%d", consts.UserFollowList, UserId),
		},
	}).Result()
	if err != nil {
		return nil, errno.RedisServiceErr
	}

	if len(interIds) == 0 {
		interIds = append(interIds, "")
	}

	users := make([]*User, len(ids))

	size := len(ids)

	interIndex := 0
	for i := 0; i < size; i++ {
		i := i
		id := ids[i]

		var isFollow bool

		if ids[i] == interIds[interIndex] {
			isFollow = true
			interIndex++
		}

		followCnt, followerCnt, err := CountFollowAndFollower(ctx, id)
		if err != nil {
			return nil, err
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		users[i] = &User{
			ID:            idInt,
			FollowCount:   followCnt,
			FollowerCount: followerCnt,
			IsFollow:      isFollow,
		}
	}

	return users, nil
}

func CountFollowAndFollower(ctx context.Context, id string) (int64, int64, error) {
	cmds, err := RDB.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.ZCard(ctx, consts.UserFollowList+id)
		pipe.ZCard(ctx, consts.UserFollowerList+id)

		return nil
	})
	if err != nil {
		return 0, 0, err
	}
	followCnt, err := cmds[0].(*redis.IntCmd).Result()
	if err != nil {
		return 0, 0, err
	}
	followerCnt, err := cmds[1].(*redis.IntCmd).Result()
	if err != nil {
		return 0, 0, err
	}
	return followCnt, followerCnt, nil
}

func IsFollow(ctx context.Context, userId, FollowId int64) (bool, error) {
	_, err := RDB.ZRank(ctx, fmt.Sprintf("%s%d", consts.UserFollowList, userId), fmt.Sprintf("%d", FollowId)).Result()
	if err != nil {
		if err != redis.Nil {
			return false, err
		} else {
			return false, nil
		}
	}
	return true, nil
}
