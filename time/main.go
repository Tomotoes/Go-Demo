package main

import (
	"math/big"
	"other/time/profile"
	"time"
)

func BigIntFactorial(x big.Int) *big.Int {
	defer profile.Duration(time.Now(), "IntFactorial")

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(&x, one) {
		y.Mul(y, &x)
	}

	return x.Set(y)
}

func main() {

}
