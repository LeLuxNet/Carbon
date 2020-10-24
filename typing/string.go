package typing

var _ Object = String{}

type String struct {
	Value string
}

func (o String) ToString() string {
	return o.Value
}

func (o String) Class() Class {
	return Class{"string"}
}
