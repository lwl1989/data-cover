package data

//接受者頭部
type ReceiverHeader struct {
	MessageLen  	[2]byte		//訊息長度
	BizCate			[1]byte		//業務類別
	TransferCode 	[1]byte    	//傳輸格式代碼
	TransferEdition [1]byte  	//傳輸格式版別
	TransferNumber 	[4]byte   	//傳輸序號
}
//接受者
type Receiver struct {
	EscCode [1]byte    //ASCII 27  0001 1011 033 27 0x1B
	Head 	ReceiverHeader
	Body    interface{}
	End     ReceiverEnd
}
//接受者尾部
type ReceiverEnd struct {
	End     		[1]byte  //XOR值
	TerminalCode 	[2]byte  //(HEXACODE：0D 0A)
}

type Receiver1 struct {

}