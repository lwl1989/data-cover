package data

import (
    "github.com/lwl1989/data-cover/cover"
    "io"
)

const StockReceiverHeaderLen = 10
const StockReceiverEndLen  = 3
//接受者頭部
type StockReceiverHeader struct {
    EscCode [1]byte    //ASCII 27  0001 1011 033 27 0x1B
	MessageLen  	[2]byte		//訊息長度
	BizCate			[1]byte		//業務類別
	TransferCode 	[1]byte    	//傳輸格式代碼
	TransferEdition [1]byte  	//傳輸格式版別
	TransferNumber 	[4]byte   	//傳輸序號
}
//接受者
type StockReceiver struct {
	Head 	StockReceiverHeader
	Body    StockReceiverBody
	End     StockReceiverEnd

	rc      io.ReadCloser
}
//接受者尾部
type StockReceiverEnd struct {
	End     		[1]byte  //XOR值
	TerminalCode 	[2]byte  //(HEXACODE：0D 0A)
}

type StockReceiverBody interface{
    ParseData()


}


func (re *StockReceiver) getBody() interface{} {
	return re.Body
}

func (rh *StockReceiverHeader) getBodyLen() int {
    return cover.BcdToInt(rh.MessageLen[:])
}

func (rh *StockReceiverHeader) getBodyType() string {
    return cover.BcdToString(rh.TransferCode[:])
}


func (re *StockReceiver) readBody() {
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
        return &StockReceiver1{}
    case "02":
        return &StockReceiver2{}
    case "03":
        return &StockReceiver3{}
    case "04":
        return &StockReceiver4{}
    case "05":
        return &StockReceiver5{}
    case "06":
        return &StockReceiver6{}
    case "07":
        return &StockReceiver7{}
    case "08":
        return &StockReceiver8{}
    case "09":
        return &StockReceiver9{}
    case "10":
        return &StockReceiver10{}
    //case "11":
    //    return &StockReceiver11{}
    case "12":
        return &StockReceiver12{}
    case "13":
        return &StockReceiver13{}
    case "14":
        return &StockReceiver14{}
    case "15":
        return &StockReceiver15{}
    //case "16":
        //return &StockReceiver1{}
    case "17":
        return &StockReceiver17{}
    case "18":
        return &StockReceiver18{}
    case "19":
        return &StockReceiver19{}
    default:
        return nil
    }
}