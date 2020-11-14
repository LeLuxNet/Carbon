package typing

import "fmt"

var _ Object = Error{}

type Error struct {
	Message string
}

func (o Error) ToString() string {
	return o.Message
}

func (o Error) Class() Class {
	return NewNativeClass("Error", Properties{})
}

type ZeroDivisionError struct{}

func (o ZeroDivisionError) ToString() string {
	return "Division by zero"
}

func (o ZeroDivisionError) Class() Class {
	return NewNativeClass("ZeroDivisionError", Properties{})
}

type IndexOutOfBoundsError struct {
	Index int64
	Len   int
}

func (o IndexOutOfBoundsError) ToString() string {
	return fmt.Sprintf("Index out of bounds: index %d, len %d", o.Index, o.Len)
}

func (o IndexOutOfBoundsError) Class() Class {
	return NewNativeClass("IndexOutOfBoundsError", Properties{})
}
