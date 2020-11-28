package type_a

import "github.com/leluxnet/carbon/typing"

var _ typing.Type = Tuple{}

type Tuple struct {
	Values []typing.Type
}

func (o Tuple) ToString() string {
	return "type_a<tuple>"
}

func (o Tuple) Class() typing.Class {
	return typing.TypeClass
}

func (o Tuple) Allows(other typing.Object) bool {
	if other, ok := other.(typing.Tuple); ok {
		if len(o.Values) != len(other.Values) {
			return false
		}

		for i, val := range o.Values {
			if !val.Allows(other.Values[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func (o Tuple) TName() string {
	return "tuple"
}
