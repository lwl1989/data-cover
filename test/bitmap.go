package test

import (
	"github.com/lwl1989/data-cover/cover"
	"testing"
)

func Test_BitMapValue(t *testing.T ) {
	var b  [1]byte
	b[0] = 3
	cover.BitMapGetValue(b, 1, 3)
}
