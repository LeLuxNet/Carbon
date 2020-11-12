package parser

import (
	"fmt"
	"github.com/leluxnet/carbon/ast"
	"github.com/leluxnet/carbon/errors"
	"github.com/leluxnet/carbon/token"
	"github.com/leluxnet/carbon/typing"
	"math/big"
)

type Parser struct {
	Tokens   []token.Token
	Position int

	errs []errors.SyntaxError
}

func (p *Parser) Parse() ([]ast.Statement, []errors.SyntaxError) {
	var statements []ast.Statement

	for p.Position < len(p.Tokens) {
		stmt, err := p.statement()
		if err != nil {
			p.catch(err)
		} else {
			statements = append(statements, stmt)
		}
	}

	return statements, p.errs
}

func (p *Parser) statement() (ast.Statement, *errors.SyntaxError) {
	var res ast.Statement
	var err *errors.SyntaxError

	if p.match(token.Var) {
		res, err = p.varStmt()
	} else if p.match(token.Val) {
		res, err = p.valStmt()
	} else if p.match(token.If) {
		return p.ifStmt()
	} else if p.match(token.While) {
		return p.whileStmt()
	} else if p.match(token.Do) {
		res, err = p.doWhileStmt()
	} else if p.match(token.Fun) {
		return p.funStmt()
	} else if p.match(token.Return) {
		res, err = p.returnStmt()
	} else if p.match(token.Break) {
		res = ast.BreakStmt{}
	} else if p.match(token.Continue) {
		res = ast.ContinueStmt{}
	} else if p.match(token.LeftBrace) {
		return p.blockStmt()
	} else {
		var success = false
		if p.Position+1 < len(p.Tokens) && p.Tokens[p.Position].Type == token.Identifier {
			t := p.Tokens[p.Position+1].Type
			if t == token.Equal ||
				t == token.PlusEqual ||
				t == token.MinusEqual ||
				t == token.AsteriskEqual ||
				t == token.SlashEqual ||
				t == token.PercentEqual ||
				t == token.AsteriskAsteriskEqual {
				success = true
				res, err = p.assignStmt(t)
			}
		}
		if !success {
			res, err = p.expressionStmt()
		}
	}

	if err != nil {
		return nil, err
	}

	err = p.consume(token.Semicolon, "Semicolon needed")
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Parser) varStmt() (ast.Statement, *errors.SyntaxError) {
	err := p.consume(token.Identifier, "Expect identifier after 'var'")
	if err != nil {
		return nil, err
	}

	name := p.previous().Literal

	err = p.consume(token.Equal, "Expect '=' after variable name")
	if err != nil {
		return nil, err
	}

	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	return ast.VarStmt{Name: name, Expr: expr}, nil
}

func (p *Parser) valStmt() (ast.Statement, *errors.SyntaxError) {
	err := p.consume(token.Identifier, "Expect identifier after 'val'")
	if err != nil {
		return nil, err
	}

	name := p.previous().Literal

	err = p.consume(token.Equal, "Expect '=' after variable name")
	if err != nil {
		return nil, err
	}

	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	return ast.ValStmt{Name: name, Expr: expr}, nil
}

func (p *Parser) assignStmt(t token.TokenType) (ast.Statement, *errors.SyntaxError) {
	name := p.Tokens[p.Position].Literal
	p.Position += 2

	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	return ast.AssignStmt{Name: name, Type: t, Expr: expr}, nil
}

func (p *Parser) ifStmt() (ast.Statement, *errors.SyntaxError) {
	err := p.consume(token.LeftParen, "Expect '(' after 'if'")
	if err != nil {
		return nil, err
	}

	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	err = p.consume(token.RightParen, "Expect ')' after if condition")
	if err != nil {
		return nil, err
	}

	then, err := p.statement()
	if err != nil {
		return nil, err
	}

	var elseBranch ast.Statement
	if p.match(token.Else) {
		elseBranch, err = p.statement()
		if err != nil {
			return nil, err
		}
	}

	return ast.IfStmt{Condition: condition, Then: then, Else: elseBranch}, nil
}

