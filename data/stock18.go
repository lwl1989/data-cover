package data

import "github.com/lwl1989/data-cover/types"

//like 12
//header  cate 01  code 18  edition 02
type StockReceiver18 struct {
	BondNumber [1]types.Bcd  //bcd
	ReceiveBondData []BondData
}