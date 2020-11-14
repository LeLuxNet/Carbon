package typing

import (
	"github.com/leluxnet/carbon/math"
	"hash/fnv"
	"math/big"
	"strings"
)

var StringClass = NewNativeClass("string", Properties{})

var _ Object = String{}

type String struct {
	Value string
}

func (o String) ToString() string {
	return o.Value
}

func (o String) Class() Class {
	return StringClass
}

func (o String) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(String); ok {
		return Bool{o.Value == other.Value}, nil
	}
	return nil, nil
}

func (o String) NEq(other Object) (Object, Throwable) {
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
		for i := new(big.Int).Set(other.Value); i.Sign() > 0; i = i.Sub(i, math.IOne) {
			b.WriteString(o.Value)
		}
		return String{b.String()}, nil
	}
	return nil, nil
}

func (o String) Hash() uint64 {
	return hashString(o.Value)
}

func hashString(s string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return h.Sum64()
}

func (o String) GetIndex(key Object) (Object, Object) {
	switch key := key.(type) {
	case Int:
		chars := []rune(o.Value)

		i, err := resolveIntIndex(len(chars), key)
		if err != nil {
			return nil, err
		}

		return Char{chars[i]}, nil
	}
	return nil, nil
}
