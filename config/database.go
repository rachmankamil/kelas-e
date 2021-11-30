package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbname string) {

	var err error
	db, err := gorm.Open(mysql.Open("root:masukaja@tcp(localhost:3306)/"+dbname+"?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
