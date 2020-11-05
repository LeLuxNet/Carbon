package builtin

import (
	"github.com/leluxnet/carbon/env"
)

func Register(e *env.Env) {
	e.Define(Print.Name, Print, nil, false, false)
	e.Define(Input.Name, Input, nil, false, false)
}
