package main

import (
	"fmt"
	"math/big"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
	two  = big.NewInt(2)
)

var Cache = map[*big.Int]int{zero: 0, one: 1}

func Fusc(n *big.Int) int {
	for k, v := range Cache {
		if n.Cmp(k) == 0 {
			return v
		}
	}

	//for n.And(n, one).Cmp(zero) == 0 {
	//	n = n.Div(n, two)
	//}

	half := n.Div(n, two)
	halfAddOne := half.Add(half, one)

	//for half.And(half, one).Cmp(zero) == 0 {
	//	half = half.Div(half, two)
	//}

	Cache[half] = Fusc(half)
	Cache[halfAddOne] = Fusc(halfAddOne)
	return Cache[half] + Cache[halfAddOne]
}

func main() {
	n := new(big.Int)
	n = n.SetBit(zero, 1000, 1)
	n = n.Add(n, one)
	fmt.Println(Fusc(n))
}