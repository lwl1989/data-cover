package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 03  edition 02
type StockReceiver10 struct {
	ExponentialCode [6]types.Ascii //Ascii

	ExponentialTime [3]types.Bcd //bcd
	NewExponential  [4]types.Bcd
}