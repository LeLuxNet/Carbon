package type_a

import "github.com/leluxnet/carbon/typing"

var _ typing.Type = Set{}

type Set struct {
	typing.Type
}

func (o Set) ToString() string {
	return "type_a<set>"
}

func (o Set) Class() typing.Class {
	return typing.TypeClass
}

func (o Set) Allows(other typing.Object) bool {
	if other, ok := other.(typing.Set); ok {
		for _, val := range other.Values {
			if !o.Type.Allows(val) {
				return false
			}
		}
		return true
	}
	return false
}

func (o Set) TName() string {
	return "set"
}
