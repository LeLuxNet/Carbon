package ast

import (
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

type Expression interface {
	astExpression()
}

type LiteralExpression struct {
	Object typing.Object
}

type UnaryExpression struct {
	Type  token.TokenType
	Right Expression
}

type BinaryExpression struct {
	Left  Expression
	Type  token.TokenType
	Right Expression
}

type GroupingExpression struct {
	Expr Expression
}

type VariableExpression struct {
	Name string
}

type CallExpression struct {
	Target    Expression
	Arguments []Expression
}

type IndexExpression struct {
	Target Expression
	Index  Expression
}

type MapExpression struct {
	Items map[Expression]Expression
}

type ArrayExpression struct {
	Values []Expression
}

type SetExpression struct {
	Values []Expression
}

type TupleExpression struct {
	Values []Expression
}

type LambdaExpression struct {
	Data ParamData
	Body Statement
}

func (e LiteralExpression) astExpression()  {}
func (e UnaryExpression) astExpression()    {}
func (e BinaryExpression) astExpression()   {}
func (e GroupingExpression) astExpression() {}
func (e VariableExpression) astExpression() {}
func (e CallExpression) astExpression()     {}
func (e IndexExpression) astExpression()    {}
func (e MapExpression) astExpression()      {}
func (e ArrayExpression) astExpression()    {}
func (e SetExpression) astExpression()      {}
func (e TupleExpression) astExpression()    {}
func (e LambdaExpression) astExpression()   {}
