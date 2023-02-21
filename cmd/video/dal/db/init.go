package db

import (
	"douyin/pkg/consts"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitVideo_db() {
	var err error

	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql failed")
		panic(err)
	}
	return
}
