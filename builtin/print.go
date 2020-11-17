package builtin

import (
	"fmt"
	"github.com/leluxnet/carbon/typing"
	"strings"
)

var Print = typing.BFunction{
	Name: "print",
	Dat: typing.ParamData{
		Args: "args",
	},
	Cal: func(_ typing.Object, args []typing.Object, _ *typing.File) typing.Throwable {
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
