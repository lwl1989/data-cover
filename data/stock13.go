package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 13  edition 02
type StockReceiver13 struct {
	StockCode [6]types.Ascii  // Ascii

	MatchTime  [3]types.Bcd //bcd

	UpDownFlag [1]types.BitMap //bit map

	DealPrice [3]types.Bcd //bcd
	DealNumber [6]types.Bcd
	BuyPrice   [3]types.Bcd
	SellPrice  [3]types.Bcd
}