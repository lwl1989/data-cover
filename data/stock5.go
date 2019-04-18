package data

import "github.com/lwl1989/data-cover/types"

//header  cate 01  code 05  edition 01
type StockReceiver5 struct {
	Category [1]types.Bcd
	ContentMessage []types.Ascii  //> = 60 bytes
}