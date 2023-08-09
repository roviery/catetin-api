package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Fullname string `gorm:"column:fullname"`
	Email    string `gorm:"column:email;unique_index"`
	Password string `gorm:"column:password;"`
}
