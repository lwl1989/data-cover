package data

import "github.com/lwl1989/data-cover/types"

//bcd
//成交信息
//header  cate 01  code 02  edition 02
type StockReceiver2 struct {   //bcd all
	TimeFrame [3]types.Bcd
	AllDealMoney [8]types.Bcd
	AllDealNumber [8]types.Bcd
	AllDealStroke [5]types.Bcd

	FundDealMoney [8]types.Bcd
	FundDealNumber [8]types.Bcd
	FundDealStroke [5]types.Bcd

	StockDealMoney [8]types.Bcd
	StockDealNumber [8]types.Bcd
	StockDealStroke [5]types.Bcd

	SubscriptionMoney [8]types.Bcd
	SubscriptionNumber [8]types.Bcd
	SubscriptionStroke [5]types.Bcd

	SaleMoney [8]types.Bcd
	SaleNumber [8]types.Bcd
	SaleStroke [5]types.Bcd
}

