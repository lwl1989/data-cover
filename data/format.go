package data

import (
    "github.com/lwl1989/data-cover/cover"
    "io"
)

const ReceiverHeaderLen = 10
const ReceiverEndLen  = 3
//接受者頭部
type ReceiverHeader struct {
    EscCode [1]byte    //ASCII 27  0001 1011 033 27 0x1B
	MessageLen  	[2]byte		//訊息長度
	BizCate			[1]byte		//業務類別
	TransferCode 	[1]byte    	//傳輸格式代碼
	TransferEdition [1]byte  	//傳輸格式版別
	TransferNumber 	[4]byte   	//傳輸序號
}
//接受者
type Receiver struct {
	Head 	ReceiverHeader
	Body    ReceiverBody
	End     ReceiverEnd

	rc      io.ReadCloser
}
//接受者尾部
type ReceiverEnd struct {
	End     		[1]byte  //XOR值
	TerminalCode 	[2]byte  //(HEXACODE：0D 0A)
}

type ReceiverBody interface{
    ParseData()


}


func (re *Receiver) getBody() interface{} {
	return re.Body
}

func (rh *ReceiverHeader) getBodyLen() int {
    return cover.BcdToInt(rh.MessageLen[:])
}

func (rh *ReceiverHeader) getBodyType() string {
    return cover.BcdToString(rh.TransferCode[:])
}


func (re *Receiver) readBody() {
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
        return &Receiver1{}
    case "02":
        return &Receiver2{}
    case "03":
        return &Receiver3{}
    case "04":
        return &Receiver4{}
    case "05":
        return &Receiver5{}
    case "06":
        return &Receiver6{}
    case "07":
        return &Receiver7{}
    case "08":
        return &Receiver8{}
    case "09":
        return &Receiver9{}
    case "10":
        return &Receiver10{}
    //case "11":
    //    return &Receiver11{}
    case "12":
        return &Receiver12{}
    case "13":
        return &Receiver13{}
    case "14":
        return &Receiver14{}
    case "15":
        return &Receiver15{}
    //case "16":
        //return &Receiver1{}
    case "17":
        return &Receiver17{}
    case "18":
        return &Receiver18{}
    case "19":
        return &Receiver19{}
    default:
        return nil
    }
}