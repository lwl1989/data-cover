package cover

import (
    "strconv"
    "github.com/lwl1989/data-cover/types"
)


func BcdToString(b []types.Bcd) string {
    bt:=make([]byte, len(b))
    for k,v := range b {
        bt[k] = byte(v)
    }
    return string(bt[:])
}

func BcdToInt(b []types.Bcd) int  {
	n,err := strconv.Atoi(BcdToString(b))
	if err == nil {
		return n
	}

	return 0
}

