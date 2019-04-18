package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 12  edition 02
type StockReceiver12 struct {
	BondNumber [1]types.Bcd  //bcd
	ReceiveBondData []BondData
}

type BondData struct {
	StockCode [6]types.Ascii //ascii

	StartPrice [3]types.Bcd //bcd
	MaxDealPrice [3]types.Bcd
	MinDealPrice [3]types.Bcd
	NearDealPrice [3]types.Bcd
	AllDealNumber [4]types.Bcd
	DataTime     [6]types.Bcd
}