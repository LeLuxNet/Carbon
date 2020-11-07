package typing

import (
	"math/big"
	"strings"
)

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

func (o String) Eq(other Object) (Object, Object) {
	if other, ok := other.(String); ok {
		return Bool{o.Value == other.Value}, nil
	}
	return nil, nil
}

func (o String) NEq(other Object) (Object, Object) {
	if other, ok := other.(String); ok {
		return Bool{o.Value != other.Value}, nil
	}
	return nil, nil
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
		var b strings.Builder
		for i := new(big.Int).Set(other.Value); i.Sign() > 0; i = i.Sub(i, IOne) {
			b.WriteString(o.Value)
		}
		return String{b.String()}, nil
	}
	return nil, nil
}
