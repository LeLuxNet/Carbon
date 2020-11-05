package typing

import (
	imath "github.com/leluxnet/carbon/math"
	"math"
	"strconv"
)

var _ Object = Int{}

type Int struct {
	Value int
}

func (o Int) ToString() string {
	return strconv.Itoa(o.Value)
}

func (o Int) Class() Class {
	return Class{"int"}
}

func (o Int) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value == other.Value}, nil
	case Double:
		return Bool{float64(o.Value) == other.Value}, nil
	case Bool:
		if other.Value {
			return Bool{o.Value == 1}, nil
		} else {
			return Bool{o.Value == 0}, nil
		}
	}
	return nil, nil
}

func (o Int) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Bool:
		return Int{o.Value + other.ToInt()}, nil
	case Double:
		return Double{float64(o.Value) + other.Value}, nil
	case Int:
		return Int{o.Value + other.Value}, nil
	}
	return nil, nil
}

func (o Int) Sub(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Bool:
			return Int{o.Value - other.ToInt()}, nil
		case Double:
			return Double{float64(o.Value) - other.Value}, nil
		case Int:
			return Int{o.Value - other.Value}, nil
		}
	}
	return nil, nil
}

func (o Int) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Bool:
		return Int{o.Value * other.ToInt()}, nil
	case Double:
		return Double{float64(o.Value) * other.Value}, nil
	case Int:
		return Int{o.Value * other.Value}, nil
	}
	return nil, nil
}

func (o Int) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Bool:
			return Double{float64(o.Value) / float64(other.ToInt())}, nil
		case Double:
			return Double{float64(o.Value) / other.Value}, nil
		case Int:
			return Double{float64(o.Value) / float64(other.Value)}, nil
		}
	}
	return nil, nil
}

func (o Int) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Bool:
			if other.Value {
				return Int{0}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		case Double:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(float64(o.Value), other.Value)}, nil

			}
		case Int:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Int{imath.IntMod(o.Value, other.Value)}, nil
			}
		}
	}
	return nil, nil
}

func (o Int) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return Int{1}, nil
			}
		case Double:
			return Double{math.Pow(float64(o.Value), other.Value)}, nil
		case Int:
			return Double{math.Pow(float64(o.Value), float64(other.Value))}, nil
		}
	}
	return nil, nil
}

func (o Int) Neg() Object {
	return Int{-o.Value}
}
