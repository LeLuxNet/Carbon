package math

import "math"

func IntMod(a, b int) int {
	return (a%b + b) % b
}

func DoubleMod(a, b float64) float64 {
	return math.Mod(math.Mod(a, b)+b, b)
}
