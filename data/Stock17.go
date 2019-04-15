package data

//like 06
//header  cate 01  code 17  edition 03
type FuturesReceiver17 struct {
	StockCode [6]byte  //ASCII

	MatchTime [6]byte //bcd

	ProjectRevealFlag [1]byte  //bit map
	UpDownFlag        [1]byte
	StatusFlag		  [1]byte

	AllDealNumber		[4]byte  //bcd

	Quotations []QuotationData
}