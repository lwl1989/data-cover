package data


//header  cate 01  code 03  edition 02
type StockReceiver10 struct {
	ExponentialCode [6]byte //Ascii

	ExponentialTime [3]byte //bcd
	NewExponential  [4]byte
}