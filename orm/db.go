package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func DBConnect() {
	dsn := "root:90p@ssw0rd@tcp(127.0.0.1:3380)/go_jwt?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Print("err", err)
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{})
}