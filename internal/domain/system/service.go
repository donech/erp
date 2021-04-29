package system

import (
	"context"
	"errors"
	"strconv"

	"github.com/donech/tool/tabler"
	"github.com/jinzhu/gorm"

	"github.com/donech/tool/xjwt"

	"github.com/donech/erp/internal/common"
	"github.com/donech/tool/xdb"

	"github.com/donech/erp/internal/domain/system/entity"
)

//Users 获取 system user 列表
//condition[0] query  string
//condition[1:] the args for query string
func Users(ctx context.Context, pager tabler.Pager, condition ...interface{}) ([]*entity.User, bool, error) {
	db := xdb.Trace(ctx, common.GetDB())
	var res []*entity.User
	err := tabler.NewTable(db, "id", pager).Do(func(db *gorm.DB) *gorm.DB {
		db = db.Model(entity.User{})
		if len(condition) > 1 && condition[0] != "" {
			db = db.Where(condition[0], condition[1:])
		}
		return db.Order("id desc").Find(&res)
	})
	if len(res) > int(pager.PageSize) {
		return res[0:pager.PageSize], true, err
	}
	return res, false, err
}

func SaveUser(ctx context.Context, account, name, password string) (int64, error) {
	user, err := GetUserByAccount(ctx, account)
	if gorm.IsRecordNotFoundError(err) {
		en := entity.User{
			Account:  account,
			Name:     name,
			Password: common.EncryptPassword(password),
		}
		err = xdb.Trace(ctx, common.GetDB()).Create(&en).Error
		if err != nil {
			return 0, err
		}
		return en.ID, nil
	}
	user.Name = name
	err = UpdateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func UpdateUser(ctx context.Context, user *entity.User) error {
	return xdb.Trace(ctx, common.GetDB()).Save(user).Error
}

func GetUserByAccount(ctx context.Context, account string) (*entity.User, error) {
	db := xdb.Trace(ctx, common.GetDB())
	user := entity.User{}
	err := db.Model(user).Where("account = ?", account).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(ctx context.Context, id int64) (*entity.User, error) {
	db := xdb.Trace(ctx, common.GetDB())
	user := entity.User{}
	err := db.Model(user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//AuthUser 通过解析 jwtToken 获取用户信息
//flag 控制是否去数据库读取详情
func AuthUser(ctx context.Context, flag bool) (*entity.User, error) {
	claims := xjwt.GetClaimsFromCtx(ctx)
	if id, ok := claims["id"].(string); ok {
		number, _ := strconv.ParseInt(id, 10, 64)
		if !flag {
			return &entity.User{
				Entity: xdb.Entity{ID: number},
				Name:   claims["name"].(string),
			}, nil
		}
		return GetUserById(ctx, number)
	}
	return nil, errors.New("no auth user found")
}
