package server

import "github.com/fsnotify/fsnotify"

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