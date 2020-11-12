package eval

import (
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/typing"
)

var _ typing.Object = Function{}
var _ ast.Callable = Function{}

type Function struct {
	Name  string
	PData ast.ParamData
	Stmt  ast.Statement
	Env   *env.Env
}

func (o Function) Data() ast.ParamData {
	return o.PData
}

func (o Function) Call(args []typing.Object) typing.Throwable {
	e := env.NewEnclosedEnv(o.Env)

	for i, param := range o.PData.Params {
		e.Define(param.Name, args[i], &param.Type, false, false)
	}

	_, err := EvalStmt(o.Stmt, e)
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
