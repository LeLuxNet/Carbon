package eval

import (
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

func Eval(stmts []ast.Statement, e *env.Env) typing.Object {
	for _, stmt := range stmts {
		err := evalStmt(stmt, e)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalStmt(stmt ast.Statement, e *env.Env) typing.Object {
	switch stmt := stmt.(type) {
	case ast.VarStmt:
		return evalVar(stmt, e)
	case ast.ValStmt:
		return evalVal(stmt, e)
	case ast.AssignStmt:
		return evalAssignment(stmt, e)
	case ast.IfStmt:
		return evalIf(stmt, e)
	case ast.WhileStmt:
		return evalWhile(stmt, e)
	case ast.DoWhileStmt:
		return evalDoWhile(stmt, e)
	case ast.BlockStmt:
		return evalBlock(stmt, e)
	case ast.ExpressionStmt:
		_, err := evalExpression(stmt.Expr, e)
		return err
	}
	return nil
}

func evalExpression(expr ast.Expression, e *env.Env) (typing.Object, typing.Object) {
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

func evalVar(expr ast.VarStmt, e *env.Env) typing.Object {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return val
	}

	return e.Define(expr.Name, val, nil, false, false)
}

func evalVal(expr ast.ValStmt, e *env.Env) typing.Object {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return val
	}

	return e.Define(expr.Name, val, nil, false, true)
}

func evalAssignment(expr ast.AssignStmt, e *env.Env) typing.Object {
	val, err := evalExpression(expr.Expr, e)
	if err != nil {
		return err
	}

	return e.Set(expr.Name, val)
}

func evalIf(expr ast.IfStmt, e *env.Env) typing.Object {
	condition, err := evalExpression(expr.Condition, e)
	if err != nil {
		return err
	}

	if typing.Truthy(condition) {
		return evalStmt(expr.Then, e)
	} else if expr.Else != nil {
		return evalStmt(expr.Else, e)
	}
	return nil
}

func evalWhile(expr ast.WhileStmt, e *env.Env) typing.Object {
	for true {
		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}

		evalStmt(expr.Body, e)
	}
	return nil
}

func evalDoWhile(expr ast.DoWhileStmt, e *env.Env) typing.Object {
	for true {
		evalStmt(expr.Body, e)

		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}
	}
	return nil
}

func evalBlock(expr ast.BlockStmt, e *env.Env) typing.Object {
	scope := env.NewEnclosedEnv(e)
	for _, stmt := range expr.Body {
		err := evalStmt(stmt, scope)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalVariable(expr ast.VariableExpression, e *env.Env) (typing.Object, typing.Object) {
	return e.Get(expr.Name)
}

func evalCall(expr ast.CallExpression, e *env.Env) (typing.Object, typing.Object) {
	callee, err := evalExpression(expr.Callee, e)
	if err != nil {
		return nil, err
	}

	fun, ok := callee.(typing.Callable)
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

	return fun.Call(args), nil
}

func evalUnary(expr ast.UnaryExpression, e *env.Env) (typing.Object, typing.Object) {
	_, err := evalExpression(expr.Right, e)
	if err != nil {
		return nil, err
	}

	// TODO: Calc unary
	return nil, typing.NewError("Not implemented")
}

func evalBinary(expr ast.BinaryExpression, e *env.Env) (typing.Object, typing.Object) {
	left, err := evalExpression(expr.Left, e)
	if err != nil {
		return nil, err
	}

	right, err := evalExpression(expr.Right, e)
	if err != nil {
		return nil, err
	}

	if expr.Type == token.EqualEqual {
		return typing.Bool{Value: typing.Eq(left, right)}, nil
	}

	// TODO: Calc binary
	return nil, typing.NewError("Not implemented")
}