func (p *Parser) whileStmt() (ast.Statement, *errors.SyntaxError) {
	err := p.consume(token.LeftParen, "Expect '(' after 'while'")
	if err != nil {
		return nil, err
	}

	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	err = p.consume(token.RightParen, "Expect ')' after while condition")
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	return ast.WhileStmt{Condition: condition, Body: body}, nil
}

func (p *Parser) doWhileStmt() (ast.Statement, *errors.SyntaxError) {
	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	err = p.consume(token.While, "Expect 'while' after body of do")
	if err != nil {
		return nil, err
	}

	err = p.consume(token.LeftParen, "Expect '(' after 'while'")
	if err != nil {
		return nil, err
	}

	condition, err := p.expression()
	if err != nil {
		return nil, err
	}

	err = p.consume(token.RightParen, "Expect ')' after while condition")
	if err != nil {
		return nil, err
	}

	return ast.DoWhileStmt{Condition: condition, Body: body}, nil
}

func (p *Parser) funStmt() (ast.Statement, *errors.SyntaxError) {
	err := p.consume(token.Identifier, "Expect function name")
	if err != nil {
		return nil, err
	}
	name := p.previous().Literal

	err = p.consume(token.LeftParen, "Expect '(' after function name")
	if err != nil {
		return nil, err
	}

	var params []ast.Parameter
	if p.Position < len(p.Tokens) &&
		p.Tokens[p.Position].Type != token.RightParen {
		for {
			err = p.consume(token.Identifier, "Expect parameter name.")
			if err != nil {
				return nil, err
			}
			params = append(params, ast.Parameter{Name: p.previous().Literal})

			if !p.match(token.Comma) {
				break
			}
		}
	}

	err = p.consume(token.RightParen, "Expect ')' after parameters")
	if err != nil {
		return nil, err
	}

	body, err := p.statement()
	if err != nil {
		return nil, err
	}

	return ast.FunStmt{
		Name: name,
		Data: ast.ParamData{Params: params},
		Body: body,
	}, nil
}

func (p *Parser) returnStmt() (ast.Statement, *errors.SyntaxError) {
	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	return ast.ReturnStmt{Expr: expr}, nil
}

func (p *Parser) blockStmt() (ast.Statement, *errors.SyntaxError) {
	var stmts []ast.Statement

	for p.Position < len(p.Tokens) &&
		p.Tokens[p.Position].Type != token.RightBrace {

		stmt, err := p.statement()
		if err != nil {
			p.catch(err)
		} else {
			stmts = append(stmts, stmt)
		}
	}

	err := p.consume(token.RightBrace, "Expect '}' after block")
	if err != nil {
		return nil, err
	}

	return ast.BlockStmt{Body: stmts}, err
}

func (p *Parser) expressionStmt() (ast.Statement, *errors.SyntaxError) {
	expr, err := p.expression()

	if err != nil {
		return nil, err
	}
	return ast.ExpressionStmt{Expr: expr}, nil
}

func (p *Parser) expression() (ast.Expression, *errors.SyntaxError) {
	return p.disjunction()
}

func (p *Parser) disjunction() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.conjunction, token.PipePipe)
}

func (p *Parser) conjunction() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.comparison, token.AmpersandAmpersand)
}

func (p *Parser) comparison() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.bitwiseOr, token.EqualEqual, token.BangEqual,
		token.EqualEqualEqual, token.BangEqualEqual,
		token.LessEqual, token.Less,
		token.GreaterEqual, token.Greater)
}

func (p *Parser) bitwiseOr() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.bitwiseXor, token.Pipe)
}

func (p *Parser) bitwiseXor() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.bitwiseAnd, token.Circumflex)
}

func (p *Parser) bitwiseAnd() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.shift, token.Ampersand)
}

func (p *Parser) shift() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.sum, token.LeftShift, token.RightShift, token.URightShift)
}

