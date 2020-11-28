package type_a

import "github.com/leluxnet/carbon/typing"

var _ typing.Type = Map{}

type Map struct {
	Key   typing.Type
	Value typing.Type
}

func (o Map) ToString() string {
	return "type_a<map>"
}

func (o Map) Class() typing.Class {
	return typing.TypeClass
}

func (o Map) Allows(other typing.Object) bool {
	if other, ok := other.(typing.Map); ok {
		for _, item := range other.Items {
			if !o.Key.Allows(item.Key) || !o.Value.Allows(item.Value) {
				return false
			}
		}
		return true
	}
	return false
}

func (o Map) TName() string {
	return "map"
}
