package typing

import (
	imath "github.com/leluxnet/carbon/math"
	"math"
	"strconv"
)

var _ Object = NDouble{}

type NDouble struct {
	Value float64
}

func (o NDouble) ToString() string {
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

func (o NDouble) Class() Class {
	return Class{"native double"}
}

func (o NDouble) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return Bool{o.Value == float64(other.Value)}, nil
	case NDouble:
		return Bool{o.Value == other.Value}, nil
	case Bool:
		return Bool{other.Value == (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o NDouble) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return Bool{o.Value != float64(other.Value)}, nil
	case NDouble:
		return Bool{o.Value != other.Value}, nil
	case Bool:
		return Bool{other.Value != (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o NDouble) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return NDouble{o.Value + float64(other.Value)}, nil
	case NDouble:
		return NDouble{o.Value + other.Value}, nil
	case Bool:
		return NDouble{o.Value + float64(other.ToInt())}, nil
	}
	return nil, nil
}

func (o NDouble) Sub(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NDouble{o.Value - float64(other.Value)}, nil
		case NDouble:
			return NDouble{o.Value - other.Value}, nil
		case Bool:
			return NDouble{o.Value - float64(other.ToInt())}, nil
		}
	}
	return nil, nil
}

func (o NDouble) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return NDouble{o.Value * float64(other.Value)}, nil
	case NDouble:
		return NDouble{o.Value * other.Value}, nil
	case Bool:
		return NDouble{o.Value * float64(other.ToInt())}, nil
	}
	return nil, nil
}

func (o NDouble) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NDouble{o.Value / float64(other.Value)}, nil
		case NDouble:
			return NDouble{o.Value / other.Value}, nil
		case Bool:
			return NDouble{o.Value / float64(other.ToInt())}, nil
		}
	}
	return nil, nil
}

func (o NDouble) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return NDouble{imath.DoubleMod(o.Value, float64(other.Value))}, nil
			}
		case NDouble:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return NDouble{imath.DoubleMod(o.Value, other.Value)}, nil

			}
		case Bool:
			if other.Value {
				return NDouble{0}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		}
	}
	return nil, nil
}

func (o NDouble) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NDouble{math.Pow(o.Value, float64(other.Value))}, nil
		case NDouble:
			return NDouble{math.Pow(o.Value, other.Value)}, nil
		case Bool:
			if other.Value {
				return o, nil
			} else {
				return NInt{1}, nil
			}
		}
	}
	return nil, nil
}

func (o NDouble) Neg() Object {
	return NDouble{-o.Value}
}