package cover

import "github.com/lwl1989/data-cover/types"

func CharToString(b []types.Char) string {
    bt:=make([]byte, len(b))
    for k,v := range b {
        bt[k] = byte(v)
    }
    return string(bt[:])
}
