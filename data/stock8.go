package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01 code 08  edition 01
type StockReceiver8 struct {
	EntrustTime [3]types.Ascii

	EntrustBuyStroke [4]types.Bcd
	EntrustSellStroke [4]types.Bcd
	EntrustBuyNumber [4]types.Bcd
	EnTrustSellNumber [4]types.Bcd

	FundEntrustBuyStroke [4]types.Bcd
	FundEntrustSellStroke [4]types.Bcd
	FundEntrustBuyNumber [4]types.Bcd
	FundEntrustSellNumber [4]types.Bcd

	AllUpEntrustBuyStroke [4]types.Bcd
	AllUpEntrustSellStroke [4]types.Bcd
	AllUpEntrustBuyNumber [4]types.Bcd
	AllUpEnTrustSellNumber [4]types.Bcd

	AllDownEntrustBuyStroke [4]types.Bcd
	AllDownEntrustSellStroke [4]types.Bcd
	AllDownEntrustBuyNumber [4]types.Bcd
	AllDownEnTrustSellNumber [4]types.Bcd

	FundUpEntrustBuyStroke [4]types.Bcd
	FundUpEntrustSellStroke [4]types.Bcd
	FundUpEntrustBuyNumber [4]types.Bcd
	FundUpEntrustSellNumber [4]types.Bcd

	FundDownEntrustBuyStroke [4]types.Bcd
	FundDownEntrustSellStroke [4]types.Bcd
	FundDownEntrustBuyNumber [4]types.Bcd
	FundDownEntrustSellNumber [4]types.Bcd
}