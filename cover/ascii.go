package cover

import "github.com/lwl1989/data-cover/types"

func AsciiToString(b []types.Ascii) string {
    bt:=make([]byte, len(b))
    for k,v := range b {
        bt[k] = byte(v)
    }
    return string(bt[:])
}
