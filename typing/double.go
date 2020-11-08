package typing

import (
	"github.com/leluxnet/carbon/math"
	"math/big"
)

var _ Object = Double{}


type Double struct {
	Value *big.Float
}

func (o Double) ToString() string {
	if o.Value.IsInf() {
		if o.Value.Signbit() {
			return "-Infinity"
		} else {
			return "Infinity"
		}
	} else if o.Value.IsInt() {
		return o.Value.Text('f', 1)
	}
	return o.Value.Text('f', -1)
}

func (o Double) Class() Class {
	return Class{"double"}
}

func (o Double) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value.Cmp(new(big.Float).SetInt(other.Value)) == 0}, nil
	case Double:
		return Bool{o.Value.Cmp(other.Value) == 0}, nil
	case Bool:
		if other.Value {
			return Bool{o.Value.Cmp(math.DOne) == 0}, nil
		} else {
			return Bool{o.Value.Sign() == 0}, nil
		}
	}
	return nil, nil
}

func (o Double) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value.Cmp(new(big.Float).SetInt(other.Value)) != 0}, nil
	case Double:
		return Bool{o.Value.Cmp(other.Value) != 0}, nil
	case Bool:
		if other.Value {
			return Bool{o.Value.Cmp(math.DOne) != 0}, nil
		} else {
			return Bool{o.Value.Sign() != 0}, nil
		}
	}
	return nil, nil
}

func (o Double) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Double{new(big.Float).Add(o.Value, new(big.Float).SetInt(other.Value))}, nil
	case Double:
		return Double{new(big.Float).Add(o.Value, other.Value)}, nil
	case Bool:
		if other.Value {
			return Double{new(big.Float).Add(o.Value, math.DOne)}, nil
		} else {
			return o, nil
		}
	}
	return nil, nil
}

func (o Double) Sub(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Double{new(big.Float).Sub(o.Value, new(big.Float).SetInt(other.Value))}, nil
	case Double:
		return Double{new(big.Float).Sub(o.Value, other.Value)}, nil
	case Bool:
		if other.Value {
			return Double{new(big.Float).Sub(o.Value, big.NewFloat(1))}, nil
		} else {
			return o, nil
		}
	}
	return nil, nil
}

func (o Double) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Double{new(big.Float).Mul(o.Value, new(big.Float).SetInt(other.Value))}, nil
	case Double:
		return Double{new(big.Float).Mul(o.Value, other.Value)}, nil
	case Bool:
		if other.Value {
			return o, nil
		} else {
			return Double{math.DZero}, nil
		}
	}
	return nil, nil
}

func (o Double) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Double{new(big.Float).Quo(o.Value, new(big.Float).SetInt(other.Value))}, nil
		case Double:
			return Double{new(big.Float).Quo(o.Value, other.Value)}, nil
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

func (o Double) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{math.DMod(o.Value, new(big.Float).SetInt(other.Value))}, nil
			}
		case Double:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{math.DMod(o.Value, other.Value)}, nil
			}
		case Bool:
			if other.Value {
				return Int{math.IZero}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		}
	}
	return nil, nil
}

func (o Double) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Double{math.Pow(o.Value, new(big.Float).SetInt(other.Value))}, nil
		case Double:
			return Double{math.Pow(o.Value, other.Value)}, nil
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return Int{math.IOne}, nil
			}
		}
	}
	return nil, nil
}

func (o Double) Neg() Object {
	return Double{new(big.Float).Neg(o.Value)}
}
