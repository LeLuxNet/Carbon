package math

import (
	"math/big"
)

func IMod(a, b *big.Int) *big.Int {
	res := new(big.Int).Rem(a, b)
	if (res.Sign() < 0 && b.Sign() > 0) || (res.Sign() > 0 && b.Sign() < 0) {
		res.Add(res, b)
	}
	return res
}
func DMod(a, b *big.Float) *big.Float {
	res := Rem(a, b)
	if (res.Sign() < 0 && b.Sign() > 0) || (res.Sign() > 0 && b.Sign() < 0) {
		res.Add(res, b)
	}
	return res
}
