package eval

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/env"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/type_a"
	"github.com/leluxnet/carbon/typing"
)

func Eval(stmts []ast.Statement, e *env.Env, printRes bool, fileName string, path string) (*typing.File, typing.Throwable) {
	file := typing.NewFile(fileName, path)
	for _, stmt := range stmts {
		val, err := evalStmt(stmt, e, file)
		if err != nil {
			return nil, err
		}

		if printRes && val != nil {
			fmt.Println(val.ToString())
		}
	}
	return file, nil
}

func evalStmt(stmt ast.Statement, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	switch stmt := stmt.(type) {
	case ast.VarStmt:
		return nil, evalVar(stmt, e, file)
	case ast.AssignStmt:
		return nil, evalAssignment(stmt, e, file)
	case ast.IfStmt:
		return nil, evalIf(stmt, e, file)
	case ast.WhileStmt:
		return nil, evalWhile(stmt, e, file)
	case ast.DoWhileStmt:
		return nil, evalDoWhile(stmt, e, file)
	case ast.ClassStmt:
		_, _, err := evalClass(stmt, e, file)
		return nil, err
	case ast.FunStmt:
		_, err := evalFun(stmt, e, file)
		return nil, err
	case ast.SwitchStmt:
		return nil, evalSwitch(stmt, e, file)
	case ast.ReturnStmt:
		return nil, evalReturn(stmt, e, file)
	case ast.BreakStmt:
		return nil, typing.Break{}
	case ast.ContinueStmt:
		return nil, typing.Continue{}
	case ast.ExportStmt:
		return nil, evalExport(stmt, e, file)
	case ast.BlockStmt:
		return nil, evalBlock(stmt, e, file)
	case ast.ExpressionStmt:
		return evalExpression(stmt.Expr, e, file)
	case ast.SetPropertyStatement:
		return nil, evalSProperty(stmt, e, file)
	}
	panic("Parser let a unknown statement through")
}

func evalExpression(expr ast.Expression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	switch expr := expr.(type) {
	case ast.LiteralExpression:
		return expr.Object, nil
	case ast.GroupingExpression:
		return evalExpression(expr.Expr, e, file)
	case ast.ArrayExpression:
		return evalArray(expr, e, file)
	case ast.MapExpression:
		return evalMap(expr, e, file)
	case ast.SetExpression:
		return evalSet(expr, e, file)
	case ast.TupleExpression:
		return evalTuple(expr, e, file)
	case ast.LambdaExpression:
		return evalLambda(expr, e)
	case ast.VariableExpression:
		return evalVariable(expr, e)
	case ast.CallExpression:
		return evalCall(expr, e, file)
	case ast.IndexExpression:
		return evalIndex(expr, e, file)
	case ast.PropertyExpression:
		return evalProperty(expr, e, file)
	case ast.UnaryExpression:
		return evalUnary(expr, e, file)
	case ast.BinaryExpression:
		return evalBinary(expr, e, file)
	}
	panic("Parser let a unknown expression through")
}

func getVar(expr ast.VarStmt, e *env.Env, file *typing.File) (map[string]DeconRes, bool, typing.Throwable) {
	_, static := expr.Annotations["static"]
	var err typing.Throwable

	var val typing.Object
	if expr.Expr == nil {
		val = typing.Null{}
	} else {
		val, err = evalExpression(expr.Expr, e, file)
		if err != nil {
			return nil, false, err
		}
	}

	data, err := deconstruct(val, expr.Names, e, file)
	return data, static, err
}

func evalVar(expr ast.VarStmt, e *env.Env, file *typing.File) typing.Throwable {
	data, static, err := getVar(expr, e, file)
	if err != nil {
		return err
	}

	if static {
		return typing.NewError("Only class fields can be static")
	}

	for name, val := range data {
		err = e.Define(name, val.Val, val.Type, val.Val == nil, expr.Const)
		if err != nil {
			return err
		}
	}

	return nil
}

type DeconRes struct {
	Val  typing.Object
	Type typing.Type
}

