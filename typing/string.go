package typing

import "strings"

var _ Object = String{}

type String struct {
	Value string
}

func (o String) ToString() string {
	return o.Value
}

func (o String) Class() Class {
	return Class{"string"}
}

func (o String) Add(other Object, first bool) (Object, Object) {
	if first {
		return String{o.Value + other.ToString()}, nil
	} else {
		return String{other.ToString() + o.Value}, nil
	}
}

func (o String) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		return String{strings.Repeat(o.Value, other.Value)}, nil
	}
	return nil, nil
}
