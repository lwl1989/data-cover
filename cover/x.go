package cover

import "github.com/lwl1989/data-cover/types"

func XtoString(x []types.X) string{
    bt:=make([]byte, len(x))
    for k,v := range x {
        bt[k] = byte(v)
    }
    return string(bt[:])
}