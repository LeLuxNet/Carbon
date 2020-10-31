package main

import (
	"fmt"
	"github.com/leluxnet/carbon/builtin"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/eval"
	"github.com/leluxnet/carbon/lexer"
	"github.com/leluxnet/carbon/parser"
	"github.com/leluxnet/carbon/throw"
	"github.com/leluxnet/carbon/typing"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	e := env.NewEnv()
	builtin.Register(e)

	if len(os.Args) > 1 {

		code := runFile(os.Args[1], e)
		os.Exit(code)
	} else {
		Repl(e)
	}
}

func runFile(name string, e *env.Env) int {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	code, _ := run(string(dat), e)
	return code
}

func run(source string, e *env.Env) (int, typing.Object) {
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

	err := eval.Eval(stmts, e)
	switch err.(type) {
	case throw.Throw:
		fmt.Fprintln(os.Stderr, err.TData())
	case throw.Return:
		return 0, err.TData()
	}

	return 0, nil
}
