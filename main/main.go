package main

import (
	"fmt"
	"github.com/leluxnet/carbon/eval"
	"os"
)

func main() {
	eval.InitImportFun()
	e := eval.BuiltinEnv()

	if len(os.Args) > 1 {
		name, err := eval.AbsPath(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		code, _ := eval.RunFile(name, e)
		os.Exit(code)
	} else {
		Repl(e)
	}
}
