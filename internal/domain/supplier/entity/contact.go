package entity

import (
	"context"

	"github.com/donech/tool/xdb"
	"github.com/donech/tool/xlog"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	xdb.Entity
	SupplierId int64      `gorm:"column:supplier_id;NOT NULL" json:"supplier_id"` // 供货商ID
	Firstname  string     `gorm:"column:firstname;NOT NULL" json:"firstname"`   // 名
	Lastname   string     `gorm:"column:lastname;NOT NULL" json:"lastname"`     // 姓
	Title      string     `gorm:"column:title;NOT NULL" json:"title"`             // 职称
	Relations  []Relation `gorm:"ForeignKey:contact_id"`
	xdb.CUDTime
}

func (c Contact) TableName() string {
	return "supplier_contact"
}

func (c *Contact) GetRelations(ctx context.Context, db *gorm.DB) []Relation {
	if c.Relations != nil {
		return c.Relations
	}
	c.Relations = make([]Relation, 0)
	if c.ID == 0 {
		return c.Relations
	}
	err := db.Model(c).Related(&c.Relations).Error
	if err != nil {
		xlog.S(ctx).Warnf("get relations error, contact=%v, err=%v", c, err)
	}
	return c.Relations
}
