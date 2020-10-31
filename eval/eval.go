package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/throw"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

func Eval(stmts []ast.Statement, e *env.Env, printRes bool) throw.Throwable {
	for _, stmt := range stmts {
		val, err := EvalStmt(stmt, e)
		if err != nil {
			return err
		}

		if printRes && val != nil {
			fmt.Println(val.ToString())
		}
	}
	return nil
}

func EvalStmt(stmt ast.Statement, e *env.Env) (typing.Object, throw.Throwable) {
	switch stmt := stmt.(type) {
	case ast.VarStmt:
		return nil, evalVar(stmt, e)
	case ast.ValStmt:
		return nil, evalVal(stmt, e)
	case ast.AssignStmt:
		return nil, evalAssignment(stmt, e)
	case ast.IfStmt:
		return nil, evalIf(stmt, e)
	case ast.WhileStmt:
		return nil, evalWhile(stmt, e)
	case ast.DoWhileStmt:
		return nil, evalDoWhile(stmt, e)
	case ast.FunStmt:
		return nil, evalFun(stmt, e)
	case ast.ReturnStmt:
		return nil, evalReturn(stmt, e)
	case ast.BreakStmt:
		return nil, throw.Break{}
	case ast.ContinueStmt:
		return nil, throw.Continue{}
	case ast.BlockStmt:
		return nil, evalBlock(stmt, e)
	case ast.ExpressionStmt:
		return evalExpression(stmt.Expr, e)
	}
	return nil, nil
}

func evalExpression(expr ast.Expression, e *env.Env) (typing.Object, throw.Throwable) {
	switch expr := expr.(type) {
	case ast.LiteralExpression:
		return expr.Object, nil
	case ast.GroupingExpression:
		return evalExpression(expr.Expr, e)
	case ast.VariableExpression:
		return evalVariable(expr, e)
	case ast.CallExpression:
		return evalCall(expr, e)
	case ast.UnaryExpression:
		return evalUnary(expr, e)
	case ast.BinaryExpression:
		return evalBinary(expr, e)
	}

	return typing.Null{}, nil
}

func evalVar(expr ast.VarStmt, e *env.Env) throw.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Define(expr.Name, val, nil, false, false)
}

func evalVal(expr ast.ValStmt, e *env.Env) throw.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Define(expr.Name, val, nil, false, true)
}

func evalAssignment(expr ast.AssignStmt, e *env.Env) throw.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Set(expr.Name, val)
}

func evalIf(expr ast.IfStmt, e *env.Env) throw.Throwable {
	condition, err := evalExpression(expr.Condition, e)
	if err != nil {
		return err
	}

	if typing.Truthy(condition) {
		_, err := EvalStmt(expr.Then, e)
		return err
	} else if expr.Else != nil {
		_, err := EvalStmt(expr.Else, e)
		return err
	}
	return nil
}

func evalWhile(expr ast.WhileStmt, e *env.Env) throw.Throwable {
	for true {
		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}

		_, err = EvalStmt(expr.Body, e)
		if err != nil {
			if _, ok := err.(throw.Break); ok {
				return nil
			} else if _, ok := err.(throw.Continue); ok {
				continue
			}
			return err
		}
	}
	return nil
}

func evalDoWhile(expr ast.DoWhileStmt, e *env.Env) throw.Throwable {
	for true {
		_, err := EvalStmt(expr.Body, e)
		if err != nil {
			if _, ok := err.(throw.Break); ok {
				return nil
			} else if _, ok := err.(throw.Continue); ok {
				continue
			}
			return err
		}

		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}
	}
	return nil
}

func evalFun(expr ast.FunStmt, e *env.Env) throw.Throwable {
	fun := Function{
		Name:  expr.Name,
		PData: expr.Data,
		Stmt:  expr.Body,
		Env:   e,
	}

	return e.Define(expr.Name, fun, nil, false, true)
}

func evalReturn(expr ast.ReturnStmt, e *env.Env) throw.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}
	return throw.Return{Data: val}
}

func evalBlock(expr ast.BlockStmt, e *env.Env) throw.Throwable {
	scope := env.NewEnclosedEnv(e)
	for _, stmt := range expr.Body {
		_, err := EvalStmt(stmt, scope)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalVariable(expr ast.VariableExpression, e *env.Env) (typing.Object, throw.Throwable) {
	return e.Get(expr.Name)
}

func evalCall(expr ast.CallExpression, e *env.Env) (typing.Object, throw.Throwable) {
	callee, err := evalExpression(expr.Callee, e)
	if err != nil {
		return nil, err
	}

	fun, ok := callee.(ast.Callable)
	if !ok {
		return nil, throw.NewError("You can only call functions")
	}

	var args []typing.Object
	for _, arg := range expr.Arguments {
		object, err := evalExpression(arg, e)
		if err != nil {
			return nil, err
		}

		args = append(args, object)
	}

	data := fun.Data()
	minArgs := 0
	for _, arg := range data.Params {
		if arg.Default == nil {
			minArgs++
		}
	}

	if len(args) < minArgs {
		return nil, throw.NewError("More args needed")
	} else if len(args) > len(data.Params) && data.Args == "" {
		return nil, throw.NewError("Less args needed")
	}

	err = fun.Call(args)
	if ret, ok := err.(throw.Return); ok {
		return ret.Data, nil
	}
	return nil, err
}

func evalUnary(expr ast.UnaryExpression, e *env.Env) (typing.Object, throw.Throwable) {
	right, err := evalExpression(expr.Right, e)
	if err != nil {
		return nil, err
	}

	switch expr.Type {
	case token.Minus:
		if right, ok := right.(typing.Negable); ok {
			return right.Neg(), nil
		}
	}
	return nil, throw.NewError("Not implemented")
}

func evalBinary(expr ast.BinaryExpression, e *env.Env) (typing.Object, throw.Throwable) {
	left, err := evalExpression(expr.Left, e)
	if err != nil {
		return nil, err
	}

	right, err := evalExpression(expr.Right, e)
	if err != nil {
		return nil, err
	}

	switch expr.Type {
	case token.EqualEqual:
		return typing.Bool{Value: typing.Eq(left, right)}, nil
	case token.Plus:
		return typing.Add(left, right), nil
	case token.Minus:
		return typing.Sub(left, right), nil
	case token.Asterisk:
		return typing.Mult(left, right), nil
	case token.Slash:
		return typing.Div(left, right), nil
	}

	return nil, throw.NewError("Not implemented")
}
