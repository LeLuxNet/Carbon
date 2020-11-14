package typing

import (
	"fmt"
	"strings"
)

var _ Object = Tuple{}

type Tuple struct {
	Values []Object
}

func (o Tuple) ToString() string {
	var vals []string
	for _, val := range o.Values {
		vals = append(vals, val.ToString())
	}

	return fmt.Sprintf("(%s)", strings.Join(vals, ", "))
}

func (o Tuple) Class() Class {
	return NewNativeClass("tuple", Properties{})
}

func (o Tuple) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(Tuple); ok {
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

func (o Tuple) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(Tuple); ok {
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

func (o Tuple) GetIndex(index Object) (Object, Object) {
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

func (o Tuple) Contains(value Object) (Object, Throwable) {
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
