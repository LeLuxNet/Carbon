package builtin

import (
	"bufio"
	"fmt"
	"github.com/leluxnet/carbon/typing"
	"os"
)

var Input = typing.BFunction{
	Name: "input",
	Dat: typing.ParamData{
		Params: []typing.Parameter{{
			Name:    "text",
			Type:    typing.StringClass,
			Default: typing.String{}}},
	},
	Cal: func(_ typing.Object, args []typing.Object) typing.Throwable {
		scanner := bufio.NewScanner(os.Stdin)

		if len(args) > 0 {
			fmt.Print(args[0].ToString())
		}

		scanner.Scan()
		text := scanner.Text()

		return typing.Return{Data: typing.String{Value: text}}
	},
}
