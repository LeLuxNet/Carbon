package builtin

import (
	"github.com/leluxnet/carbon/env"
)

func Register(e *env.Env) {
	e.Define("print", Print{}, nil, false, false)
	e.Define("input", Input{}, nil, false, false)
}
