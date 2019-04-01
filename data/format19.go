package data

//header  cate 01  code 19  edition 01
type Receiver19 struct {
	StockCode [6]byte  //ascii

	TodayStopDealStockTime [3]byte     //bcd
	TodayReStartDealStockTime [3]byte

	TransferFlag  	[1]byte //ascii
}