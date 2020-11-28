package math

import (
	"math"
	"math/big"
)

func Atan2(y *big.Float, x *big.Float) *big.Float {
	a, _ := y.Float64()
	b, _ := x.Float64()
	return big.NewFloat(math.Atan2(a, b))
}
