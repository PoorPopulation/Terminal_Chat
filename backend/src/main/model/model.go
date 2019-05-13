package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Nick     string
	Password string
}

type LoginLog struct {
	gorm.Model
	UserID  uint
	LoginIp string
}

type Message struct {
	gorm.Model
	Content  string
	UserId   uint
	UserName string
}
