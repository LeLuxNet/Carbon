package main

import (
	"fmt"
	"github.com/leluxnet/carbon/builtin"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/eval"
	"github.com/leluxnet/carbon/lexer"
	"github.com/leluxnet/carbon/parser"
	"github.com/leluxnet/carbon/typing"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	e := env.NewEnv()
	builtin.Register(e)

	if len(os.Args) > 1 {

		code, _ := runFile(os.Args[1], e)
		os.Exit(code)
	} else {
		Repl(e)
	}
}

func runFile(name string, e *env.Env) (int, map[string]typing.Object) {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return 1, nil
	}

	return run(string(dat), e, false)
}

func run(source string, e *env.Env, printRes bool) (int, map[string]typing.Object) {
	source = strings.ReplaceAll(source, "\r", "")

	lex := lexer.Lexer{Source: source}
	token, errs := lex.ScanTokens()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2, nil
	}
	// fmt.Println(token)

	parse := parser.Parser{Tokens: token}
	stmts, errs := parse.Parse()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2, nil
	}
	// fmt.Println(stmts)

	props, err := eval.Eval(stmts, e, printRes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.TData().ToString())
		return 1, nil
	}

	return 0, props
}
