package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	db, err := gorm.Open(mysql.Open("root:masukaja@tcp(37.59.55.185:3306)/kelas-e?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
