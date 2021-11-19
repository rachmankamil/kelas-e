package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	db, err := gorm.Open(mysql.Open("root:masukaja@/kelas-e?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db.Debug()
}
