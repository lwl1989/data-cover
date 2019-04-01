package data


//header  cate 01  code 06  edition 03
type Receiver6 struct {
	StockCode  [6]byte //Ascii

	MatchTime  [6]byte //bcd

	ProjectRevealFlag [1]byte  //bit map
	UpDownFlag        [1]byte
	StatusFlag		  [1]byte

	AllDealNumber		[4]byte  //bcd

	Quotations []QuotationData
}

type QuotationData struct {
	Price [3]byte
	Number [4]byte
}