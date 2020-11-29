package eval

import (
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/typing"
)

var _ typing.Getter = Getter{}
var _ typing.Setter = Setter{}

type Getter struct {
	Name string
	Stmt ast.Statement
	Env  *env.Env
}

func (o Getter) Get(this typing.Object, file *typing.File) (typing.Object, typing.Throwable) {
	e := env.NewEnclosedEnv(o.Env)

	if this != nil {
		e.Define("this", this, nil, false, true)
	}

	_, err := evalStmt(o.Stmt, e, file)
	if ret, ok := err.(typing.Return); ok {
		return ret.Data, nil
	} else {
		return nil, err
	}
}

type Setter struct {
	Name  string
	Param typing.Parameter
	Stmt  ast.Statement
	Env   *env.Env
}

func (o Setter) Set(this, val typing.Object, file *typing.File) typing.Throwable {
	e := env.NewEnclosedEnv(o.Env)

	if this != nil {
		e.Define("this", this, nil, false, true)
	}

	e.Define(o.Param.Name, val, o.Param.Type, false, true)

	_, err := evalStmt(o.Stmt, e, file)
	return err
}
