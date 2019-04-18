package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 19  edition 01
type StockReceiver19 struct {
	StockCode [6]types.Ascii  //ascii

	TodayStopDealStockTime [3]types.Bcd     //bcd
	TodayReStartDealStockTime [3]types.Bcd

	TransferFlag  	[1]types.Ascii //ascii
}