package env

import "github.com/leluxnet/carbon/typing"

type Variable struct {
	Type     *typing.Class
	Value    typing.Object
	Nullable bool
	Constant bool
}
