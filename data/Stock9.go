package data


//header  cate 01  code 06  edition 02
type FuturesReceiver9 struct {
	StockCode  [6]byte //Ascii

	MatchTime  [3]byte //bcd
	DealPrice   [3]byte
	DealNumber  [4]byte
}




