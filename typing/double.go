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

func (o Double) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value == float64(other.Value)}, nil
	case Double:
		return Bool{o.Value == other.Value}, nil
	case Bool:
		return Bool{other.Value == (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o Double) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Bool{o.Value != float64(other.Value)}, nil
	case Double:
		return Bool{o.Value != other.Value}, nil
	case Bool:
		return Bool{other.Value != (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o Double) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Double{o.Value + float64(other.Value)}, nil
	case Double:
		return Double{o.Value + other.Value}, nil
	case Bool:
		return Double{o.Value + float64(other.ToInt())}, nil
	}
	return nil, nil
}

func (o Double) Sub(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Double{o.Value - float64(other.Value)}, nil
		case Double:
			return Double{o.Value - other.Value}, nil
		case Bool:
			return Double{o.Value - float64(other.ToInt())}, nil
		}
	}
	return nil, nil
}

func (o Double) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return Double{o.Value * float64(other.Value)}, nil
	case Double:
		return Double{o.Value * other.Value}, nil
	case Bool:
		return Double{o.Value * float64(other.ToInt())}, nil
	}
	return nil, nil
}

func (o Double) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			return Double{o.Value / float64(other.Value)}, nil
		case Double:
			return Double{o.Value / other.Value}, nil
		case Bool:
			return Double{o.Value / float64(other.ToInt())}, nil
		}
	}
	return nil, nil
}

func (o Double) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Int:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(o.Value, float64(other.Value))}, nil
			}
		case Double:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return Double{imath.DoubleMod(o.Value, other.Value)}, nil

			}
		case Bool:
			if other.Value {
				return Double{0}, nil
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
			return Double{math.Pow(o.Value, float64(other.Value))}, nil
		case Double:
			return Double{math.Pow(o.Value, other.Value)}, nil
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return Int{1}, nil
			}
		}
	}
	return nil, nil
}

func (o Double) Neg() Object {
	return Double{-o.Value}
}
