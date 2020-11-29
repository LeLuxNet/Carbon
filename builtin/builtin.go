package builtin

import (
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/typing"
)

func Register(e *env.Env) {
	e.Define("string", typing.StringClass, nil, false, true)
	e.Define("bool", typing.Bool{}.Class(), nil, false, true)
	e.Define("int", typing.IntClass, nil, false, true)
	e.Define("double", typing.DoubleClass, nil, false, true)
	e.Define("char", typing.Char{}.Class(), nil, false, true)

	e.Define(Print.Name, Print, nil, false, false)
	e.Define(Input.Name, Input, nil, false, false)
}
