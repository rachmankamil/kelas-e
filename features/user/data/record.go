package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);unique"`
	Password string `gorm:"type:varchar(255)"`
	Gender   string `gorm:"type:varchar(10)"`
	Address  string
}
