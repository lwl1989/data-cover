package data


//header  cate 01  code 12  edition 02
type FuturesReceiver12 struct {
	BondNumber [1]byte  //bcd
	ReceiveBondData []BondData
}

type BondData struct {
	StockCode [6]byte //ascii

	StartPrice [3]byte //bcd
	MaxDealPrice [3]byte
	MinDealPrice [3]byte
	NearDealPrice [3]byte
	AllDealNumber [4]byte
	DataTime     [6]byte
}