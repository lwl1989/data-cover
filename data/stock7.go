package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01 code 07  edition 01
type StockReceiver7 struct {
	TimeFrame [3]types.Bcd

	DealMoney [8]types.Bcd
	DealNumber [8]types.Bcd
	DealStroke [5]types.Bcd
}