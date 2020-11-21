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

func (o Function) Call(this typing.Object, params map[string]typing.Object, args []typing.Object, _ map[string]typing.Object, file *typing.File) typing.Throwable {
	e := env.NewEnclosedEnv(o.Env)

	if this != nil {
		e.Define("this", this, nil, false, true)
	}

	for _, param := range o.PData.Params {
		val, _ := params[param.Name]
		e.Define(param.Name, val, &param.Type, false, false)
	}

	if o.PData.Args != "" {
		e.Define(o.PData.Args, typing.Array{Values: args}, nil, false, true)
	}

	_, err := evalStmt(o.Stmt, e, file)
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
