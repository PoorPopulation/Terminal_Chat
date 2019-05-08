package model

import "time"

type User struct {
	ID int64
	Name string
	Nick string
	Password string
	CreateTime time.Time
	ModifyTime time.Time
}

type LoginLog struct {
	ID int64
	UserID int64
	LoginIp string
	CreateTime time.Time
	ModifyTime time.Time
}

type Message struct {
	ID int64
	Content string
	UserId int64
	UserName string
	CreateTime time.Time
	ModifyTime time.Time
}
