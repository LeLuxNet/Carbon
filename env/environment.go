package env

import (
	"fmt"
	"github.com/leluxnet/carbon/typing"
)

type Env struct {
	vars  map[string]*Variable
	outer *Env
}

func NewEnclosedEnv(outer *Env) *Env {
	env := NewEnv()
	env.outer = outer
	return env
}

func NewEnv() *Env {
	return &Env{vars: make(map[string]*Variable)}
}

func (e Env) Declare(name string, class *typing.Class) typing.Throwable {
	if _, ok := e.vars[name]; ok {
		return typing.NewError(fmt.Sprintf("Variable '%s' is already declared", name))
	}

	e.vars[name] = &Variable{Type: class, Value: typing.Null{}, Nullable: true}
	return nil
}

func (e Env) Define(name string, object typing.Object, class *typing.Class, nullable bool, constant bool) typing.Throwable {
	if _, ok := e.vars[name]; ok {
		return typing.NewError(fmt.Sprintf("Variable '%s' is already declared", name))
	}

	if class == nil {
		t := object.Class()
		class = &t
	}

	e.vars[name] = &Variable{Type: class, Value: object, Nullable: nullable, Constant: constant}
	return nil
}

func (e Env) Set(name string, object typing.Object) typing.Throwable {
	if v, ok := e.vars[name]; ok {
		if v.Constant {
			return typing.NewError(fmt.Sprintf("Variable '%s' is constant and can't be reassigned", name))
		} else if v.Type.IsInstance(object) || (object.Class().Name == "null" && v.Nullable) {
			v.Value = object
			return nil
		} else {
			return typing.NewError(fmt.Sprintf("Variable '%s' of type '%s' can't be assigned to '%s'",
				name, v.Type.Name, object.Class().Name))
		}
	}

	if e.outer != nil {
		return e.outer.Set(name, object)
	}

	return typing.NewError("Variable is not declared")
}

func (e Env) Get(name string) (typing.Object, typing.Throwable) {
	if v, ok := e.vars[name]; ok {
		return v.Value, nil
	}

	if e.outer != nil {
		return e.outer.Get(name)
	}

	return nil, typing.NewError(fmt.Sprintf("No such var '%s'", name))
}
