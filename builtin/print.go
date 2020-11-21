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
	Cal: func(_ typing.Object, params map[string]typing.Object, args []typing.Object, kwArgs map[string]typing.Object, _ *typing.File) typing.Throwable {
		end := "\n"
		if s, ok := kwArgs["end"]; ok {
			end = s.ToString()
		}

		if len(args) == 0 {
			fmt.Print(end)
			return nil
		}

		sep := " "
		if s, ok := kwArgs["sep"]; ok {
			sep = s.ToString()
		}

		var builder strings.Builder

		builder.WriteString(args[0].ToString())
		for _, arg := range args[1:] {
			builder.WriteString(sep)
			builder.WriteString(arg.ToString())
		}
		builder.WriteString(end)

		fmt.Print(builder.String())
		return nil
	},
}
