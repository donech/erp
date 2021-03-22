package entity

import (
	"github.com/donech/tool/xdb"
)

type Contact struct {
	xdb.Entity
	SupplierId int64      `gorm:"column:supplier_id;NOT NULL" json:"supplier_id"` // 供货商ID
	FirstName  string     `gorm:"column:first_name;NOT NULL" json:"first_name"`   // 名
	LastName   string     `gorm:"column:last_name;NOT NULL" json:"last_name"`     // 姓
	Title      string     `gorm:"column:title;NOT NULL" json:"title"`             // 职称
	Relations  []Relation `gorm:"ForeignKey:contact_id"`
	xdb.CUDTime
}

func (c Contact) TableName() string {
	return "supplier_contact"
}
