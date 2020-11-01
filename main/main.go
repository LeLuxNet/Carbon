package main

import (
	"fmt"
	"github.com/leluxnet/carbon/builtin"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/eval"
	"github.com/leluxnet/carbon/lexer"
	"github.com/leluxnet/carbon/parser"
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

	return run(string(dat), e, false)
}

func run(source string, e *env.Env, printRes bool) int {
	source = strings.ReplaceAll(source, "\r", "")

	lex := lexer.Lexer{Source: source}
	token, errs := lex.ScanTokens()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2
	}
	// fmt.Println(token)

	parse := parser.Parser{Tokens: token}
	stmts, errs := parse.Parse()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2
	}
	// fmt.Println(stmts)

	err := eval.Eval(stmts, e, printRes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.TData().ToString())
		return 1
	}

	return 0
}
