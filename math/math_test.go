package math

import (
	"math/big"
	"testing"
)

func float(str string) *big.Float {
	num, _ := new(big.Float).SetString(str)
	return num
}

var num1 = float("85849042357678784738274893289745872394753245983274985532583480543789149802347892349087218.947985348579823")
var num2 = float("-4328423804823092857934598723985623485763724982738072375098975289937957324978967696785.2727544")

func BenchmarkMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DMod(num1, num2)
	}
}
