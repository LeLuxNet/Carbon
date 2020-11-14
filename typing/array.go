package typing

import (
	"fmt"
	"math/big"
	"strings"
)

var _ Object = Array{}

type Array struct {
	Values []Object
}

func (o Array) ToString() string {
	var vals []string
	for _, val := range o.Values {
		vals = append(vals, val.ToString())
	}

	return fmt.Sprintf("[%s]", strings.Join(vals, ", "))
}

func (o Array) Class() Class {
	return NewNativeClass("array", Properties{})
}

func (o Array) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(Array); ok {
		if len(o.Values) != len(other.Values) {
			return Bool{false}, nil
		}

		for i, val := range o.Values {
			eq, err := Eq(val, other.Values[i])
			if err != nil {
				return nil, err
			}

			if !Truthy(eq) {
				return Bool{false}, nil
			}
		}
		return Bool{true}, nil
	}
	return nil, nil
}

func (o Array) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(Array); ok {
		if len(o.Values) != len(other.Values) {
			return Bool{true}, nil
		}

		for i, val := range o.Values {
			eq, err := Eq(val, other.Values[i])
			if err != nil {
				return nil, err
			}

			if !Truthy(eq) {
				return Bool{true}, nil
			}
		}
		return Bool{false}, nil
	}
	return nil, nil
}

func (o Array) SetIndex(index, value Object) Object {
	switch index := index.(type) {
	case Int:
		i, err := resolveIntIndex(len(o.Values), index)
		if err != nil {
			return err
		}

		o.Values[i] = value
	}
	return nil
}

func (o Array) GetIndex(index Object) (Object, Object) {
	switch index := index.(type) {
	case Int:
		i, err := resolveIntIndex(len(o.Values), index)
		if err != nil {
			return nil, err
		}

		val := o.Values[i]
		return val, nil
	}
	return nil, Error{fmt.Sprintf("'%s' is not of type int", index.ToString())}
}

func (o Array) Contains(value Object) (Object, Throwable) {
	for _, v := range o.Values {
		eq, err := Eq(value, v)
		if err != nil {
			return nil, err
		}

		if Truthy(eq) {
			return Bool{true}, nil
		}
	}

	return Bool{false}, nil
}

func (o Array) Append(value Object) Object {
	o.Values = append(o.Values, value)
	return nil
}

func resolveIntIndex(l int, index Int) (int64, Object) {
	l64 := int64(l)

	var i int64
	if index.Value.Sign() >= 0 {
		i = index.Value.Int64()
	} else {
		le := big.NewInt(l64)

		// with len = 10:
		//   10 + (-1) = 9 = last element
		i = le.Add(le, index.Value).Int64()
	}

	if l64 > i {
		return i, nil
	} else {
		return 0, IndexOutOfBoundsError{Index: i, Len: l}
	}
}
