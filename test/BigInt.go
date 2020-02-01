package main

import (
	"fmt"
	"math/big"
)
var (
	zero = big.NewInt(0)
	one = big.NewInt(1)
)
func main() {

	n := new(big.Int)
	n = n.SetBit(zero, 1000, 1)
	n = n.Add(n, one)
	fmt.Println(n.Text(10))
}
