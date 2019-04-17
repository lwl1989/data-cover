package cover

import (
    "strconv"
    "github.com/lwl1989/data-cover/types"
)


func BcdToString(b []types.Bcd) string {
	return string(b[:])
}

func BcdToInt(b []types.Bcd) int  {
	n,err := strconv.Atoi(BcdToString(b))
	if err == nil {
		return n
	}

	return 0
}

