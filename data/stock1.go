package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 01  edition 07
type StockReceiver1 struct {
	ContentBody
	WarrantData
	ContentOther
}


type ContentBody struct {   //47 byte
	Content [47]types.Ascii    //ASCII
	StockCode [6]types.Ascii
	StockName [16]types.Ascii
	StockCate [2]types.Ascii
	StockBondCate [2]types.Ascii
	StockNumFlag [2]types.Ascii

	StockExCode [1]types.Bcd //bcd
	TodayPrice  [3]types.Bcd
	UpPrice     [3]types.Bcd
	DownPrice   [3]types.Bcd

	NotTenFlag [1]types.Ascii  //ASCII
	ExFlag  [1]types.Ascii
	SpecFlag [1]types.Ascii
	SpecExFlag [1]types.Ascii
	NotFlatFinancingStock [1]types.Ascii
	NotFlatBorrowStock [1]types.Ascii
}

type WarrantData struct {
	WarrantCode [1]types.Ascii  //ASCII

	PerformancePrice [4]types.Bcd  //bcd
	YesterdayPerformanceNumber [5]types.Bcd
	YesterdayLostNumber [5]types.Bcd
	SurplusNumber [5]types.Bcd  //发行余量
	ExerciseRatio [4]types.Bcd  //行使比率
	MaxPrice	[4]types.Bcd
	MinPrice	[4]types.Bcd
	ExpireDay	[4]types.Bcd
}

type ContentOther struct {  //7
	ForeignStockCode [1]types.Ascii//ascii

	TradingUnit  [3]types.Bcd  //bcd

	TradingCurrencyCode [3]types.Ascii //ascii

}