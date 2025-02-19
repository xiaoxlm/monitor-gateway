package model

import (
	"github.com/xiaoxlm/monitor-gateway/config"
)

func init() {
	config.Config.Mysql.AppendModel(&User{})
}

type User struct {
	ID
	Username string
	Nickname string
	TimeAt
}

func (User) TableName() string {
	return "users"
}
