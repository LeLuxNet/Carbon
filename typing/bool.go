package typing

import (
	imath "github.com/leluxnet/carbon/math"
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

func (o Bool) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Bool:
		return Bool{o.Value == other.Value}, nil
	case Int:
	case Double:
		return Bool{o.Value == (other.Value == 1)}, nil
	}
	return nil, nil
}

func (o Bool) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Bool:
		return Bool{o.Value != other.Value}, nil
	case Int:
	case Double:
		return Bool{o.Value != (other.Value == 1)}, nil
	}
	return nil, nil
}

func (o Bool) Neg() Object {
	var i int
	if o.Value {
		i = -1
	} else {
		i = 0
	}
	return Int{i}
}

func (o Bool) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Int{Value: o.ToInt() + other.Value}, nil
	case Double:
		return Double{Value: float64(o.ToInt()) + other.Value}, nil
	case Bool:
		return Int{o.ToInt() + other.ToInt()}, nil
	}
	return nil, nil
}

func (o Bool) Sub(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Int{o.ToInt() - other.Value}, nil
		case Double:
			return Double{float64(o.ToInt()) - other.Value}, nil
		case Bool:
			return Int{o.ToInt() - other.ToInt()}, nil
		}
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
			return Int{0}, nil
		case Double:
			return Double{0}, nil
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
			return Int{o.ToInt() / other.Value}, nil
		case Double:
			return Double{float64(o.ToInt()) / other.Value}, nil
		case Bool:
			return Int{o.ToInt() / other.ToInt()}, nil
		}
	}
	return nil, nil
}

func (o Bool) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Int{imath.IntMod(o.ToInt(), other.Value)}, nil
			}
		case Double:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(float64(o.ToInt()), other.Value)}, nil

			}
		case Bool:
			if other.Value {
				return Int{0}, nil
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
			if o.Value || other.Value != 0 {
				return Int{1}, nil
			} else {
				return Int{0}, nil
			}
		case Bool:
			if o.Value || !other.Value {
				return Int{1}, nil
			} else {
				return Int{0}, nil
			}
		}
	}
	return nil, nil
}
