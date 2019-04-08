package test



import (
    "testing"
    "net"
    "fmt"
)

func TestSocket(t *testing.T ) {
    conn, err := net.Dial("tcp", ":8899")
    if err != nil {
        fmt.Printf(" dial error: %s",  err)
    }
    for {
        var b =  make([]byte,1024)
        n,err := conn.Read(b)
        if err != nil {
            fmt.Println(err)
        }
        if n == 0 {
            continue
        }

        fmt.Println(string(b))
    }

}

func HandleConn(conn net.Conn) {

}
