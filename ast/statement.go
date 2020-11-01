package ast

import "github.com/leluxnet/carbon/token"

type Statement interface {
	astStatement()
}

type VarStmt struct {
	Name string
	Expr Expression
}

type ValStmt struct {
	Name string
	Expr Expression
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
	Data ParamData
	Body Statement
}

type ReturnStmt struct {
	Expr Expression
}

type BreakStmt struct{}

type ContinueStmt struct{}

type BlockStmt struct {
	Body []Statement
}

type ExpressionStmt struct {
	Expr Expression
}

func (s VarStmt) astStatement()        {}
func (s ValStmt) astStatement()        {}
func (s AssignStmt) astStatement()     {}
func (s IfStmt) astStatement()         {}
func (s WhileStmt) astStatement()      {}
func (s DoWhileStmt) astStatement()    {}
func (s FunStmt) astStatement()        {}
func (s ReturnStmt) astStatement()     {}
func (s BreakStmt) astStatement()      {}
func (s ContinueStmt) astStatement()   {}
func (s BlockStmt) astStatement()      {}
func (s ExpressionStmt) astStatement() {}
