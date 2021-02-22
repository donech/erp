package entity

import "github.com/donech/tool/xdb"

type User struct {
	xdb.Entity
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"password"`
	xdb.CUDTime
}

func (receiver User) TableName() string {
	return "system_user"
}