func (p *Parser) sum() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.term, token.Plus, token.Minus)
}

func (p *Parser) term() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.unary, token.Asterisk, token.Slash, token.Percent)
}

func (p *Parser) unary() (ast.Expression, *errors.SyntaxError) {
	if p.match(token.Plus, token.Minus, token.Tilde, token.Bang) {
		op := p.previous().Type
		right, err := p.unary()
		if err != nil {
			return nil, err
		}

		return ast.UnaryExpression{Type: op, Right: right}, nil
	}

	return p.power()
}

func (p *Parser) power() (ast.Expression, *errors.SyntaxError) {
	return p.loopMatch(p.primary, token.AsteriskAsterisk)
}

func (p *Parser) primary() (ast.Expression, *errors.SyntaxError) {
	expr, err := p.literal()
	if err != nil {
		return nil, err
	}

	if p.match(token.LeftParen) {
		return p.call(expr)
	} else if p.match(token.LeftBracket) {
		return p.index(expr)
	}

	return expr, nil
}

func (p *Parser) call(expr ast.Expression) (ast.Expression, *errors.SyntaxError) {
	var args []ast.Expression

	if p.Position < len(p.Tokens) &&
		p.Tokens[p.Position].Type != token.RightParen {

		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		args = append(args, expr)

		for p.match(token.Comma) {
			expr, err := p.expression()
			if err != nil {
				return nil, err
			}
			args = append(args, expr)
		}
	}

	err := p.consume(token.RightParen, "Expect ')' after arguments")
	if err != nil {
		return nil, err
	}

	return ast.CallExpression{Target: expr, Arguments: args}, nil
}

func (p *Parser) index(expr ast.Expression) (ast.Expression, *errors.SyntaxError) {
	index, err := p.expression()
	if err != nil {
		return nil, err
	}

	err = p.consume(token.RightBracket, "Expect ']' after index")
	if err != nil {
		return nil, err
	}

	return ast.IndexExpression{Target: expr, Index: index}, nil
}

func (p *Parser) literal() (ast.Expression, *errors.SyntaxError) {
	if p.match(token.True) {
		return ast.LiteralExpression{Object: typing.Bool{Value: true}}, nil
	} else if p.match(token.False) {
		return ast.LiteralExpression{Object: typing.Bool{Value: false}}, nil
	} else if p.match(token.Null) {
		return ast.LiteralExpression{Object: typing.Null{}}, nil
	} else if p.match(token.Int) {
		num, success := new(big.Int).SetString(p.previous().Literal, 10)
		if !success {
			fmt.Println(p.previous())
			return nil, &errors.SyntaxError{Message: "Can't parse int"}
		}
		return ast.LiteralExpression{Object: typing.Int{Value: num}}, nil
	} else if p.match(token.Double) {
		num, _, err := new(big.Float).Parse(p.previous().Literal, 10)
		if err != nil {
			return nil, &errors.SyntaxError{Message: "Can't parse double"}
		}
		return ast.LiteralExpression{Object: typing.Double{Value: num}}, nil
	} else if p.match(token.String) {
		return ast.LiteralExpression{Object: typing.String{Value: p.previous().Literal}}, nil
	} else if p.match(token.LeftParen) {
		return p.tuple()
	} else if p.match(token.Char) {
		// TODO: Maybe use types from lexing directly
		c := []rune(p.previous().Literal)[0]

		return ast.LiteralExpression{Object: typing.Char{Value: c}}, nil
	} else if p.match(token.LeftBracket) {
		return p.array()
	} else if p.match(token.LeftBrace) {
		return p.hMap()
	} else if p.match(token.Identifier) {
		return ast.VariableExpression{Name: p.previous().Literal}, nil
	}

	return nil, errors.NewSyntaxError("Expect expression", p.Tokens[p.Position])
}

