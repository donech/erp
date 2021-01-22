package common

import (
	"github.com/donech/tool/xdb"
	"github.com/donech/tool/xlog"
	"github.com/donech/tool/xredis"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

//dbClient global dbClient
var dbClient *gorm.DB

//redisClient global redisClient
var redisClient *redis.Client

//InitGlobal init global vat and return a clean up func.
func InitGlobal() func() {
	redisViper := viper.Sub("redis")
	clean := func() {}
	if redisViper == nil {
		xlog.SS().Fatalf("init global redis error: not found config for redis in yaml")
		return clean
	}
	redisCfg := xredis.Config{}
	err := redisViper.Unmarshal(&redisCfg)
	if err != nil {
		xlog.SS().Fatalf("init global redis error: ", err)
		return clean
	}
	rd := xredis.New(redisCfg)
	RegisteGlobalRedis(rd)

	dbViper := viper.Sub("db")
	if dbViper == nil {
		xlog.SS().Fatalf("init global db error: not found config for db in yaml")
		return clean
	}
	dbCfg := xdb.Config{}
	err = dbViper.Unmarshal(&dbCfg)
	if err != nil {
		xlog.SS().Fatalf("init global dbClient error: ", err)
	}
	db, clean := xdb.Open(dbCfg)
	RegisteGlobalDB(db)
	return clean
}

//RegisteGlobalDB  registe global dbClient
func RegisteGlobalDB(db *gorm.DB) {
	dbClient = db
}

//RegisteGlobalRedis registe global redisClient
func RegisteGlobalRedis(client *redis.Client) {
	redisClient = client
}

//GetDB get global dbClient
func GetDB() *gorm.DB {
	return dbClient
}

//GetRedis get global redisClient
func GetRedis() *redis.Client {
	return redisClient
}
