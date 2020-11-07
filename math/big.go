package math

import "math/big"

func Exp(a *big.Float, n *big.Int) *big.Float {
	tmp := a
	res := big.NewFloat(1)
	for n.Sign() == 1 {
		temp := new(big.Float)
		if n.Bit(0) == 1 {
			temp.Mul(res, tmp)
			res = temp
		}
		temp = new(big.Float)
		temp.Mul(tmp, tmp)
		tmp = temp
		n.Quo(n, big.NewInt(2))
	}
	return res
}

func Mod(a *big.Float, b *big.Float) *big.Float {
	iQuo, _ := new(big.Float).Quo(a, b).Int(nil)
	quo := new(big.Float).SetInt(iQuo)
	return new(big.Float).Sub(a, quo.Mul(quo, b))
}
