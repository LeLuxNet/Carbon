package typing

import (
	"fmt"
	"strings"
)

var _ Object = Map{}

type Pair struct {
	Key   Object
	Value Object
}

type Map struct {
	Items map[uint64]Pair
}

func NewMap() Map {
	return Map{
		make(map[uint64]Pair),
	}
}

func (o Map) ToString() string {
	var pairs []string
	for _, pair := range o.Items {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.ToString(), pair.Value.ToString()))
	}

	return fmt.Sprintf("{ %s }", strings.Join(pairs, ", "))
}

func (o Map) Class() Class {
	return Class{"map"}
}

func (o Map) SetIndex(key, value Object) Object {
	if hkey, ok := key.(Hashable); ok {
		o.Items[hkey.Hash()] = Pair{key, value}
		return nil
	}
	return Error{fmt.Sprintf("'%s' is not hashable", key.Class().Name)}
}

func (o Map) GetIndex(key Object) (Object, Object) {
	if hkey, ok := key.(Hashable); ok {
		if val, ok := o.Items[hkey.Hash()]; ok {
			return val.Value, nil
		} else {
			return nil, Error{fmt.Sprintf("No such key '%s'", key.ToString())}
		}
	}
	return nil, Error{fmt.Sprintf("'%s' is not hashable", key.Class().Name)}
}
