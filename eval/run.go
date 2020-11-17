package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/builtin"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/lexer"
	"github.com/leluxnet/carbon/parser"
	"github.com/leluxnet/carbon/typing"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func BuiltinEnv() *env.Env {
	e := env.NewEnv()
	builtin.Register(e)

	e.Define(ImportFun.Name, ImportFun, nil, false, false)

	return e
}

func RunFile(name string, e *env.Env) (int, *typing.File) {
	dat, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return 1, nil
	}

	return Run(string(dat), e, false, name, filepath.Dir(name))
}

func Run(source string, e *env.Env, printRes bool, fileName string, path string) (int, *typing.File) {
	source = strings.ReplaceAll(source, "\r", "")

	lex := lexer.Lexer{Source: source}
	token, errs := lex.ScanTokens()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err.ToString(source))
		}
		return 2, nil
	}
	// fmt.Println(token)

	parse := parser.Parser{Tokens: token}
	stmts, errs := parse.Parse()
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Fprintln(os.Stderr, err.ToString(source))
		}
		return 2, nil
	}
	// fmt.Println(stmts)

	file, err2 := Eval(stmts, e, printRes, fileName, path)
	if err2 != nil {
		fmt.Fprintln(os.Stderr, err2.TData().ToString())
		return 1, nil
	}

	return 0, file
}
