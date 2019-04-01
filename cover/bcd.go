package cover

import "strconv"


func BcdToString(b []byte) string {
	return string(b[:])
}

func BcdToInt(b []byte) int  {
	n,err := strconv.Atoi(BcdToString(b))
	if err == nil {
		return n
	}

	return 0
}
