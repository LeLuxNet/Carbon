package typing

import (
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
	} else {
		switch other := other.(type) {
		case Bool:
			return Double{float64(other.ToInt()) - o.Value}
		case Double:
			return Double{other.Value - o.Value}
		case Int:
			return Double{float64(other.Value) - o.Value}
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
	} else {
		switch other := other.(type) {
		case Bool:
			return Double{float64(other.ToInt()) / o.Value}
		case Double:
			return Double{other.Value / o.Value}
		case Int:
			return Double{float64(other.Value) / o.Value}
		}
	}
	return nil
}

func (o Double) Neg() Object {
	return Double{-o.Value}
}
