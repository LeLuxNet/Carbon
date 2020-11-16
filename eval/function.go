package eval

import (
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/typing"
)

var _ typing.Object = Function{}
var _ typing.Callable = Function{}

type Function struct {
	Name  string
	PData typing.ParamData
	Stmt  ast.Statement
	Env   *env.Env
}

func (o Function) Data() typing.ParamData {
	return o.PData
}

func (o Function) Call(this typing.Object, args []typing.Object) typing.Throwable {
	e := env.NewEnclosedEnv(o.Env)

	e.Define("this", this, nil, false, true)

	for i, param := range o.PData.Params {
		e.Define(param.Name, args[i], &param.Type, false, false)
	}

	_, err := evalStmt(o.Stmt, e, map[string]typing.Object{})
	return err
}

func (o Function) ToString() string {
	if o.Name == "" {
		return "lambda function"
	} else {
		return "function<" + o.Name + ">"
	}
}

func (o Function) Class() typing.Class {
	return typing.Class{Name: "function"}
}
