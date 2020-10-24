package ast

import (
	"github.com/leluxnet/carbon/throw"
	"github.com/leluxnet/carbon/typing"
)

type Callable interface {
	Data() ParamData
	Call(args []typing.Object) throw.Throwable
}

type Parameter struct {
	Name    string
	Type    typing.Class
	Default typing.Object
}

type ParamData struct {
	Params []Parameter
	Args   string
	KwArgs string
}
