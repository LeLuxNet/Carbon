package typing

var ErrorClass = Class{"error"}

var _ Object = Error{}

type Error struct {
	Message string
}

func (o Error) ToString() string {
	return o.Message
}

func (o Error) Class() Class {
	return ErrorClass
}

func (o Error) Error() string {
	return o.Message
}
