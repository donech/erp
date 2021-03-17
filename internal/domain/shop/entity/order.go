package entity

import "github.com/donech/tool/xdb"

type Order struct {
	xdb.Entity
	TradeId  int64 `json:"trade_id"`
	EntityId int64 `json:"entity_id"`
	Amount   int64 `json:"amount"`
	Price    int64 `json:"price"`
	xdb.CUDTime
}
