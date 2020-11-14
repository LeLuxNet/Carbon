package typing

import (
	"github.com/leluxnet/carbon/math"
	"math/big"
	"strings"
)

var _ Object = Char{}

type Char struct {
	Value rune
}

func (o Char) ToString() string {
	return string(o.Value)
}

func (o Char) Class() Class {
	return NewNativeClass("char", Properties{})
}

func (o Char) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(Char); ok {
		return Bool{o.Value == other.Value}, nil
	}
	return nil, nil
}

func (o Char) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(Char); ok {
		return Bool{o.Value != other.Value}, nil
	}
	return nil, nil
}

func (o Char) Add(other Object, first bool) (Object, Object) {
	if first {
		switch other := other.(type) {
		case Char:
			return String{string(o.Value) + string(other.Value)}, nil
		case String:
			return String{string(o.Value) + other.Value}, nil
		}
	}
	return nil, nil
}

func (o Char) Mul(other Object, _ bool) (Object, Object) {
	switch other := other.(type) {
	case Int:
		var b strings.Builder
		for i := new(big.Int).Set(other.Value); i.Sign() > 0; i = i.Sub(i, math.IOne) {
			b.WriteRune(o.Value)
		}
		return String{b.String()}, nil
	}
	return nil, nil
}

func (o Char) Hash() uint64 {
	return uint64(o.Value)
}
