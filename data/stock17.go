package data

import "github.com/lwl1989/data-cover/types"

//like 06
//header  cate 01  code 17  edition 03
type StockReceiver17 struct {
	StockCode [6]types.Ascii  //ASCII

	MatchTime [6]types.Bcd //bcd

	ProjectRevealFlag [1]types.BitMap  //bit map
	UpDownFlag        [1]types.BitMap
	StatusFlag		  [1]types.BitMap

	AllDealNumber		[4]types.Bcd  //bcd

	Quotations []QuotationData
}