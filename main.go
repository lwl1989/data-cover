package main

import (
	"fmt"
	"github.com/lwl1989/data-cover/cover"
)

func main() {
	var b  [1]byte
	b[0] = 127  // 01111111
	            // 00111111
	            // 11111000
	            // 00011111
	fmt.Println(cover.BitMapGetValue(b, 1, 3))
}
