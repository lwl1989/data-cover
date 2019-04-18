package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 14  edition 02
type StockReceiver14 struct {
	SubscriptionCode [6]types.Ascii   //ascii
	SubscriptionName [50]types.Ascii
}