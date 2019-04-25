package cover

import "github.com/lwl1989/data-cover/types"

func N9toString(n9 []types.N9) string{
    bt:=make([]byte, len(n9))
    for k,v := range n9 {
        bt[k] = byte(v)
    }
    return string(bt[:])
}