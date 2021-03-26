package entity

import (
	"context"
	"fmt"
	"testing"

	"github.com/donech/tool/xdb"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db, _ = xdb.New(xdb.Config{
		Dsn:     "root:123456@tcp(127.0.0.1:3306)/erp?charset=utf8mb4&parseTime=true&loc=Local",
		LogMode: true,
	})

}

func TestSupplier(t *testing.T) {
	var s []Supplier
	err := db.Preload("Contacts", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Relations")
	}).Find(&s).Error
	if err != nil {
		t.Fatal("Supplier error=", err)
	}
	//err = db.Model(&s).Related(&s.Contacts).Error
	//if err != nil {
	//	t.Fatal("Contact error=", err)
	//}
	fmt.Printf("%+v", s)
	//fmt.Printf("%+v", c)
}

func TestCreateSupplier(t *testing.T) {
	s := Supplier{
		Entity:  xdb.Entity{ID: 4},
		Name:    "solar",
		Address: "外国语大学",
		Contacts: []Contact{{
			Entity:    xdb.Entity{ID: 4},
			FirstName: "熊1熊",
			LastName:  "朴1",
			Title:     "教1师",
			Relations: []Relation{{
				Entity: xdb.Entity{ID: 4},
				Type:   1,
				Value:  "123456",
			}},
		}},
	}
	err := db.Save(&s).Error
	if err != nil {
		t.Fatal("create Supplier error=", err)
	}
	fmt.Printf("%+v", s)
}

func TestContact_GetRelations(t *testing.T) {
	c := Contact{
		Entity: xdb.Entity{ID: 4},
	}
	relations := c.GetRelations(context.Background(), db)
	fmt.Printf("relations: %v", relations)
	for _, relation := range relations {
		if relation.ContactId != c.ID {
			t.Fatal("GetRelations fail", c, relation)
		}
	}
}
