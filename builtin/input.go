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
	Cal: func(_ typing.Object, params map[string]typing.Object, args []typing.Object, _ map[string]typing.Object, _ *typing.File) typing.Throwable {
		scanner := bufio.NewScanner(os.Stdin)

		if val, ok := params["text"]; ok {
			fmt.Print(val.ToString())
		}

		scanner.Scan()
		text := scanner.Text()

		return typing.Return{Data: typing.String{Value: text}}
	},
}