func (p *Parser) array() (ast.Expression, *errors.SyntaxError) {
	var values []ast.Expression

	for len(p.Tokens) > p.Position && p.Tokens[p.Position].Type != token.RightBracket {
		val, err := p.expression()
		if err != nil {
			return nil, err
		}

		values = append(values, val)

		if !p.match(token.Comma) {
			break
		}
	}

	p.consume(token.RightBracket, "Expect ']' after array")

	return ast.ArrayExpression{Values: values}, nil
}

func (p *Parser) hMap() (ast.Expression, *errors.SyntaxError) {
	items := make(map[ast.Expression]ast.Expression)

	for len(p.Tokens) > p.Position && p.Tokens[p.Position].Type != token.RightBrace {
		key, err := p.expression()
		if err != nil {
			return nil, err
		}

		if len(p.Tokens) > p.Position &&
			p.Tokens[p.Position].Type == token.Comma ||
			p.Tokens[p.Position].Type == token.RightBrace {

			if key, ok := key.(ast.VariableExpression); ok {
				sKey := ast.LiteralExpression{Object: typing.String{Value: key.Name}}
				items[sKey] = key
			} else {
				return nil, errors.NewSyntaxError("You can only use the key shorthand if you provide a variable", p.Tokens[p.Position])
			}
		} else {
			err := p.consume(token.Colon, "Expect ':' after key in map")
			if err != nil {
				return nil, err
			}

			value, err := p.expression()
			if err != nil {
				return nil, err
			}
			items[key] = value
		}

		if !p.match(token.Comma) {
			break
		}
	}

	p.consume(token.RightBrace, "Expect '}' after map")

	return ast.MapExpression{Items: items}, nil
}

func (p *Parser) tuple() (ast.Expression, *errors.SyntaxError) {
	var values []ast.Expression

	for len(p.Tokens) > p.Position && p.Tokens[p.Position].Type != token.RightParen {
		val, err := p.expression()
		if err != nil {
			return nil, err
		}

		values = append(values, val)

		if !p.match(token.Comma) {
			break
		}
	}

	if len(values) == 1 {
		err := p.consume(token.RightParen, "Expect ')' after expression")
		if err != nil {
			return nil, err
		}

		return ast.GroupingExpression{Expr: values[0]}, nil
	} else {
		err := p.consume(token.RightParen, "Expect ')' after tuple")
		if err != nil {
			return nil, err
		}

		return ast.TupleExpression{Values: values}, nil
	}
}

func (p *Parser) match(types ...token.TokenType) bool {
	if p.Position >= len(p.Tokens) {
		return false
	}

	for _, t := range types {
		if p.Tokens[p.Position].Type == t {
			p.Position++
			return true
		}
	}
	return false
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Position-1]
}

func (p *Parser) consume(t token.TokenType, msg string) *errors.SyntaxError {
	if p.Position < len(p.Tokens) && p.Tokens[p.Position].Type == t {
		p.Position++
		return nil
	}

	if p.Position < len(p.Tokens) {
		tok := p.Tokens[p.Position]
		return errors.NewSyntaxError(msg, tok)
	} else {
		tok := p.Tokens[p.Position-1]
		return &errors.SyntaxError{
			Message: msg,
			Line:    tok.Line,
			Column:  tok.Column + 1,
		}
	}
}

func (p *Parser) loopMatch(fn func() (ast.Expression, *errors.SyntaxError), types ...token.TokenType) (ast.Expression, *errors.SyntaxError) {
	expr, err := fn()
	if err != nil {
		return nil, err
	}

	for p.match(types...) {
		op := p.previous().Type

		right, err := fn()
		if err != nil {
			return nil, err
		}

		expr = ast.BinaryExpression{Left: expr, Type: op, Right: right}
	}

	return expr, nil
}

func (p *Parser) catch(err *errors.SyntaxError) {
	p.errs = append(p.errs, *err)
	for p.Position < len(p.Tokens) {
		if p.Tokens[p.Position].Type == token.Semicolon {
			break
		}
		p.Position++
	}
	p.Position++
}
