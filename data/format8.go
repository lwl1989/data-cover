package data

//header  cate 01 code 08  edition 01
type Receiver8 struct {
	EntrustTime [3]byte

	EntrustBuyStroke [4]byte
	EntrustSellStroke [4]byte
	EntrustBuyNumber [4]byte
	EnTrustSellNumber [4]byte

	FundEntrustBuyStroke [4]byte
	FundEntrustSellStroke [4]byte
	FundEntrustBuyNumber [4]byte
	FundEntrustSellNumber [4]byte

	AllUpEntrustBuyStroke [4]byte
	AllUpEntrustSellStroke [4]byte
	AllUpEntrustBuyNumber [4]byte
	AllUpEnTrustSellNumber [4]byte

	AllDownEntrustBuyStroke [4]byte
	AllDownEntrustSellStroke [4]byte
	AllDownEntrustBuyNumber [4]byte
	AllDownEnTrustSellNumber [4]byte

	FundUpEntrustBuyStroke [4]byte
	FundUpEntrustSellStroke [4]byte
	FundUpEntrustBuyNumber [4]byte
	FundUpEntrustSellNumber [4]byte

	FundDownEntrustBuyStroke [4]byte
	FundDownEntrustSellStroke [4]byte
	FundDownEntrustBuyNumber [4]byte
	FundDownEntrustSellNumber [4]byte
}