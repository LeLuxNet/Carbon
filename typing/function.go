package typing

import (
	"fmt"
)

var _ Callable = BFunction{}

type BFunction struct {
	Name string
	Dat  ParamData
	Cal  func(this Object, args []Object) Throwable
}

func (o BFunction) Data() ParamData {
	return o.Dat
}

func (o BFunction) ToString() string {
	return fmt.Sprintf("builtin-function<%s>", o.Name)
}

func (o BFunction) Class() Class {
	return Class{Name: fmt.Sprintf("function<%s>", o.Name)}
}

func (o BFunction) Call(this Object, args []Object) Throwable {
	return o.Cal(this, args)
}
