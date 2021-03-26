package entity

import (
	"context"

	"github.com/donech/tool/xdb"
	"github.com/donech/tool/xlog"
	"github.com/jinzhu/gorm"
)

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

func (s *Supplier) GetContacts(ctx context.Context, db *gorm.DB) []Contact {
	if s.Contacts != nil {
		return s.Contacts
	}
	s.Contacts = make([]Contact, 0)
	err := db.Model(s).Related(&s.Contacts).Error
	if err != nil {
		xlog.S(ctx).Warnf("get contacts error, supplier=%v, err=%v", s, err)
	}
	return s.Contacts
}
