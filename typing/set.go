package typing

import (
	"fmt"
	"strings"
)

var _ Object = Set{}

type Set struct {
	Values map[uint64]Object
}

func NewSet() Set {
	return Set{make(map[uint64]Object)}
}

func (o Set) ToString() string {
	var vals []string
	for _, val := range o.Values {
		vals = append(vals, val.ToString())
	}

	return fmt.Sprintf("{%s}", strings.Join(vals, ", "))
}

func (o Set) Class() Class {
	return Class{"set"}
}

func (o Set) Contains(value Object) (Object, Throwable) {
	if hVal, ok := value.(Hashable); ok {
		_, ok := o.Values[hVal.Hash()]
		return Bool{ok}, nil
	}
	return nil, NewError(fmt.Sprintf("'%s' is not hashable", value.Class().Name))
}

func (o Set) Append(value Object) Object {
	if hkey, ok := value.(Hashable); ok {
		o.Values[hkey.Hash()] = value
		return nil
	}
	return Error{fmt.Sprintf("'%s' is not hashable", value.Class().Name)}
}
