package system

import (
	"context"

	"github.com/donech/erp/internal/common"
	"github.com/donech/tool/xdb"

	"github.com/donech/erp/internal/domain/system/entity"
)

func CreateUser(ctx context.Context, account, name, password string) (int64, error) {
	en := entity.User{
		Account:  account,
		Name:     name,
		Password: password,
	}
	err := xdb.Trace(ctx, common.GetDB()).Create(&en).Error
	if err != nil {
		return 0, err
	}
	return en.ID, nil
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
