package server

import (
    "net"
    "time"
    "github.com/lwl1989/data-cover/data"
    "io"
)

type StockClient struct {
    Ip   string
    Port string
    Conn net.Conn
}


func (stock *StockClient) GetConn() net.Conn{
    if stock.Conn == nil {
        conn,err := net.Dial("tcp", stock.Ip+":"+stock.Port)
        if err != nil {
            //todo:
        }
        //conn.SetDeadline(time.Now().Add(8*time.Second))
        stock.Conn = conn
    }

    return stock.Conn
}


func (stock *StockClient) ReadHeader() {
    buf := make([]byte, data.ReceiverHeaderLen)

    for {
        n, err := stock.Conn.Read(buf)
        if err != nil {
            if err == io.EOF {
                //end reconnection
            }
            //todo:
        }
        if n == 0 {
            //todo:
        }

        //todo: parse header
        time.Sleep(2*time.Second)
    }
}