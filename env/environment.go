package env

import (
	"fmt"
	"github.com/leluxnet/carbon/throw"
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

func (e Env) Declare(name string, class *typing.Class) throw.Throwable {
	_, ok := e.vars[name]
	if ok {
		return throw.NewError("Already declared")
	}

	e.vars[name] = &Variable{Type: class, Value: typing.Null{}, Nullable: true}
	return nil
}

func (e Env) Define(name string, object typing.Object, class *typing.Class, nullable bool, constant bool) throw.Throwable {
	_, ok := e.vars[name]
	if ok {
		return throw.NewError("Variable is already declared")
	}

	if class == nil {
		t := object.Class()
		class = &t
	}

	e.vars[name] = &Variable{Type: class, Value: object, Nullable: nullable, Constant: constant}
	return nil
}

func (e Env) Set(name string, object typing.Object) throw.Throwable {
	v, ok := e.vars[name]
	if ok {
		if object.Class() == *v.Type || (object.Class().Name == "null" && v.Nullable) {
			v.Value = object
			return nil
		} else {
			return throw.NewError("Variable of type '" + v.Type.Name + "' cannot be assigned to '" + object.Class().Name + "'")
		}
	}

	if e.outer != nil {
		return e.outer.Set(name, object)
	}

	return throw.NewError("Variable is not declared")
}

func (e Env) Get(name string) (typing.Object, throw.Throwable) {
	v, ok := e.vars[name]

	if ok {
		return v.Value, nil
	}

	if e.outer != nil {
		return e.outer.Get(name)
	}

	return nil, throw.NewError(fmt.Sprintf("No such var '%s'", name))
}
