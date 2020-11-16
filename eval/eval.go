package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/hash"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

type strObject map[string]typing.Object

func Eval(stmts []ast.Statement, e *env.Env, printRes bool) (map[string]typing.Object, typing.Throwable) {
	props := make(strObject)
	for _, stmt := range stmts {
		val, err := evalStmt(stmt, e, props)
		if err != nil {
			return nil, err
		}

		if printRes && val != nil {
			fmt.Println(val.ToString())
		}
	}
	return props, nil
}

func evalStmt(stmt ast.Statement, e *env.Env, props strObject) (typing.Object, typing.Throwable) {
	switch stmt := stmt.(type) {
	case ast.VarStmt:
		return nil, evalVar(stmt, e)
	case ast.ValStmt:
		return nil, evalVal(stmt, e)
	case ast.AssignStmt:
		return nil, evalAssignment(stmt, e)
	case ast.IfStmt:
		return nil, evalIf(stmt, e, props)
	case ast.WhileStmt:
		return nil, evalWhile(stmt, e, props)
	case ast.DoWhileStmt:
		return nil, evalDoWhile(stmt, e, props)
	case ast.ClassStmt:
		name, class, err := getClass(stmt, e, props)
		if err != nil {
			return nil, err
		}
		return nil, e.Define(name, class, nil, false, true)
	case ast.FunStmt:
		return nil, evalFun(stmt, e)
	case ast.ReturnStmt:
		return nil, evalReturn(stmt, e)
	case ast.BreakStmt:
		return nil, typing.Break{}
	case ast.ContinueStmt:
		return nil, typing.Continue{}
	case ast.ImportStmt:
		return nil, evalImport(stmt, e)
	case ast.ExportStmt:
		return nil, evalExport(stmt, e, props)
	case ast.BlockStmt:
		return nil, evalBlock(stmt, e, props)
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
	case ast.ArrayExpression:
		return evalArray(expr, e)
	case ast.MapExpression:
		return evalMap(expr, e)
	case ast.SetExpression:
		return evalSet(expr, e)
	case ast.TupleExpression:
		return evalTuple(expr, e)
	case ast.LambdaExpression:
		return evalLambda(expr, e)
	case ast.VariableExpression:
		return evalVariable(expr, e)
	case ast.CallExpression:
		return evalCall(expr, e)
	case ast.IndexExpression:
		return evalIndex(expr, e)
	case ast.PropertyExpression:
		return evalProperty(expr, e)
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

func evalIf(expr ast.IfStmt, e *env.Env, props strObject) typing.Throwable {
	condition, err := evalExpression(expr.Condition, e)
	if err != nil {
		return err
	}

	if typing.Truthy(condition) {
		_, err := evalStmt(expr.Then, e, props)
		return err
	} else if expr.Else != nil {
		_, err := evalStmt(expr.Else, e, props)
		return err
	}
	return nil
}

func evalWhile(expr ast.WhileStmt, e *env.Env, props strObject) typing.Throwable {
	for {
		condition, err := evalExpression(expr.Condition, e)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}

		_, err = evalStmt(expr.Body, e, props)
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

func evalDoWhile(expr ast.DoWhileStmt, e *env.Env, props strObject) typing.Throwable {
	for {
		_, err := evalStmt(expr.Body, e, props)
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

func getClass(expr ast.ClassStmt, e *env.Env, props strObject) (string, typing.Object, typing.Throwable) {
	p := make(typing.Properties, len(expr.Properties))
	for name, val := range expr.Properties {
		if val, ok := val.(ast.FunStmt); ok {
			fun := Function{
				Name:  val.Name,
				PData: val.Data,
				Stmt:  val.Body,
				Env:   e,
			}

			p[name] = fun
		} else {
			v, err := evalStmt(val, e, props)
			if err != nil {
				return "", nil, err
			}

			p[name] = v
		}
	}

	class := typing.Class{
		Name:       expr.Name,
		Properties: p,
	}

	return "", class, nil
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

func evalImport(expr ast.ImportStmt, e *env.Env) typing.Throwable {
	props := Import(expr.Module)

	m := typing.NewMap()
	for name, o := range props {
		m.Items[hash.HashString(name)] = typing.Pair{Key: typing.String{Value: name}, Value: o}
	}

	return e.Define(expr.Name, m, nil, false, true)
}

func evalExport(expr ast.ExportStmt, e *env.Env, props strObject) typing.Throwable {
	switch body := expr.Body.(type) {
	case ast.ClassStmt:
		name, class, err := getClass(body, e, props)
		if err != nil {
			return err
		}
		props[name] = class
		return nil
	case ast.FunStmt:
		props[body.Name] = Function{
			Name:  body.Name,
			PData: body.Data,
			Stmt:  body.Body,
			Env:   e,
		}
		return nil
	case ast.ExpressionStmt:
		if expr, ok := body.Expr.(ast.VariableExpression); ok {
			val, err := e.Get(expr.Name)
			if err != nil {
				return err
			}
			props[expr.Name] = val
			return nil
		}
	}
	return typing.NewError("Only functions, classes and variables can be exported")
}

func evalBlock(expr ast.BlockStmt, e *env.Env, props strObject) typing.Throwable {
	scope := env.NewEnclosedEnv(e)
	for _, stmt := range expr.Body {
		_, err := evalStmt(stmt, scope, props)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalArray(expr ast.ArrayExpression, e *env.Env) (typing.Object, typing.Throwable) {
	values := make([]typing.Object, len(expr.Values))

	for i, rVal := range expr.Values {
		val, err := evalExpression(rVal, e)
		if err != nil {
			return nil, err
		}

		values[i] = val
	}

	return typing.Array{Values: values}, nil
}

func evalMap(expr ast.MapExpression, e *env.Env) (typing.Object, typing.Throwable) {
	res := typing.NewMap()

	for rKey, rValue := range expr.Items {
		key, err := evalExpression(rKey, e)
		if err != nil {
			return nil, err
		}

		value, err := evalExpression(rValue, e)
		if err != nil {
			return nil, err
		}

		res.SetIndex(key, value)
	}

	return res, nil
}

func evalSet(expr ast.SetExpression, e *env.Env) (typing.Object, typing.Throwable) {
	res := typing.NewSet()

	for _, rVal := range expr.Values {
		val, err := evalExpression(rVal, e)
		if err != nil {
			return nil, err
		}

		res.Append(val)
	}

	return res, nil
}

func evalTuple(expr ast.TupleExpression, e *env.Env) (typing.Object, typing.Throwable) {
	var values []typing.Object

	for _, rVal := range expr.Values {
		val, err := evalExpression(rVal, e)
		if err != nil {
			return nil, err
		}

		values = append(values, val)
	}

	return typing.Tuple{Values: values}, nil
}

func evalVariable(expr ast.VariableExpression, e *env.Env) (typing.Object, typing.Throwable) {
	return e.Get(expr.Name)
}

func evalLambda(expr ast.LambdaExpression, e *env.Env) (typing.Object, typing.Throwable) {
	fun := Function{
		PData: expr.Data,
		Stmt:  expr.Body,
		Env:   e,
	}

	return fun, nil
}

func evalCall(expr ast.CallExpression, e *env.Env) (typing.Object, typing.Throwable) {
	var this typing.Object
	var callee typing.Object
	var err typing.Throwable

	if prop, ok := expr.Target.(ast.PropertyExpression); ok {
		this, err = evalExpression(prop.Target, e)
		if err != nil {
			return nil, err
		}

		callee, err = getProperty(this.Class(), prop.Name)
	} else {
		callee, err = evalExpression(expr.Target, e)
	}
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

	err = fun.Call(this, args)
	if ret, ok := err.(typing.Return); ok {
		return ret.Data, nil
	}
	return nil, err
}

func evalIndex(expr ast.IndexExpression, e *env.Env) (typing.Object, typing.Throwable) {
	target, err := evalExpression(expr.Target, e)
	if err != nil {
		return nil, err
	}

	if t, ok := target.(typing.IndexGettable); ok {
		index, err := evalExpression(expr.Index, e)
		if err != nil {
			return nil, err
		}

		res, err2 := t.GetIndex(index)
		if err2 != nil {
			return nil, typing.Throw{Data: err2}
		}
		return res, nil
	}

	return nil, typing.NewError(fmt.Sprintf("'%s' has not support for getting indexes", target.Class().Name))
}

func evalProperty(expr ast.PropertyExpression, e *env.Env) (typing.Object, typing.Throwable) {
	target, err := evalExpression(expr.Target, e)
	if err != nil {
		return nil, err
	}

	return getProperty(target.Class(), expr.Name)
}

func getProperty(o typing.Object, name string) (typing.Object, typing.Throwable) {
	p, ok := o.Class().Properties[name]
	if !ok {
		return nil, typing.NewError(fmt.Sprintf("'%s' has no such property '%s'", o.ToString(), name))
	}

	return p, nil
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
