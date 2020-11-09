package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

func Eval(stmts []ast.Statement, e *env.Env, printRes bool) typing.Throwable {
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

func EvalStmt(stmt ast.Statement, e *env.Env) (typing.Object, typing.Throwable) {
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
		return nil, typing.Break{}
	case ast.ContinueStmt:
		return nil, typing.Continue{}
	case ast.BlockStmt:
		return nil, evalBlock(stmt, e)
	case ast.ExpressionStmt:
		return evalExpression(stmt.Expr, e)
	}
	return nil, nil
}

func evalExpression(expr ast.Expression, e *env.Env) (typing.Object, typing.Throwable) {
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

func evalVar(expr ast.VarStmt, e *env.Env) typing.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Define(expr.Name, val, nil, false, false)
}

func evalVal(expr ast.ValStmt, e *env.Env) typing.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Define(expr.Name, val, nil, false, true)
}

func evalAssignment(expr ast.AssignStmt, e *env.Env) typing.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	if expr.Type == token.Equal {
		return e.Set(expr.Name, val)
	}

	oldVal, err := e.Get(expr.Name)
	if err != nil {
		return err
	}

	switch expr.Type {
	case token.PlusEqual:
		val, err = typing.Add(oldVal, val)
	case token.MinusEqual:
		val, err = typing.Sub(oldVal, val)
	case token.AsteriskEqual:
		val, err = typing.Mul(oldVal, val)
	case token.SlashEqual:
		val, err = typing.Div(oldVal, val)
	case token.PercentEqual:
		val, err = typing.Mod(oldVal, val)
	case token.AsteriskAsteriskEqual:
		val, err = typing.Pow(oldVal, val)
	}

	if err != nil {
		return err
	}

	return e.Set(expr.Name, val)
}

func evalIf(expr ast.IfStmt, e *env.Env) typing.Throwable {
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

func evalWhile(expr ast.WhileStmt, e *env.Env) typing.Throwable {
	for {
		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}

		_, err = EvalStmt(expr.Body, e)
		if err != nil {
			if _, ok := err.(typing.Break); ok {
				return nil
			} else if _, ok := err.(typing.Continue); ok {
				continue
			}
			return err
		}
	}
	return nil
}

func evalDoWhile(expr ast.DoWhileStmt, e *env.Env) typing.Throwable {
	for {
		_, err := EvalStmt(expr.Body, e)
		if err != nil {
			if _, ok := err.(typing.Break); ok {
				return nil
			} else if _, ok := err.(typing.Continue); ok {
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

func evalFun(expr ast.FunStmt, e *env.Env) typing.Throwable {
	fun := Function{
		Name:  expr.Name,
		PData: expr.Data,
		Stmt:  expr.Body,
		Env:   e,
	}

	return e.Define(expr.Name, fun, nil, false, true)
}

func evalReturn(expr ast.ReturnStmt, e *env.Env) typing.Throwable {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}
	return typing.Return{Data: val}
}

func evalBlock(expr ast.BlockStmt, e *env.Env) typing.Throwable {
	scope := env.NewEnclosedEnv(e)
	for _, stmt := range expr.Body {
		_, err := EvalStmt(stmt, scope)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalVariable(expr ast.VariableExpression, e *env.Env) (typing.Object, typing.Throwable) {
	return e.Get(expr.Name)
}

func evalCall(expr ast.CallExpression, e *env.Env) (typing.Object, typing.Throwable) {
	callee, err := evalExpression(expr.Callee, e)
	if err != nil {
		return nil, err
	}

	fun, ok := callee.(ast.Callable)
	if !ok {
		return nil, typing.NewError("You can only call functions")
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
		return nil, typing.NewError("More args needed")
	} else if len(args) > len(data.Params) && data.Args == "" {
		return nil, typing.NewError("Less args needed")
	}

	err = fun.Call(args)
	if ret, ok := err.(typing.Return); ok {
		return ret.Data, nil
	}
	return nil, err
}

func evalUnary(expr ast.UnaryExpression, e *env.Env) (typing.Object, typing.Throwable) {
	right, err := evalExpression(expr.Right, e)
	if err != nil {
		return nil, err
	}

	switch expr.Type {
	case token.Minus:
		if right, ok := right.(typing.Negable); ok {
			return right.Neg(), nil
		}
	case token.Bang:
		return typing.Bool{Value: !typing.Truthy(right)}, nil
	case token.Tilde:
		if right, ok := right.(typing.Notable); ok {
			return right.Not(), nil
		}
	}
	return nil, typing.NewError("Not implemented")
}

func evalBinary(expr ast.BinaryExpression, e *env.Env) (typing.Object, typing.Throwable) {
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
		return typing.Eq(left, right)
	case token.BangEqual:
		return typing.NEq(left, right)
	case token.Less:
		return typing.Lt(left, right)
	case token.LessEqual:
		return typing.Le(left, right)
	case token.Greater:
		return typing.Gt(left, right)
	case token.GreaterEqual:
		return typing.Ge(left, right)

	case token.Plus:
		return typing.Add(left, right)
	case token.Minus:
		return typing.Sub(left, right)
	case token.Asterisk:
		return typing.Mul(left, right)
	case token.Slash:
		return typing.Div(left, right)
	case token.Percent:
		return typing.Mod(left, right)
	case token.AsteriskAsterisk:
		return typing.Pow(left, right)

	case token.AmpersandAmpersand:
		if typing.Truthy(left) {
			return right, nil
		} else {
			return left, nil
		}
	case token.PipePipe:
		if typing.Truthy(left) {
			return left, nil
		} else {
			return right, nil
		}

	case token.LeftShift:
		return typing.LShift(left, right)
	case token.RightShift:
		return typing.RShift(left, right)
	case token.URightShift:
		return typing.URShift(left, right)
	case token.Pipe:
		return typing.Or(left, right)
	case token.Ampersand:
		return typing.And(left, right)
	case token.Circumflex:
		return typing.Xor(left, right)
	}

	return nil, typing.NewError("Not implemented")
}
