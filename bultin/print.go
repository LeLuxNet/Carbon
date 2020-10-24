package builtin

import (
	"fmt"
	"github.com/leluxnet/carbon/typing"
	"strings"
)

var _ typing.Callable = Print{}

type Print struct{}

func (o Print) ToString() string {
	return "print"
}

func (o Print) Class() typing.Class {
	return typing.Class{Name: "print"}
}

func (o Print) Call(args []typing.Object) typing.Object {
	if len(args) == 0 {
		return typing.Null{}
	}

	var builder strings.Builder

	builder.WriteString(args[0].ToString())
	for _, arg := range args[1:] {
		builder.WriteString(" ")
		builder.WriteString(arg.ToString())
	}

	fmt.Println(builder.String())
	return typing.Null{}
}