func fromAstType(t ast.Type, e *env.Env) (typing.Type, typing.Throwable) {
	switch t := t.(type) {
	case ast.Class:
		cl, err := e.Get(t.Name)
		if err != nil {
			return nil, err
		}

		class, ok := cl.(typing.Class)
		if !ok {
			return nil, typing.NewError("Type annotation can't be an object")
		}

		return class, nil
	case ast.Array:
		val, err := fromAstType(t.Type, e)
		if err != nil {
			return nil, err
		}
		return type_a.Array{Type: val}, nil
	case ast.Map:
		key, err := fromAstType(t.Key, e)
		if err != nil {
			return nil, err
		}

		val, err := fromAstType(t.Value, e)
		if err != nil {
			return nil, err
		}

		return type_a.Map{Key: key, Value: val}, nil
	case ast.Set:
		val, err := fromAstType(t.Type, e)
		if err != nil {
			return nil, err
		}
		return type_a.Set{Type: val}, nil
	case ast.Tuple:
		var res []typing.Type
		for _, v := range t.Values {
			val, err := fromAstType(v, e)
			if err != nil {
				return nil, err
			}
			res = append(res, val)
		}
		return type_a.Tuple{Values: res}, nil
	}
	panic("Parser let some a unknown type annotation through")
}

func deconstruct(val typing.Object, names map[string]ast.DeconData, e *env.Env, file *typing.File) (map[string]DeconRes, typing.Throwable) {
	res := make(map[string]DeconRes, len(names))
	for name, prop := range names {
		var t typing.Type
		if prop.T != nil {
			var err typing.Throwable
			t, err = fromAstType(prop.T, e)
			if err != nil {
				return nil, err
			}
		}

		if prop.Expr == nil {
			res[name] = DeconRes{val, t}
		} else {
			p, err := evalExpression(prop.Expr, e, file)
			if err != nil {
				return nil, err
			}

			pVal, err := getIndex(val, p)
			if err != nil {
				return nil, err
			}

			res[name] = DeconRes{pVal, nil}
		}
	}
	return res, nil
}

