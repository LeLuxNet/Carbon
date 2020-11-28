package env

import "github.com/leluxnet/carbon/typing"

type Variable struct {
	Type     typing.Type
	Value    typing.Object
	Nullable bool
	Constant bool
}
