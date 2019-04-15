package data


//委托交易信息  all bcd
//header  cate 01  code 04  edition 02
type StockReceiver4 struct {
	EntrustTime  [3]byte

	AllEntrustBuyStroke [4]byte
	AllEntrustSellStroke [4]byte
	AllEntrustBuyNumber [4]byte
	AllEntrustSellNumber [4]byte

	FundEntrustBuyStroke [4]byte
	FundEntrustSellStroke [4]byte
	FundEntrustBuyNumber [4]byte
	FundEntrustSellNumber [4]byte

	StockEntrustBuyStroke [4]byte
	StockEntrustSellStroke [4]byte
	StockEntrustBuyNumber [4]byte
	StockEntrustSellNumber [4]byte

	SubscriptionEntrustBuyStroke [4]byte
	SubscriptionEntrustSellStroke [4]byte
	SubscriptionEntrustBuyNumber [4]byte
	SubscriptionEntrustSellNumber [4]byte

	SaleEntrustBuyStroke [4]byte
	SaleEntrustSellStroke [4]byte
	SaleEntrustBuyNumber [4]byte
	SaleEntrustSellNumber [4]byte

	AllEntrustUpBuyStroke [4]byte
	AllEntrustUpSellStroke [4]byte
	AllEntrustUpBuyNumber [4]byte
	AllEntrustUpSellNumber [4]byte

	AllEntrustDownBuyStroke [4]byte
	AllEntrustDownSellStroke [4]byte
	AllEntrustDownBuyNumber [4]byte
	AllEntrustDownSellNumber [4]byte

	FundEntrustUpBuyStroke [4]byte
	FundEntrustUpSellStroke [4]byte
	FundEntrustUpBuyNumber [4]byte
	FundEntrustUpSellNumber [4]byte

	FundEntrustDownBuyStroke [4]byte
	FundEntrustDownSellStroke [4]byte
	FundEntrustDownBuyNumber [4]byte
	FundEntrustDownSellNumber [4]byte

	StockEntrustUpBuyStroke [4]byte
	StockEntrustUpSellStroke [4]byte
	StockEntrustUpBuyNumber [4]byte
	StockEntrustUpSellNumber [4]byte

	StockEntrustDownBuyStroke [4]byte
	StockEntrustDownSellStroke [4]byte
	StockEntrustDownBuyNumber [4]byte
	StockEntrustDownSellNumber [4]byte

	SubscriptionEntrustUpBuyStroke [4]byte
	SubscriptionEntrustUpSellStroke [4]byte
	SubscriptionEntrustUpBuyNumber [4]byte
	SubscriptionEntrustUpSellNumber [4]byte

	SubscriptionEntrustDownBuyStroke [4]byte
	SubscriptionEntrustDownSellStroke [4]byte
	SubscriptionEntrustDownBuyNumber [4]byte
	SubscriptionEntrustDownSellNumber [4]byte

	SaleEntrustUpBuyStroke [4]byte
	SaleEntrustUpSellStroke [4]byte
	SaleEntrustUpBuyNumber [4]byte
	SaleEntrustUpSellNumber [4]byte

	SaleEntrustDownBuyStroke [4]byte
	SaleEntrustDownSellStroke [4]byte
	SaleEntrustDownBuyNumber [4]byte
	SaleEntrustDownSellNumber [4]byte
}