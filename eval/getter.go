package eval

import (
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/typing"
)

var _ typing.Getter = Getter{}

type Getter struct {
	Name string
	Stmt ast.Statement
	Env  *env.Env
}

func (o Getter) Call(this typing.Object, file *typing.File) (typing.Object, typing.Throwable) {
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

func (o Getter) ToString() string {
	panic("This should not be called! A getter is not a type")
}

func (o Getter) Class() typing.Class {
	panic("This should not be called! A getter is not a type")
}
