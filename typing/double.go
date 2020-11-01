package typing

import (
	imath "github.com/leluxnet/carbon/math"
	"math"
	"strconv"
)

var _ Object = Double{}

type Double struct {
	Value float64
}

func (o Double) ToString() string {
	if math.IsInf(o.Value, 1) {
		return "Infinity"
	} else if math.IsInf(o.Value, -1) {
		return "-Infinity"
	}
	if o.Value == math.Floor(o.Value) {
		return strconv.FormatFloat(o.Value, 'f', 1, 64)
	}
	return strconv.FormatFloat(o.Value, 'f', -1, 64)
}

func (o Double) Class() Class {
	return Class{"double"}
}

func (o Double) Add(other Object, _ bool) Object {
	switch other := other.(type) {
	case Bool:
		return Double{o.Value + float64(other.ToInt())}
	case Double:
		return Double{o.Value + other.Value}
	case Int:
		return Double{o.Value + float64(other.Value)}
	}
	return nil
}

func (o Double) Sub(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Double{o.Value - float64(other.ToInt())}
		case Double:
			return Double{o.Value - other.Value}
		case Int:
			return Double{o.Value - float64(other.Value)}
		}
	}
	return nil
}

func (o Double) Mult(other Object, _ bool) Object {
	switch other := other.(type) {
	case Bool:
		return Double{o.Value * float64(other.ToInt())}
	case Double:
		return Double{o.Value * other.Value}
	case Int:
		return Double{o.Value * float64(other.Value)}
	}
	return nil
}

func (o Double) Div(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Double{o.Value / float64(other.ToInt())}
		case Double:
			return Double{o.Value / other.Value}
		case Int:
			return Double{o.Value / float64(other.Value)}
		}
	}
	return nil
}

func (o Double) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Bool:
			if other.Value {
				return Double{0}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		case Double:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(o.Value, other.Value)}, nil

			}
		case Int:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(o.Value, float64(other.Value))}, nil
			}
		}
	}
	return nil, nil
}

func (o Double) Pow(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			if other.Value {
				return o
			} else {
				return Int{1}
			}
		case Double:
			return Double{math.Pow(o.Value, other.Value)}
		case Int:
			return Double{math.Pow(o.Value, float64(other.Value))}
		}
	}
	return nil
}

func (o Double) Neg() Object {
	return Double{-o.Value}
}
