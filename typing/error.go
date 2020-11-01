package typing

var _ Object = Error{}

type Error struct {
	Message string
}

func (o Error) ToString() string {
	return o.Message
}

func (o Error) Class() Class {
	return Class{"Error"}
}

type ZeroDivisionError struct{}

func (o ZeroDivisionError) ToString() string {
	return "Division by zero"
}

func (o ZeroDivisionError) Class() Class {
	return Class{"ZeroDivisionError"}
}
