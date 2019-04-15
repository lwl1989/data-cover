package data

import (
    "github.com/lwl1989/data-cover/cover"
    "io"
)

const FuturesReceiverHeaderLen = 10
const FuturesReceiverEndLen  = 3
//接受者頭部
type FuturesReceiverHeader struct {
    EscCode [1]byte    //ASCII 27  0001 1011 033 27 0x1B
	MessageLen  	[2]byte		//訊息長度
	BizCate			[1]byte		//業務類別
	TransferCode 	[1]byte    	//傳輸格式代碼
	TransferEdition [1]byte  	//傳輸格式版別
	TransferNumber 	[4]byte   	//傳輸序號
}
//接受者
type FuturesReceiver struct {
	Head 	FuturesReceiverHeader
	Body    FuturesReceiverBody
	End     FuturesReceiverEnd

	rc      io.ReadCloser
}
//接受者尾部
type FuturesReceiverEnd struct {
	End     		[1]byte  //XOR值
	TerminalCode 	[2]byte  //(HEXACODE：0D 0A)
}

type FuturesReceiverBody interface{
    ParseData()


}


func (re *FuturesReceiver) getBody() interface{} {
	return re.Body
}

func (rh *FuturesReceiverHeader) getBodyLen() int {
    return cover.BcdToInt(rh.MessageLen[:])
}

func (rh *FuturesReceiverHeader) getBodyType() string {
    return cover.BcdToString(rh.TransferCode[:])
}


func (re *FuturesReceiver) readBody() {
    l := re.Head.getBodyLen()
    if l > 0 {
        buf := make([]byte, l)
        n,err := re.rc.Read(buf)
        if err != nil {
            //todo:
        }
        if n != l {
            //todo:
        }

        rec := GetFormatWithType(re.Head.getBodyType())
        if rec == nil {
            //todo:
        }

        //todo:
        //rec.ParseData(buf)
        //rec parse data

    }
}


func GetFormatWithType(t string) interface{} {
    switch t {
    case "01":
        return &FuturesReceiver1{}
    case "02":
        return &FuturesReceiver2{}
    case "03":
        return &FuturesReceiver3{}
    case "04":
        return &FuturesReceiver4{}
    case "05":
        return &FuturesReceiver5{}
    case "06":
        return &FuturesReceiver6{}
    case "07":
        return &FuturesReceiver7{}
    case "08":
        return &FuturesReceiver8{}
    case "09":
        return &FuturesReceiver9{}
    case "10":
        return &FuturesReceiver10{}
    //case "11":
    //    return &FuturesReceiver11{}
    case "12":
        return &FuturesReceiver12{}
    case "13":
        return &FuturesReceiver13{}
    case "14":
        return &FuturesReceiver14{}
    case "15":
        return &FuturesReceiver15{}
    //case "16":
        //return &FuturesReceiver1{}
    case "17":
        return &FuturesReceiver17{}
    case "18":
        return &FuturesReceiver18{}
    case "19":
        return &FuturesReceiver19{}
    default:
        return nil
    }
}