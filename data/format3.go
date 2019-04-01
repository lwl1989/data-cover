package data


type ExponentialData struct {
	Content [4]byte
}
//header  cate 01  code 03  edition 02
type Receiver3 struct {
	ExponentialTime [3]byte
	ExponentialNumber [1]byte
	Exponential []ExponentialData
}