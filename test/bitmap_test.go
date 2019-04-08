package test

import (
	"github.com/lwl1989/data-cover/cover"
	"testing"
    "fmt"
)

func TestBitMapValue(t *testing.T ) {
	var b  [1]byte
	b[0] = 3
	fmt.Println(cover.BitMapGetValue(b, 1, 3))
}
