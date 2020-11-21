package typing

import (
	"fmt"
)

var _ Callable = BFunction{}

type BFunction struct {
	Name string
	Dat  ParamData
	Cal  func(this Object, params map[string]Object, args []Object, kwArgs map[string]Object, file *File) Throwable
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

func (o BFunction) Call(this Object, params map[string]Object, args []Object, kwArgs map[string]Object, file *File) Throwable {
	return o.Cal(this, params, args, kwArgs, file)
}
