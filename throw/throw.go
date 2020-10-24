package throw

import (
	"github.com/leluxnet/carbon/typing"
)

type Throwable interface {
	TData() typing.Object
}

type Throw struct {
	Data typing.Object
}
type Return struct {
	Data typing.Object
}
type Break struct{}
type Continue struct{}

func (t Throw) TData() typing.Object    { return t.Data }
func (t Return) TData() typing.Object   { return t.Data }
func (t Break) TData() typing.Object    { return typing.Error{Message: "'break' outside of loop"} }
func (t Continue) TData() typing.Object { return typing.Error{Message: "'continue' outside of loop"} }

func NewError(msg string) Throw {
	return Throw{Data: typing.Error{Message: msg}}
}
