package main

import (
	"fmt"
	"github.com/lwl1989/data-cover/cover"
    "net"
    "time"
    "github.com/fsnotify/fsnotify"
    "log"
)

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Println(event.Name, uint32(event.Op))
                log.Println("event:", event)
                if event.Op&fsnotify.Write == fsnotify.Write {
                    log.Println("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Add("/Users/limars/Music")
    if err != nil {
        log.Fatal(err)
    }
    <-done
    return
    listen, err := net.Listen("tcp", ":8899")
    if err != nil {
        fmt.Println("listen error: ", err)
        return
    }
    i:=0
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept error:", err)
        }
        for {
            n, err := conn.Write([]byte("hello"))

            fmt.Println(n, err)
            i++
            fmt.Printf("%d: accept a new connection\n", i)
            time.Sleep(5 * time.Second)
        }

    var b  [1]byte
	b[0] = 127  // 01111111
	            // 00111111
	            // 11111000
	            // 00011111
	fmt.Println(cover.BitMapGetValue(b, 1, 3))
}
