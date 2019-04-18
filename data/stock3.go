package data

import "github.com/lwl1989/data-cover/types"

type ExponentialData struct {
	Content [4]types.Bcd
}

//header  cate 01  code 03  edition 02
type StockReceiver3 struct {
	ExponentialTime [3]types.Bcd
	ExponentialNumber [1]types.Bcd
	Exponential []ExponentialData
}