func evalAssignment(expr ast.AssignStmt, e *env.Env, file *typing.File) typing.Throwable {
	val, err := evalExpression(expr.Expr, e, file)
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

func evalIf(expr ast.IfStmt, e *env.Env, file *typing.File) typing.Throwable {
	condition, err := evalExpression(expr.Condition, e, file)
	if err != nil {
		return err
	}

	if typing.Truthy(condition) {
		_, err := evalStmt(expr.Then, e, file)
		return err
	} else if expr.Else != nil {
		_, err := evalStmt(expr.Else, e, file)
		return err
	}
	return nil
}

func evalWhile(expr ast.WhileStmt, e *env.Env, file *typing.File) typing.Throwable {
	for {
		condition, err := evalExpression(expr.Condition, e, file)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}

		_, err = evalStmt(expr.Body, e, file)
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

func evalDoWhile(expr ast.DoWhileStmt, e *env.Env, file *typing.File) typing.Throwable {
	for {
		_, err := evalStmt(expr.Body, e, file)
		if err != nil {
			if _, ok := err.(typing.Break); ok {
				return nil
			} else if _, ok := err.(typing.Continue); ok {
				continue
			}
			return err
		}

		condition, err := evalExpression(expr.Condition, e, file)
		if err != nil {
			return err
		} else if !typing.Truthy(condition) {
			break
		}
	}
	return nil
}

func fromAstParamData(data ast.ParamData, e *env.Env, file *typing.File) (*typing.ParamData, typing.Throwable) {
	params := make([]typing.Parameter, len(data.Params))
	var err typing.Throwable

	for i, p := range data.Params {
		var t typing.Type
		if p.Type != nil {
			t, err = fromAstType(p.Type, e)
		}

		var d typing.Object
		if p.Default != nil {
			d, err = evalExpression(p.Default, e, file)
		}

		params[i] = typing.Parameter{
			Name:    p.Name,
			Type:    t,
			Default: d,
		}
	}

	var r typing.Type
	if data.Return != nil {
		r, err = fromAstType(data.Return, e)
		if err != nil {
			return nil, err
		}
	}

	return &typing.ParamData{
		Params: params,
		Args:   data.Args,
		KwArgs: data.KwArgs,
		Return: r,
	}, nil
}

func evalClass(expr ast.ClassStmt, e *env.Env, file *typing.File) (string, typing.Object, typing.Throwable) {
	p := make(typing.Properties)
	sp := make(typing.Properties)

	var staticProps []ast.Statement
	for _, val := range expr.Properties {
		switch val := val.(type) {
		case ast.VarStmt:
			_, static := val.Annotations["static"]
			if static {
				staticProps = append(staticProps, val)
				continue
			}

			data, _, err := getVar(val, e, file)
			if err != nil {
				return "", nil, err
			}

			for name, val := range data {
				p[name] = val.Val
			}
		case ast.FunStmt:
			_, static := val.Annotations["static"]
			if static {
				staticProps = append(staticProps, val)
				continue
			}

			fun, static, err := getFun(val, e, file)
			if err != nil {
				return "", nil, err
			}

			if static {
				sp[fun.Name] = *fun
			} else {
				p[fun.Name] = *fun
			}
		case ast.ConStmt:
			data, err := fromAstParamData(val.Data, e, file)
			if err != nil {
				return "", nil, err
			}

			con := Constructor{
				Name:  val.Name,
				PData: *data,
				Stmt:  val.Body,
				Env:   e,
			}

			sp[con.Name] = con
		case ast.GetterStmt:
			get := Getter{
				Name: val.Name,
				Stmt: val.Body,
				Env:  e,
			}

			if v, ok := p[get.Name]; ok {
				x := v.(typing.GSetter)
				x.Getter = get
				p[get.Name] = x
			} else {
				p[get.Name] = typing.GSetter{Getter: get}
			}
		case ast.SetterStmt:
			params, err := fromAstParamData(ast.ParamData{Params: []ast.Parameter{val.Param}}, e, file)
			if err != nil {
				return "", nil, err
			}

			set := Setter{
				Name:  val.Name,
				Param: params.Params[0],
				Stmt:  val.Body,
				Env:   e,
			}

			if v, ok := p[set.Name]; ok {
				x := v.(typing.GSetter)
				x.Setter = set
				p[set.Name] = x
			} else {
				p[set.Name] = typing.GSetter{Setter: set}
			}
		default:
			panic("Parser allowed unknown statement inside a class")
		}
	}

	class := typing.Class{
		Name:       expr.Name,
		Properties: p,
		Static:     sp,
	}

	err := e.Define(expr.Name, class, nil, false, true)
	if err != nil {
		return "", nil, err
	}

	for _, val := range staticProps {
		switch val := val.(type) {
		case ast.VarStmt:
			data, _, err := getVar(val, e, file)
			if err != nil {
				return "", nil, err
			}

			for name, val := range data {
				sp[name] = val.Val
			}
		case ast.FunStmt:
			fun, _, err := getFun(val, e, file)
			if err != nil {
				return "", nil, err
			}
			sp[fun.Name] = fun
		}
	}

	return expr.Name, class, nil
}

func getFun(expr ast.FunStmt, e *env.Env, file *typing.File) (*Function, bool, typing.Throwable) {
	_, static := expr.Annotations["static"]

	data, err := fromAstParamData(expr.Data, e, file)
	if err != nil {
		return nil, false, err
	}

	return &Function{
		Name:  expr.Name,
		PData: *data,
		Stmt:  expr.Body,
		Env:   e,
	}, static, nil
}

func evalFun(expr ast.FunStmt, e *env.Env, file *typing.File) (*Function, typing.Throwable) {
	fun, static, err := getFun(expr, e, file)
	if err != nil {
		return nil, err
	}

	if static {
		return fun, typing.NewError("Only functions in class can be static")
	}

	return fun, e.Define(expr.Name, fun, nil, false, true)
}

func evalSwitch(expr ast.SwitchStmt, e *env.Env, file *typing.File) typing.Throwable {
	val, err := evalExpression(expr.Value, e, file)
	if err != nil {
		return err
	}

	for _, c := range expr.Body {
		for _, cas := range c.Cases {
			o, err := evalExpression(cas, e, file)
			if err != nil {
				return err
			}

			eq, err := typing.Eq(val, o)
			if err != nil {
				return err
			}

			if typing.Truthy(eq) {
				_, err = evalStmt(c.Body, e, file)
				return err
			}
		}
	}
	return nil
}

func evalReturn(expr ast.ReturnStmt, e *env.Env, file *typing.File) typing.Throwable {
	val, err := evalExpression(expr.Expr, e, file)
	if err != nil {
		return err
	}
	return typing.Return{Data: val}
}

func evalExport(expr ast.ExportStmt, e *env.Env, file *typing.File) typing.Throwable {
	switch body := expr.Body.(type) {
	case ast.ClassStmt:
		name, class, err := evalClass(body, e, file)
		if err != nil {
			return err
		}

		file.Props[name] = class
		return nil
	case ast.FunStmt:
		fun, err := evalFun(body, e, file)
		if err != nil {
			return err
		}
		file.Props[body.Name] = *fun
		return nil
	case ast.ExpressionStmt:
		if expr, ok := body.Expr.(ast.VariableExpression); ok {
			val, err := e.Get(expr.Name)
			if err != nil {
				return err
			}
			file.Props[expr.Name] = val
			return nil
		}
	}
	return typing.NewError("Only functions, classes and variables can be exported")
}

func evalBlock(expr ast.BlockStmt, e *env.Env, file *typing.File) typing.Throwable {
	scope := env.NewEnclosedEnv(e)
	for _, stmt := range expr.Body {
		_, err := evalStmt(stmt, scope, file)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalArray(expr ast.ArrayExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	values := make([]typing.Object, len(expr.Values))

	for i, rVal := range expr.Values {
		val, err := evalExpression(rVal, e, file)
		if err != nil {
			return nil, err
		}

		values[i] = val
	}

	return typing.Array{Values: values}, nil
}

func evalMap(expr ast.MapExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	res := typing.NewMap()

	for rKey, rValue := range expr.Items {
		key, err := evalExpression(rKey, e, file)
		if err != nil {
			return nil, err
		}

		value, err := evalExpression(rValue, e, file)
		if err != nil {
			return nil, err
		}

		res.SetIndex(key, value)
	}

	return res, nil
}

func evalSet(expr ast.SetExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	res := typing.NewSet()

	for _, rVal := range expr.Values {
		val, err := evalExpression(rVal, e, file)
		if err != nil {
			return nil, err
		}

		err2 := res.Append(val)
		if err2 != nil {
			return nil, typing.Throw{Data: err2}
		}
	}

	return res, nil
}

func evalTuple(expr ast.TupleExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	var values []typing.Object

	for _, rVal := range expr.Values {
		val, err := evalExpression(rVal, e, file)
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

func evalCall(expr ast.CallExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	var this typing.Object
	var callee typing.Object
	var err typing.Throwable

	if prop, ok := expr.Target.(ast.PropertyExpression); ok {
		this, err = evalExpression(prop.Target, e, file)
		if err != nil {
			return nil, err
		}

		callee, err = getProperty(this, prop.Name, file)
	} else {
		callee, err = evalExpression(expr.Target, e, file)
	}
	if err != nil {
		return nil, err
	}

	var data typing.ParamData

	var fun typing.Callable
	var con *Constructor

	if expr.New {
		switch callee := callee.(type) {
		case Constructor:
			con = &callee
			data = con.PData
		case typing.Class:
			this = callee
			tmp, ok := callee.Static[""]
			if ok {
				tmpCon := tmp.(Constructor)
				con = &tmpCon
				data = con.PData
			} else {
				data = typing.ParamData{}
			}
		default:
			return nil, typing.NewError("You can only use new to call a constructor")
		}
	} else {
		var ok bool
		fun, ok = callee.(typing.Callable)
		if !ok {
			return nil, typing.NewError("You can only call functions")
		}
		data = fun.Data()
	}

	params := make(map[string]typing.Object)

	minArgs := 0
	scope := 0
	var j int
	for i, arg := range data.Params {
		j = i
		if arg.Default == nil {
			minArgs++
		}

		if scope == 0 {
			if len(expr.Args) > i {
				object, err := evalExpression(expr.Args[i], e, file)
				if err != nil {
					return nil, err
				}

				if arg.Type != nil && !arg.Type.Allows(object) {
					return nil, typing.NewError(fmt.Sprintf("Type error: Expect '%s', got '%s'", arg.Type.TName(), object.Class().Name))
				}

				params[arg.Name] = object
			} else {
				scope = 1
			}
		}

		if scope == 1 {
			if a, ok := expr.KwArgs[arg.Name]; ok {
				object, err := evalExpression(a, e, file)
				if err != nil {
					return nil, err
				}

				params[arg.Name] = object
				delete(expr.KwArgs, arg.Name)
			} else if arg.Default != nil {
				params[arg.Name] = arg.Default
			} else {
				return nil, typing.NewError("Missing a non optional argument")
			}
		}
	}

	arg := append(expr.Args[j:], expr.Args2...)
	args := make([]typing.Object, len(arg))
	for i, a := range arg {
		object, err := evalExpression(a, e, file)
		if err != nil {
			return nil, err
		}

		args[i] = object
	}

	for name, exp := range expr.KwArgs {
		expr.KwArgs2[name] = exp
	}
	kwArgs := make(map[string]typing.Object, len(expr.KwArgs2))
	for name, a := range expr.KwArgs2 {
		object, err := evalExpression(a, e, file)
		if err != nil {
			return nil, err
		}

		kwArgs[name] = object
	}

	if expr.New {
		class := this.(typing.Class)
		i := typing.Instance{Clazz: class, Fields: make(map[string]typing.Object), File: file}
		if con != nil {
			err = con.Const(i, params, args, kwArgs, file)
		}
		return i, nil
	} else {
		err = fun.Call(this, params, args, kwArgs, file)
		if ret, ok := err.(typing.Return); ok {
			return ret.Data, nil
		}
	}
	return nil, err
}

func evalIndex(expr ast.IndexExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	target, err := evalExpression(expr.Target, e, file)
	if err != nil {
		return nil, err
	}

	if t, ok := target.(typing.IndexGettable); ok {
		index, err := evalExpression(expr.Index, e, file)
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

func getIndex(o typing.Object, i typing.Object) (typing.Object, typing.Throwable) {
	if t, ok := o.(typing.IndexGettable); ok {
		res, err := t.GetIndex(i)
		if err != nil {
			return nil, typing.Throw{Data: err}
		}
		return res, nil
	}

	return nil, typing.NewError(fmt.Sprintf("'%s' has not support for getting indexes", o.Class().Name))
}

func evalProperty(expr ast.PropertyExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	target, err := evalExpression(expr.Target, e, file)
	if err != nil {
		return nil, err
	}

	return getProperty(target, expr.Name, file)
}

func evalSProperty(expr ast.SetPropertyStatement, e *env.Env, file *typing.File) typing.Throwable {
	target, err := evalExpression(expr.Target, e, file)
	if err != nil {
		return err
	}

	o, err := evalExpression(expr.Object, e, file)
	if err != nil {
		return err
	}

	return setProperty(target, expr.Name, o, file)
}

func getProperty(o typing.Object, name string, file *typing.File) (typing.Object, typing.Throwable) {
	var p typing.Object
	var err typing.Object
	if o2, ok := o.(typing.PropertyGettable); ok {
		p, err = o2.GetProperty(name)
		if err != nil {
			return nil, typing.Throw{Data: err}
		}
	} else {
		p, ok = o.Class().Properties[name]
		if !ok {
			return nil, typing.NewError(fmt.Sprintf("'%s' has no such property '%s'", o.ToString(), name))
		}
	}

	if get, ok := p.(typing.GSetter); ok {
		return get.Get(o, file)
	}

	return p, nil
}

func setProperty(o typing.Object, name string, val typing.Object, file *typing.File) typing.Throwable {
	if o, ok := o.(typing.PropertySettable); ok {
		err := o.SetProperty(name, val, file)
		if err != nil {
			return typing.Throw{Data: err}
		}
		return nil
	}

	old := o.Class().Properties[name]
	if old, ok := old.(typing.GSetter); ok {
		return old.Set(o, val, file)
	}

	o.Class().Properties[name] = val
	return nil
}

func evalUnary(expr ast.UnaryExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	right, err := evalExpression(expr.Right, e, file)
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

func evalBinary(expr ast.BinaryExpression, e *env.Env, file *typing.File) (typing.Object, typing.Throwable) {
	left, err := evalExpression(expr.Left, e, file)
	if err != nil {
		return nil, err
	}

	right, err := evalExpression(expr.Right, e, file)
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

	panic("Parsed allowed unknown binary expression type_a")
}
