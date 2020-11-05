package builtin

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/typing"
	"strings"
)

var Print = BFunction{
	Name: "print",
	data: ast.ParamData{
		Args: "args",
	},
	call: func(args []typing.Object) typing.Throwable {
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
	},
}
