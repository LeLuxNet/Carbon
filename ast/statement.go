package ast

import (
	"github.com/leluxnet/carbon/token"
)

type Statement interface {
	astStatement()
}

type DeconData struct {
	Expr Expression
	T    Type
}

type VarStmt struct {
	Annotations
	Names map[string]DeconData
	Expr  Expression
	Const bool
}

type AssignStmt struct {
	Name string
	Type token.TokenType
	Expr Expression
}

type IfStmt struct {
	Condition Expression
	Then      Statement
	Else      Statement
}

type WhileStmt struct {
	Condition Expression
	Body      Statement
}

type DoWhileStmt struct {
	Condition Expression
	Body      Statement
}

type FunStmt struct {
	Annotations
	Name string
	Data ParamData
	Body Statement
}

type ConStmt struct {
	Name string
	Data ParamData
	Body Statement
}

type GetterStmt struct {
	Name string
	Body Statement
}

type ClassStmt struct {
	Name       string
	Properties []Statement
}

type ReturnStmt struct {
	Expr Expression
}

type BreakStmt struct{}

type ContinueStmt struct{}

type ExportStmt struct {
	Body Statement
}

type BlockStmt struct {
	Body []Statement
}

type ExpressionStmt struct {
	Expr Expression
}

type SetPropertyStatement struct {
	Target Expression
	Name   string
	Object Expression
}

func (s VarStmt) astStatement()              {}
func (s AssignStmt) astStatement()           {}
func (s IfStmt) astStatement()               {}
func (s WhileStmt) astStatement()            {}
func (s DoWhileStmt) astStatement()          {}
func (s ClassStmt) astStatement()            {}
func (s FunStmt) astStatement()              {}
func (s ConStmt) astStatement()              {}
func (s GetterStmt) astStatement()           {}
func (s ReturnStmt) astStatement()           {}
func (s BreakStmt) astStatement()            {}
func (s ContinueStmt) astStatement()         {}
func (s ExportStmt) astStatement()           {}
func (s BlockStmt) astStatement()            {}
func (s ExpressionStmt) astStatement()       {}
func (s SetPropertyStatement) astStatement() {}
