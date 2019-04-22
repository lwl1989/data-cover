package server

import (
    "github.com/fsnotify/fsnotify"
    "log"
)

type FileSystem struct {
    FsNotify fsnotify.Watcher
    FilePath []string
}

func (fsn *FileSystem) AddNotify(path string) {
    found := false

    for _,v := range fsn.FilePath {
        if v == path {
            found = true
            break
        }
    }

    if !found {
        fsn.FilePath = append(fsn.FilePath, path)
    }
}

func (fsn *FileSystem) RemoveNotify(path string) {
    k := -1

    for k1,v := range fsn.FilePath {
        if v == path {
            k = k1
            break
        }
    }

    if k != -1 {
        count := len(fsn.FilePath)
        if k < (count - 1) {
            fsn.FilePath = append(fsn.FilePath[0:k], fsn.FilePath[k:count]...)
        }
    }
}

/**
go func() {
        for {
            select {
            case event, ok := <-.Events:
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
 */
func (fsn *FileSystem) Watch(f func(watcher *fsnotify.Watcher)) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
       log.Fatal("watch fail:",err)
    }
    defer watcher.Close()

    //loop
    f(watcher)
}


func filechange(watcher *fsnotify.Watcher){
    for {
        select {
        case event, ok := <-watcher.Events:
            if !ok {
                //todo: send error
                return
            }

            switch event.Op {
                case fsnotify.Create:
                case fsnotify.Write:
                    fallthrough
                case fsnotify.Remove:
                    f.Receive(event.Name)
                case fsnotify.Rename:
                    // not do any thing
                case fsnotify.Chmod:
                    // not do any thing
                default:
                 //todo: send error
            }
            //Create Op = 1 << iota
            //Write
            //Remove
            //Rename
            //Chmod
            // 001  010  011 100 101
            // 001&001 = 001

            //log.Println("event:", event)
            //if event.Op&fsnotify.Write == fsnotify.Write {
            //    log.Println("modified file:", event.Name)
            //}
        case err, ok := <-watcher.Errors:
            //todo: send error
            if !ok {
                return
            }
            log.Println("error:", err)
        }
    }
}




