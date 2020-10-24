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

func (o Function) Call(args []typing.Object) typing.Object {
	if len(args) != len(o.PData.Params) {
		return typing.NewError("Wrong arg count")
	}

	e := env.NewEnclosedEnv(o.Env)

	for i, param := range o.PData.Params {
		e.Define(param.Name, args[i], &param.Type, false, false)
	}

	return EvalStmt(o.Stmt, e)
}

func (o Function) ToString() string {
	return "function <" + o.Name + ">"
}

func (o Function) Class() typing.Class {
	return typing.Class{Name: "function"}
}
