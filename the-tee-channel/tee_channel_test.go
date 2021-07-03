package theteechannel

import (
	common "ConcurrencyPatternsGolang/utility/common"
	"fmt"
	"testing"
)

/*
out1: 1, out2: 1
out1: 2, out2: 2
out1: 1, out2: 1
out1: 2, out2: 2
*/
func TestTee(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	out1, out2 := Tee(done, common.Take(done, common.Repeat(done, 1, 2), 3))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}
