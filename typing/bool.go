package typing

import "strconv"

var _ Object = Bool{}

type Bool struct {
	Value bool
}

func (o Bool) ToString() string {
	return strconv.FormatBool(o.Value)
}

func (o Bool) Class() Class {
	return Class{"bool"}
}

func (o Bool) ToInt() Int {
	var i int
	if o.Value {
		i = 1
	} else {
		i = 0
	}
	return Int{i}
}

func (o Bool) Neg() Object {
	return o.ToInt().Neg()
}
