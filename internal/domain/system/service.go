package system

import (
	"context"
	"errors"

	"github.com/donech/tool/xjwt"

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
	if id, ok := claims["id"]; ok {
		if !flag {
			return &entity.User{
				Entity: xdb.Entity{ID: int64(id.(float64))},
				Name:   claims["name"].(string),
			}, nil
		}
		return GetUserById(ctx, int64(id.(float64)))
	}
	return nil, errors.New("no auth user found")
}
