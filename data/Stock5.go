package data


//header  cate 01  code 05  edition 01
type StockReceiver5 struct {
	Category [1]byte
	ContentMessage []byte  //> = 60 bytes
}