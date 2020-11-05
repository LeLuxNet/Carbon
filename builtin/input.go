package builtin

import (
	"bufio"
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/typing"
	"os"
)

var Input = BFunction{
	Name: "input",
	data: ast.ParamData{
		Params: []ast.Parameter{{
			Name:    "text",
			Type:    typing.String{}.Class(),
			Default: typing.String{}}},
	},
	call: func(args []typing.Object) typing.Throwable {
		scanner := bufio.NewScanner(os.Stdin)

		if len(args) > 0 {
			fmt.Print(args[0].ToString())
		}

		scanner.Scan()
		text := scanner.Text()

		return typing.Return{Data: typing.String{Value: text}}
	},
}
