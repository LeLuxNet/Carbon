package builtin

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/typing"
)

var _ ast.Callable = BFunction{}

type BFunction struct {
	Name string
	data ast.ParamData
	call func(args []typing.Object) typing.Throwable
}

func (o BFunction) Data() ast.ParamData {
	return o.data
}

func (o BFunction) ToString() string {
	return fmt.Sprintf("builtin-function<%s>", o.Name)
}

func (o BFunction) Class() typing.Class {
	return typing.Class{Name: fmt.Sprintf("function<%s>", o.Name)}
}

func (o BFunction) Call(args []typing.Object) typing.Throwable {
	return o.call(args)
}
