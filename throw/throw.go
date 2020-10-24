package throw

import (
	"github.com/leluxnet/carbon/typing"
)

type Throwable interface {
	throwable()
}

type Throw struct {
	Data typing.Object
}
type Return struct {
	Data typing.Object
}

func (t Throw) throwable()  {}
func (t Return) throwable() {}

func NewError(msg string) Throw {
	return Throw{Data: typing.Error{Message: msg}}
}
