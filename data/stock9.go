package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 06  edition 02
type StockReceiver9 struct {
	StockCode  [6]types.Ascii //Ascii

	MatchTime  [3]types.Bcd //bcd
	DealPrice   [3]types.Bcd
	DealNumber  [4]types.Bcd
}




