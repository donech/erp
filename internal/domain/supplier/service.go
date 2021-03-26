package supplier

import (
	"context"

	"github.com/donech/erp/internal/common"

	"github.com/donech/erp/internal/domain/supplier/entity"
	"github.com/donech/tool/xlog"
	"github.com/jinzhu/gorm"
)

func CreateSupplier(ctx context.Context, db *gorm.DB, s entity.Supplier) (int64, error) {
	if db == nil {
		db = common.GetDB()
	}
	err := db.Save(&s).Error
	if err != nil {
		xlog.S(ctx).Errorf("CreateSupplier error, supplier=%v, err=%v", s, err)
		return 0, err
	}
	return s.ID, nil
}

func CreateContact(ctx context.Context, db *gorm.DB, s entity.Contact) (int64, error) {
	if db == nil {
		db = common.GetDB()
	}
	err := db.Save(&s).Error
	if err != nil {
		xlog.S(ctx).Errorf("CreateContact error, Contact=%v, err=%v", s, err)
		return 0, err
	}
	return s.ID, nil
}

func CreateRelation(ctx context.Context, s entity.Relation, db *gorm.DB) (int64, error) {
	if db == nil {
		db = common.GetDB()
	}
	err := db.Save(&s).Error
	if err != nil {
		xlog.S(ctx).Errorf("CreateRelation error, Relation=%v, err=%v", s, err)
		return 0, err
	}
	return s.ID, nil
}

func GetSupplierByPage(ctx context.Context, db *gorm.DB, page, pageSize int) ([]entity.Supplier, error) {
	if db == nil {
		db = common.GetDB()
	}
	if page < 1 {
		page = 1
	}
	var s  []entity.Supplier
	err := db.Preload("Contacts", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Relations")
	}).Find(&s).Offset((page-1) * pageSize).Limit(pageSize).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetSupplierByPage error, page=%d,pageSize=%d, err=%v", page, pageSize, err)
		return nil, err
	}
	return s, nil
}

func GetSupplierByID(ctx context.Context, db *gorm.DB, id int64) (entity.Supplier, error) {
	if db == nil {
		db = common.GetDB()
	}
	s := entity.Supplier{}
	err := db.Preload("Contacts", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Relations")
	}).Where("id = ?", id).Find(&s).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetSupplierByID error, id=%d, err=%v", id, err)
		return entity.Supplier{}, err
	}
	return s, nil
}

func GetSuppliersByName(ctx context.Context, db *gorm.DB, name string) ([]entity.Supplier, error) {
	if db == nil {
		db = common.GetDB()
	}
	var ss []entity.Supplier
	err := db.Preload("Contacts", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Relations")
	}).Where("name like '%?%'", name).Find(&ss).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetSuppliersByName error, name=%s, err=%v", name, err)
		return nil, err
	}
	return ss, nil
}

func GetContractByID(ctx context.Context, db *gorm.DB, id int64) (entity.Contact, error) {
	if db == nil {
		db = common.GetDB()
	}
	c := entity.Contact{}
	err := db.Preload("Relations").Where("id = ?", id).Find(&c).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetContractByID error, id=%d, err=%v", id, err)
		return entity.Contact{}, err
	}
	return c, nil
}

func GetRelationByID(ctx context.Context, db *gorm.DB, id int64) (entity.Relation, error) {
	if db == nil {
		db = common.GetDB()
	}
	c := entity.Relation{}
	err := db.Where("id = ?", id).Find(&c).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetContractByID error, id=%d, err=%v", id, err)
		return entity.Relation{}, err
	}
	return c, nil
}

func GetContractsByName(ctx context.Context, db *gorm.DB, name string) ([]entity.Contact, error) {
	if db == nil {
		db = common.GetDB()
	}
	var cs []entity.Contact
	err := db.Preload("Relations").Where(`name like "%?%"`, name).Find(&cs).Error
	if err != nil {
		xlog.S(ctx).Errorf("GetContractsByName error, name=%s, err=%v", name, err)
		return nil, err
	}
	return cs, nil
}

func DeleteEntityByID(ctx context.Context, tx *gorm.DB, entity interface{}, id int64) error {
	if tx == nil {
		tx = common.GetDB()
	}
	return tx.Where("id = ?", id).Delete(entity).Error
}
