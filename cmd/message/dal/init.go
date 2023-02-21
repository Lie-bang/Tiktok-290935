package dal

import (
	"douyin/cmd/message/dal/db"
	"douyin/cmd/message/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
