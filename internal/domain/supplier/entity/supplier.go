package entity

import "github.com/donech/tool/xdb"

type Supplier struct {
	xdb.Entity
	Name     string    `gorm:"column:name;NOT NULL" json:"name"`       // 供货商名称
	Address  string    `gorm:"column:address;NOT NULL" json:"address"` // 供货商地址
	Contacts []Contact `gorm:"ForeignKey:supplier_id"`
	xdb.CUDTime
}

func (s Supplier) TableName() string {
	return "supplier"
}
