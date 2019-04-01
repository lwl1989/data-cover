package data


//header  cate 01  code 13  edition 02
type Receiver13 struct {
	StockCode [6]byte  // Ascii

	MatchTime  [3]byte //bcd

	UpDownFlag [1]byte //bit map

	DealPrice [3]byte //bcd
	DealNumber [6]byte
	BuyPrice   [3]byte
	SellPrice  [3]byte
}