package data

import "github.com/lwl1989/data-cover/types"

//委托交易信息  all bcd
//header  cate 01  code 04  edition 02
type StockReceiver4 struct {
	EntrustTime  [3]types.Ascii

	AllEntrustBuyStroke [4]types.Bcd
	AllEntrustSellStroke [4]types.Bcd
	AllEntrustBuyNumber [4]types.Bcd
	AllEntrustSellNumber [4]types.Bcd

	FundEntrustBuyStroke [4]types.Bcd
	FundEntrustSellStroke [4]types.Bcd
	FundEntrustBuyNumber [4]types.Bcd
	FundEntrustSellNumber [4]types.Bcd

	StockEntrustBuyStroke [4]types.Bcd
	StockEntrustSellStroke [4]types.Bcd
	StockEntrustBuyNumber [4]types.Bcd
	StockEntrustSellNumber [4]types.Bcd

	SubscriptionEntrustBuyStroke [4]types.Bcd
	SubscriptionEntrustSellStroke [4]types.Bcd
	SubscriptionEntrustBuyNumber [4]types.Bcd
	SubscriptionEntrustSellNumber [4]types.Bcd

	SaleEntrustBuyStroke [4]types.Bcd
	SaleEntrustSellStroke [4]types.Bcd
	SaleEntrustBuyNumber [4]types.Bcd
	SaleEntrustSellNumber [4]types.Bcd

	AllEntrustUpBuyStroke [4]types.Bcd
	AllEntrustUpSellStroke [4]types.Bcd
	AllEntrustUpBuyNumber [4]types.Bcd
	AllEntrustUpSellNumber [4]types.Bcd

	AllEntrustDownBuyStroke [4]types.Bcd
	AllEntrustDownSellStroke [4]types.Bcd
	AllEntrustDownBuyNumber [4]types.Bcd
	AllEntrustDownSellNumber [4]types.Bcd

	FundEntrustUpBuyStroke [4]types.Bcd
	FundEntrustUpSellStroke [4]types.Bcd
	FundEntrustUpBuyNumber [4]types.Bcd
	FundEntrustUpSellNumber [4]types.Bcd

	FundEntrustDownBuyStroke [4]types.Bcd
	FundEntrustDownSellStroke [4]types.Bcd
	FundEntrustDownBuyNumber [4]types.Bcd
	FundEntrustDownSellNumber [4]types.Bcd

	StockEntrustUpBuyStroke [4]types.Bcd
	StockEntrustUpSellStroke [4]types.Bcd
	StockEntrustUpBuyNumber [4]types.Bcd
	StockEntrustUpSellNumber [4]types.Bcd

	StockEntrustDownBuyStroke [4]types.Bcd
	StockEntrustDownSellStroke [4]types.Bcd
	StockEntrustDownBuyNumber [4]types.Bcd
	StockEntrustDownSellNumber [4]types.Bcd

	SubscriptionEntrustUpBuyStroke [4]types.Bcd
	SubscriptionEntrustUpSellStroke [4]types.Bcd
	SubscriptionEntrustUpBuyNumber [4]types.Bcd
	SubscriptionEntrustUpSellNumber [4]types.Bcd

	SubscriptionEntrustDownBuyStroke [4]types.Bcd
	SubscriptionEntrustDownSellStroke [4]types.Bcd
	SubscriptionEntrustDownBuyNumber [4]types.Bcd
	SubscriptionEntrustDownSellNumber [4]types.Bcd

	SaleEntrustUpBuyStroke [4]types.Bcd
	SaleEntrustUpSellStroke [4]types.Bcd
	SaleEntrustUpBuyNumber [4]types.Bcd
	SaleEntrustUpSellNumber [4]types.Bcd

	SaleEntrustDownBuyStroke [4]types.Bcd
	SaleEntrustDownSellStroke [4]types.Bcd
	SaleEntrustDownBuyNumber [4]types.Bcd
	SaleEntrustDownSellNumber [4]types.Bcd
}