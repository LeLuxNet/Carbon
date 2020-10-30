package typing

import "strconv"

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
	} else {
		switch other := other.(type) {
		case Bool:
			return Int{other.ToInt() - o.ToInt()}
		case Double:
			return Double{other.Value - float64(o.ToInt())}
		case Int:
			return Int{other.Value - o.ToInt()}
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
	} else {
		switch other := other.(type) {
		case Bool:
			return Int{other.ToInt() / o.ToInt()}
		case Double:
			return Double{other.Value / float64(o.ToInt())}
		case Int:
			return Int{other.Value / o.ToInt()}
		}
	}
	return nil
}
