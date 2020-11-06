package typing

import (
	imath "github.com/leluxnet/carbon/math"
	"math"
	"strconv"
)

var _ Object = NInt{}

type NInt struct {
	Value int
}

func (o NInt) ToString() string {
	return strconv.Itoa(o.Value)
}

func (o NInt) Class() Class {
	return Class{"native int"}
}

func (o NInt) Eq(other Object) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return Bool{o.Value == other.Value}, nil
	case NDouble:
		return Bool{float64(o.Value) == other.Value}, nil
	case Bool:
		return Bool{other.Value == (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o NInt) NEq(other Object) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return Bool{o.Value != other.Value}, nil
	case NDouble:
		return Bool{float64(o.Value) != other.Value}, nil
	case Bool:
		return Bool{other.Value != (o.Value == 1)}, nil
	}
	return nil, nil
}

func (o NInt) Add(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return NInt{o.Value + other.Value}, nil
	case NDouble:
		return NDouble{float64(o.Value) + other.Value}, nil
	case Bool:
		return NInt{o.Value + other.ToInt()}, nil
	}
	return nil, nil
}

func (o NInt) Sub(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NInt{o.Value - other.Value}, nil
		case NDouble:
			return NDouble{float64(o.Value) - other.Value}, nil
		case Bool:
			return NInt{o.Value - other.ToInt()}, nil
		}
	}
	return nil, nil
}

func (o NInt) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case NInt:
		return NInt{o.Value * other.Value}, nil
	case NDouble:
		return NDouble{float64(o.Value) * other.Value}, nil
	case Bool:
		return NInt{o.Value * other.ToInt()}, nil
	}
	return nil, nil
}

func (o NInt) Div(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NDouble{float64(o.Value) / float64(other.Value)}, nil
		case NDouble:
			return NDouble{float64(o.Value) / other.Value}, nil
		case Bool:
			return NDouble{float64(o.Value) / float64(other.ToInt())}, nil
		}
	}
	return nil, nil
}

func (o NInt) Mod(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return NInt{imath.IntMod(o.Value, other.Value)}, nil
			}
		case NDouble:
			if other.Value == 0 {
				return nil, ZeroDivisionError{}
			} else {
				return NDouble{imath.DoubleMod(float64(o.Value), other.Value)}, nil
			}
		case Bool:
			if other.Value {
				return NInt{0}, nil
			} else {
				return nil, ZeroDivisionError{}
			}
		}
	}
	return nil, nil
}

func (o NInt) Pow(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NDouble{math.Pow(float64(o.Value), float64(other.Value))}, nil
		case NDouble:
			return NDouble{math.Pow(float64(o.Value), other.Value)}, nil
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

func (o NInt) Neg() Object {
	return NInt{-o.Value}
}

func (o NInt) LShift(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NInt{o.Value << other.Value}, nil
		}
	}
	return nil, nil
}

func (o NInt) RShift(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case NInt:
			return NInt{o.Value >> other.Value}, nil
		}
	}
	return nil, nil
}