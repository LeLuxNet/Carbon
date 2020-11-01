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

func (o Bool) Neg() Object {
	var i int
	if o.Value {
		i = -1
	} else {
		i = 0
	}
	return Int{i}
}

func (o Bool) Add(other Object, _ bool) Object {
	switch other := other.(type) {
	case Bool:
		return Int{o.ToInt() + other.ToInt()}
	case Double:
		return Double{Value: float64(o.ToInt()) + other.Value}
	case Int:
		return Int{Value: o.ToInt() + other.Value}
	}
	return nil
}

func (o Bool) Sub(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Int{o.ToInt() - other.ToInt()}
		case Double:
			return Double{float64(o.ToInt()) - other.Value}
		case Int:
			return Int{o.ToInt() - other.Value}
		}
	}
	return nil
}

func (o Bool) Mult(other Object, _ bool) Object {
	if o.Value {
		switch other := other.(type) {
		case Bool:
		case Double:
		case Int:
		case String:
			return other
		}
	} else {
		switch other.(type) {
		case Bool:
			return Bool{false}
		case Double:
			return Double{0}
		case Int:
			return Int{0}
		case String:
			return String{""}
		}
	}
	return nil
}

func (o Bool) Div(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Int{o.ToInt() / other.ToInt()}
		case Double:
			return Double{float64(o.ToInt()) / other.Value}
		case Int:
			return Int{o.ToInt() / other.Value}
		}
	}
	return nil
}

func (o Bool) Mod(other Object, first bool) (Object, Object) {
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
				return Double{imath.DoubleMod(float64(o.ToInt()), other.Value)}, nil

			}
		case Int:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Int{imath.IntMod(o.ToInt(), other.Value)}, nil
			}
		}
	}
	return nil, nil
}

func (o Bool) Pow(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			if o.Value || !other.Value {
				return Int{1}
			} else {
				return Int{0}
			}
		case Double:
		case Int:
			if o.Value || other.Value != 0 {
				return Int{1}
			} else {
				return Int{0}
			}
		}
	}
	return nil
}
