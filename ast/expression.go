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
	Callee    Expression
	Arguments []Expression
}

func (e LiteralExpression) astExpression()  {}
func (e UnaryExpression) astExpression()    {}
func (e BinaryExpression) astExpression()   {}
func (e GroupingExpression) astExpression() {}
func (e VariableExpression) astExpression() {}
func (e CallExpression) astExpression()     {}
