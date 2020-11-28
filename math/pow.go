package math

import (
	"math"
	"math/big"
)

func Pow(a *big.Float, b *big.Float) *big.Float {
	if b.Sign() == 0 {
		return DOne
	}

	logA := log(a.Abs(a))
	return exp(logA.Mul(b, logA))
}

// TODO: Use big
func log(a *big.Float) *big.Float {
	val, _ := a.Float64()
	return big.NewFloat(math.Log(val))
}

// TODO: Use big
func exp(a *big.Float) *big.Float {
	val, _ := a.Float64()
	return big.NewFloat(math.Exp(val))
}
