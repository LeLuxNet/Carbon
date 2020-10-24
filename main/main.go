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
	code := runFile("main.car")

	os.Exit(code)
}

func runFile(name string) int {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return run(string(dat))
}

func run(source string) int {
	source = strings.ReplaceAll(source, "\r", "")

	lex := lexer.Lexer{Source: source}
	token, errs := lex.ScanTokens()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2
	}
	fmt.Println(token)

	parse := parser.Parser{Tokens: token}
	stmts, errs := parse.Parse()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err.ToString(source))
		}
		return 2
	}
	fmt.Println(stmts)

	e := env.NewEnv()
	builtin.Register(e)

	err := eval.Eval(stmts, e)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return 0
}
