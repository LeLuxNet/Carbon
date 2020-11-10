package typing

var _ Object = Null{}

type Null struct{}

func (o Null) ToString() string {
	return "null"
}

func (o Null) Class() Class {
	return Class{"null"}
}

func (o Null) Eq(other Object) (Object, Throwable) {
	// Other objects can not implement eq to null
	_, ok := other.(Null)
	return Bool{ok}, nil
}

func (o Null) NEq(other Object) (Object, Throwable) {
	// Other objects can not implement eq to null
	_, ok := other.(Null)
	return Bool{!ok}, nil
}
