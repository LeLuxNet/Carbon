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

func (o String) Add(other Object, first bool) Object {
	if first {
		return String{o.Value + other.ToString()}
	} else {
		return String{other.ToString() + o.Value}
	}
}

func (o String) Mult(other Object, _ bool) Object {
	switch other := other.(type) {
	case Int:
		return String{strings.Repeat(o.Value, other.Value)}
	}
	return nil
}
