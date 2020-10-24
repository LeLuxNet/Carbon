package typing

var _ Object = Null{}

type Null struct{}

func (o Null) ToString() string {
	return "null"
}

func (o Null) Class() Class {
	return Class{"null"}
}
