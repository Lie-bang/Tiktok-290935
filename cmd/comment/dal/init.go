package dal

import "douyin/cmd/comment/dal/db"

func Init() {
	db.InitComment_db()
}
