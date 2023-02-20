package dal

import "douyin/cmd/video/dal/db"

func Init() {
	db.InitVideo_db()
}
