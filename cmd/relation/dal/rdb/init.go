package rdb

import (
	"context"
	"douyin/pkg/consts"
	"github.com/redis/go-redis/v9"
	"time"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",     // redis地址
		Password: consts.RedisPassword, // redis密码，没有则留空
		DB:       0,                    // 默认数据库，默认是0

		// 命令执行失败时的重试策略
		MaxRetries:      3,                      // 命令执行失败时, 最多重试多少次, 默认为0即不错红石
		MinRetryBackoff: 8 * time.Millisecond,   // 每次计算重试间隔时间的下限, 默认8毫秒, -1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, // 每次计算重试间隔时间的上线, 默认512毫秒, -1表示取消间隔

		// 超时
		DialTimeout:  5 * time.Second, // 连接建立超时时间, 默认5秒
		ReadTimeout:  3 * time.Second, // 读超时, 默认3秒, -1表示取消读超时
		WriteTimeout: 3 * time.Second, // 写超时, 默认等于读超时
		PoolTimeout:  4 * time.Second, // 当所有连接都处在繁忙状态时, 客户端等待可用连接的最大等待时长, 默认为读超时+1秒
	})
	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
