package builtin

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/throw"
	"github.com/leluxnet/carbon/typing"
	"strings"
)

var _ ast.Callable = Print{}

type Print struct{}

func (o Print) Data() ast.ParamData {
	return ast.ParamData{
		Args: "args",
	}
}

func (o Print) ToString() string {
	return "print"
}

func (o Print) Class() typing.Class {
	return typing.Class{Name: "function<print>"}
}

func (o Print) Call(args []typing.Object) throw.Throwable {
	if len(args) == 0 {
		return nil
	}

	var builder strings.Builder

	builder.WriteString(args[0].ToString())
	for _, arg := range args[1:] {
		builder.WriteString(" ")
		builder.WriteString(arg.ToString())
	}

	fmt.Println(builder.String())
	return nil
}
