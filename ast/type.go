package ast

type Parameter struct {
	Name    string
	Type    Type
	Default Expression
}

type ParamData struct {
	Params []Parameter
	Args   string
	KwArgs string
	Return Type
}

type Type interface {
	astType()
}

type Class struct {
	Name string
}

type Array struct {
	Type
}

type Map struct {
	Key   Type
	Value Type
}

type Set struct {
	Type
}

type Tuple struct {
	Values []Type
}

func (s Class) astType() {}
func (s Array) astType() {}
func (s Map) astType()   {}
func (s Set) astType()   {}
func (s Tuple) astType() {}
