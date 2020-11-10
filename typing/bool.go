package typing

import (
	"github.com/leluxnet/carbon/math"
	"math/big"
	"strconv"
)

var _ Object = Bool{}

type Bool struct {
	Value bool
}

func (o Bool) ToString() string {
	return strconv.FormatBool(o.Value)
}

func (o Bool) Class() Class {
	return Class{"bool"}
}

func (o Bool) ToInt() int {
	var i int
	if o.Value {
		i = 1
	} else {
		i = 0
	}
	return i
}

func (o Bool) Eq(other Object) (Object, Throwable) {
	switch other := other.(type) {
	case Bool:
		return Bool{o.Value == other.Value}, nil
	case Int:
		return Bool{o.Value == (other.Value.Sign() == 0)}, nil
	case Double:
		return Bool{o.Value == (other.Value.Cmp(math.DOne) == 0)}, nil
	}
	return nil, nil
}

func (o Bool) NEq(other Object) (Object, Throwable) {
	switch other := other.(type) {
	case Bool:
		return Bool{o.Value != other.Value}, nil
	case Int:
		return Bool{o.Value != (other.Value.Sign() == 0)}, nil
	case Double:
		return Bool{o.Value != (other.Value.Cmp(math.DOne) == 0)}, nil
	}
	return nil, nil
}

func (o Bool) Neg() Object {
	if o.Value {
		return Int{math.INegOne}
	} else {
		return Int{math.IZero}
	}
}

func (o Bool) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		if o.Value {
			return Int{Value: new(big.Int).Add(other.Value, math.IOne)}, nil
		} else {
			return other, nil
		}
	case Double:
		if o.Value {
			return Double{Value: new(big.Float).Add(other.Value, math.DOne)}, nil
		} else {
			return other, nil
		}
	case Bool:
		return Int{big.NewInt(int64(o.ToInt() + other.ToInt()))}, nil
	}
	return nil, nil
}

func (o Bool) Sub(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		if o.Value {
			return Int{Value: new(big.Int).Sub(other.Value, math.IOne)}, nil
		} else {
			return other, nil
		}
	case Double:
		if o.Value {
			return Double{Value: new(big.Float).Sub(other.Value, math.DOne)}, nil
		} else {
			return other, nil
		}
	case Bool:
		return Int{big.NewInt(int64(o.ToInt() - other.ToInt()))}, nil
	}
	return nil, nil
}

func (o Bool) Mul(other Object, _ bool) (Object, Object) {
	if o.Value {
		switch other := other.(type) {
		case Int:
		case Double:
		case String:
		case Bool:
			return other, nil
		}
	} else {
		switch other.(type) {
		case Int:
			return Int{math.IZero}, nil
		case Double:
			return Double{math.DZero}, nil
		case String:
			return String{""}, nil
		case Bool:
			return Bool{false}, nil
		}
	}
	return nil, nil
}

func (o Bool) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else if o.Value {
				return Int{new(big.Int).Quo(math.IOne, other.Value)}, nil
			} else {
				return Int{math.IZero}, nil
			}
		case Double:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else if o.Value {
				return Double{new(big.Float).Quo(math.DOne, other.Value)}, nil
			} else {
				return Int{math.IZero}, nil
			}
		case Bool:
			if !other.Value {
				return nil, ZeroDivisionError{}
			} else if o.Value {
				return Int{math.IOne}, nil
			} else {
				return Int{math.IZero}, nil
			}
		}
	}
	return nil, nil
}

func (o Bool) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			switch other.Value.Sign() {
			case 1:
				return Int{math.IOne}, nil
			case 0:
				return nil, ZeroDivisionError{}
			case -1:
				return Int{math.INegOne}, nil
			}
		case Double:
			if other.Value.Sign() == 0 {
				return nil, ZeroDivisionError{}
			} else {
				panic("Not implemented")
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

func (o Bool) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
		case Double:
			if o.Value || other.Value.Sign() != 0 {
				return Int{math.IOne}, nil
			} else {
				return Int{math.IZero}, nil
			}
		case Bool:
			if o.Value || !other.Value {
				return Int{math.IOne}, nil
			} else {
				return Int{math.IZero}, nil
			}
		}
	}
	return nil, nil
}

func (o Bool) Hash() uint64 {
	if o.Value {
		return 1
	} else {
		return 0
	}
}