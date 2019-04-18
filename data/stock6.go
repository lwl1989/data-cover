package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 06  edition 03
type StockReceiver6 struct {
	StockCode  [6]types.Ascii //Ascii

	MatchTime  [6]types.Bcd //bcd

	ProjectRevealFlag [1]types.BitMap  //bit map
	UpDownFlag        [1]types.BitMap
	StatusFlag		  [1]types.BitMap

	AllDealNumber		[4]types.Bcd  //bcd

	Quotations []QuotationData
}

type QuotationData struct {
	Price [3]types.Bcd
	Number [4]types.Bcd
}