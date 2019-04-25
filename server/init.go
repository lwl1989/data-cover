package server

import "strings"

type fileSystem struct {
   
}

var f *fileSystem
var beforeDish = map[string]string{
        "GoodsBase":"I010",
        "StockSelectRefer":"I120",
        "StockSelectChange":"I130",
        "ConTractBase":"I140",
        "ConTractCafe":"PB4",
        "DynamicPrice":"PB8",
}
//todo:据说这里也改ftp文件咯
var stockRange = map[string]string {
    "GoodsBase":"I010",
    "StockSelectRefer":"I120",
    "StockSelectChange":"I130",
    "ConTractBase":"I140",
    "ConTractCafe":"PB4",
    "DynamicPrice":"PB8",
}

//文件变更了
func (f *fileSystem) Receive(rec string) string {

    fType := ""
    for k,v := range beforeDish{
        if strings.Index(rec, v) == 0 {
            fType = k
            break
        }
    }

    if fType == "" {
        for k,v := range stockRange{
            if strings.Index(rec, v) == 0 {
                fType = k
                break
            }
        }
    }

    return fType
}

func init() {
    f = &fileSystem{

    }
}
