package type_a

import "github.com/leluxnet/carbon/typing"

var _ typing.Type = Array{}

type Array struct {
	typing.Type
}

func (o Array) Allows(other typing.Object) bool {
	if other, ok := other.(typing.Array); ok {
		for _, val := range other.Values {
			if !o.Type.Allows(val) {
				return false
			}
		}
		return true
	}
	return false
}

func (o Array) TName() string {
	return "array"
}
