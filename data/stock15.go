package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 15  edition 01
type StockReceiver15 struct {
	TodayStopDealStockCode [6]types.Ascii   //ascii
	TodayStopDealStockRes [1]types.Ascii
}