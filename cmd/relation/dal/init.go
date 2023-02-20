package dal

import (
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/rdb"
)

func Init() {
	db.Init()
	rdb.Init()
}
