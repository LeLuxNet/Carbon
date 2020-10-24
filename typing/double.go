package typing

import "strconv"

var _ Object = Double{}

type Double struct {
	Value float64
}

func (o Double) ToString() string {
	return strconv.FormatFloat(o.Value, 'f', -1, 64)
}

func (o Double) Class() Class {
	return Class{"double"}
}

func (o Double) Neg() Object {
	return Double{-o.Value}
}
