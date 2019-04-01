package data

//header  cate 01  code 01  edition 07
type Receiver1 struct {
	ContentBody
	WarrantData
	ContentOther
}


type ContentBody struct {   //47 byte
	Content [47]byte    //ASCII
	StockCode [6]byte
	StockName [16]byte
	StockCate [2]byte
	StockBondCate [2]byte
	StockNumFlag [2]byte

	StockExCode [1]byte //bcd
	TodayPrice  [3]byte
	UpPrice     [3]byte
	DownPrice   [3]byte

	NotTenFlag [1]byte  //ASCII
	ExFlag  [1]byte
	SpecFlag [1]byte
	SpecExFlag [1]byte
	NotFlatFinancingStock [1]byte
	NotFlatBorrowStock [1]byte
}

type WarrantData struct {
	WarrantCode [1]byte  //ASCII

	PerformancePrice [4]byte  //bcd
	YesterdayPerformanceNumber [5]byte
	YesterdayLostNumber [5]byte
	SurplusNumber [5]byte  //发行余量
	ExerciseRatio [4]byte  //行使比率
	MaxPrice	[4]byte
	MinPrice	[4]byte
	ExpireDay	[4]byte
}

type ContentOther struct {  //7
	ForeignStockCode [1]byte//ascii

	TradingUnit  [3]byte  //bcd

	TradingCurrencyCode [3]byte //ascii

}