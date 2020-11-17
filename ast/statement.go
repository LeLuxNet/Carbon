package ast

import (
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
)

type Statement interface {
	astStatement()
}

type VarStmt struct {
	Names map[string]Expression
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
	Name string
	Data typing.ParamData
	Body Statement
}

type ClassStmt struct {
	Name       string
	Properties map[string]Statement
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

func (s VarStmt) astStatement()        {}
func (s AssignStmt) astStatement()     {}
func (s IfStmt) astStatement()         {}
func (s WhileStmt) astStatement()      {}
func (s DoWhileStmt) astStatement()    {}
func (s ClassStmt) astStatement()      {}
func (s FunStmt) astStatement()        {}
func (s ReturnStmt) astStatement()     {}
func (s BreakStmt) astStatement()      {}
func (s ContinueStmt) astStatement()   {}
func (s ExportStmt) astStatement()     {}
func (s BlockStmt) astStatement()      {}
func (s ExpressionStmt) astStatement() {}
