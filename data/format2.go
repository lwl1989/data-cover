package data

//bcd
//成交信息
//header  cate 01  code 02  edition 02
type Receiver2 struct {   //bcd all
	TimeFrame [3]byte
	AllDealMoney [8]byte
	AllDealNumber [8]byte
	AllDealStroke [5]byte

	FundDealMoney [8]byte
	FundDealNumber [8]byte
	FundDealStroke [5]byte

	StockDealMoney [8]byte
	StockDealNumber [8]byte
	StockDealStroke [5]byte

	SubscriptionMoney [8]byte
	SubscriptionNumber [8]byte
	SubscriptionStroke [5]byte

	SaleMoney [8]byte
	SaleNumber [8]byte
	SaleStroke [5]byte
}

