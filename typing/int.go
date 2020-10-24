package typing

import "strconv"

var _ Object = Int{}

type Int struct {
	Value int
}

func (o Int) ToString() string {
	return strconv.Itoa(o.Value)
}

func (o Int) ToDouble() Double {
	return Double{float64(o.Value)}
}

func (o Int) Class() Class {
	return Class{"int"}
}

func (o Int) Neg() Object {
	return Int{-o.Value}
}
