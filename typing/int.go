package typing

import (
	"github.com/leluxnet/carbon/math"
	"math/big"
)

var _ Object = Int{}

var IZero = big.NewInt(0)
var IOne = big.NewInt(1)
var INegOne = big.NewInt(-1)

type Int struct {
	Value *big.Int
}

func (o Int) ToString() string {
	return o.Value.String()
}

func (o Int) Class() Class {
	return Class{"int"}
}

func (o Int) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value.Cmp(o.Value) == 0}, nil
	case Double:
		return Bool{other.Value.Cmp(new(big.Float).SetInt(o.Value)) == 0}, nil
	case Bool:
		if other.Value {
			return Bool{o.Value.Cmp(IOne) == 0}, nil
		} else {
			return Bool{o.Value.Sign() == 0}, nil
		}
	}
	return nil, nil
}

func (o Int) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value.Cmp(o.Value) != 0}, nil
	case Double:
		return Bool{other.Value.Cmp(new(big.Float).SetInt(o.Value)) != 0}, nil
	case Bool:
		if other.Value {
			return Bool{o.Value.Cmp(IOne) != 0}, nil
		} else {
			return Bool{o.Value.Sign() != 0}, nil
		}
	}
	return nil, nil
}

func (o Int) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Int{new(big.Int).Add(o.Value, other.Value)}, nil
	case Double:
		return Double{new(big.Float).Add(new(big.Float).SetInt(o.Value), other.Value)}, nil
	case Bool:
		if other.Value {
			return Int{new(big.Int).Add(o.Value, IOne)}, nil
		} else {
			return o, nil
		}
	}
	return nil, nil
}

func (o Int) Sub(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Int{new(big.Int).Sub(o.Value, other.Value)}, nil
	case Double:
		return Double{new(big.Float).Sub(new(big.Float).SetInt(o.Value), other.Value)}, nil
	case Bool:
		if other.Value {
			return Int{new(big.Int).Sub(o.Value, IOne)}, nil
		} else {
			return o, nil
		}
	}
	return nil, nil
}

func (o Int) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Int{new(big.Int).Mul(o.Value, other.Value)}, nil
	case Double:
		return Double{new(big.Float).Mul(new(big.Float).SetInt(o.Value), other.Value)}, nil
	case Bool:
		if other.Value {
			return o, nil
		} else {
			return Int{Value: IZero}, nil
		}
	}
	return nil, nil
}

func (o Int) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Double{new(big.Float).Quo(new(big.Float).SetInt(o.Value), new(big.Float).SetInt(other.Value))}, nil
		case Double:
			return Double{new(big.Float).Quo(new(big.Float).SetInt(o.Value), other.Value)}, nil
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		}
	}
	return nil, nil
}

func (o Int) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Int{new(big.Int).Mod(o.Value, other.Value)}, nil
			}
		case Double:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{math.Mod(new(big.Float).SetInt(o.Value), other.Value)}, nil
			}
		case Bool:
			if other.Value {
				return Int{IZero}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		}
	}
	return nil, nil
}

func (o Int) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Int{new(big.Int).Exp(o.Value, other.Value, nil)}, nil
		case Double:
			panic("Not implemented")
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return Int{IOne}, nil
			}
		}
	}
	return nil, nil
}

func (o Int) Neg() Object {
	return Int{new(big.Int).Neg(o.Value)}
}

func (o Int) LShift(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Int{new(big.Int).Lsh(o.Value, uint(other.Value.Uint64()))}, nil
		}
	}
	return nil, nil
}

func (o Int) RShift(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Int{new(big.Int).Rsh(o.Value, uint(other.Value.Uint64()))}, nil
		}
	}
	return nil, nil
}
