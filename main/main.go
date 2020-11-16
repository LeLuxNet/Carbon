package main

import (
	"github.com/leluxnet/carbon/eval"
	"os"
)

func main() {
	e := eval.BuiltinEnv()

	if len(os.Args) > 1 {

		code, _ := eval.RunFile(os.Args[1], e)
		os.Exit(code)
	} else {
		Repl(e)
	}
}
