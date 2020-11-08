package math

import "math/big"

func Rem(a, b *big.Float) *big.Float {
	iQuo, _ := new(big.Float).Quo(a, b).Int(nil)
	quo := new(big.Float).SetInt(iQuo)
	return new(big.Float).Sub(a, quo.Mul(quo, b))
}
