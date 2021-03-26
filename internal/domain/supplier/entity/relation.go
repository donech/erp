package entity

import (
	"github.com/donech/tool/xdb"
)

type Relation struct {
	xdb.Entity
	ContactId int64  `gorm:"column:contact_id;NOT NULL" json:"contact_id"`
	Type      int    `gorm:"column:type;NOT NULL" json:"type"`   // 联系方式类型
	Value     string `gorm:"column:value;NOT NULL" json:"value"` // 联系方式值
	xdb.CUDTime
}

func (Relation) TableName() string {
	return "supplier_contact_relation"
}
