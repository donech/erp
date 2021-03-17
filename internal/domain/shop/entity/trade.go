package entity

import "github.com/donech/tool/xdb"

type Trade struct {
	xdb.Entity
	EntityId int64 `json:"entity_id"`
	Amount   int64 `json:"amount"`
	Price    int64 `json:"price"`
	Status   int32 `json:"status"`
	xdb.CUDTime
}
