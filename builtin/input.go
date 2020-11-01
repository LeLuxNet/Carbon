package builtin

import (
	"bufio"
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/typing"
	"os"
)

var _ ast.Callable = Input{}

type Input struct{}

func (o Input) Data() ast.ParamData {
	return ast.ParamData{
		Params: []ast.Parameter{{
			Name:    "text",
			Type:    typing.String{}.Class(),
			Default: typing.String{}}},
	}
}

func (o Input) ToString() string {
	return "input"
}

func (o Input) Class() typing.Class {
	return typing.Class{Name: "function<input>"}
}

func (o Input) Call(args []typing.Object) typing.Throwable {
	scanner := bufio.NewScanner(os.Stdin)

	if len(args) > 0 {
		fmt.Print(args[0].(typing.String).Value)
	}

	scanner.Scan()
	text := scanner.Text()

	return typing.Return{Data: typing.String{Value: text}}
}
