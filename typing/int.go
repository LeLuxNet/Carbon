package typing

import (
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

func (o Int) Add(other Object, _ bool) Object {
	switch other := other.(type) {
	case Bool:
		return Int{o.Value + other.ToInt()}
	case Double:
		return Double{float64(o.Value) + other.Value}
	case Int:
		return Int{o.Value + other.Value}
	}
	return nil
}

func (o Int) Sub(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Int{o.Value - other.ToInt()}
		case Double:
			return Double{float64(o.Value) - other.Value}
		case Int:
			return Int{o.Value - other.Value}
		}
	} else {
		switch other := other.(type) {
		case Bool:
			return Int{other.ToInt() - o.Value}
		case Double:
			return Double{other.Value - float64(o.Value)}
		case Int:
			return Int{other.Value - o.Value}
		}
	}
	return nil
}

func (o Int) Mult(other Object, _ bool) Object {
	switch other := other.(type) {
	case Bool:
		return Int{o.Value * other.ToInt()}
	case Double:
		return Double{float64(o.Value) * other.Value}
	case Int:
		return Int{o.Value * other.Value}
	}
	return nil
}

func (o Int) Div(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			return Double{float64(o.Value) / float64(other.ToInt())}
		case Double:
			return Double{float64(o.Value) / other.Value}
		case Int:
			return Double{float64(o.Value) / float64(other.Value)}
		}
	} else {
		switch other := other.(type) {
		case Bool:
			return Int{other.ToInt() / o.Value}
		case Double:
			return Double{other.Value / float64(o.Value)}
		case Int:
			return Int{other.Value / o.Value}
		}
	}
	return nil
}

func (o Int) Pow(other Object, first bool) Object {
	if first {
		switch other := other.(type) {
		case Bool:
			if other.Value {
				return o
			} else {
				return Int{1}
			}
		case Double:
			return Double{math.Pow(float64(o.Value), other.Value)}
		case Int:
			return Double{math.Pow(float64(o.Value), float64(other.Value))}
		}
	} else {
		switch other := other.(type) {
		case Bool:
			if other.Value || o.Value == 0 {
				return Int{1}
			} else {
				return Int{0}
			}
		case Double:
			return Double{math.Pow(other.Value, float64(o.Value))}
		case Int:
			return Double{math.Pow(float64(other.Value), float64(o.Value))}
		}
	}
	return nil
}

func (o Int) Neg() Object {
	return Int{-o.Value}
}
