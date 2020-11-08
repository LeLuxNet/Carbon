package lexer

import (
	"github.com/leluxnet/carbon/errors"
	"github.com/leluxnet/carbon/token"
)

type Lexer struct {
	Position int
	Source   string
	Chars    []rune

	Line   int
	Column int
}

func (l *Lexer) ScanTokens() ([]token.Token, []errors.SyntaxError) {
	l.Chars = []rune(l.Source)

	var tokens []token.Token
	var errs []errors.SyntaxError

	var tok *token.Token
	var semi bool
	var err *errors.SyntaxError

	for l.Position < len(l.Chars) {
		tok, semi, err = l.scanToken(semi)
		if err != nil {
			errs = append(errs, *err)
		} else if tok.Type != token.Nothing {
			tokens = append(tokens, *tok)
		}
	}

	if semi {
		tokens = append(tokens, token.Token{
			Type:     token.Semicolon,
			Line:     l.Line,
			Column:   l.Column,
			ToLine:   l.Line,
			ToColumn: l.Column,
		})
	}

	return tokens, errs
}

func (l *Lexer) scanToken(lastSemi bool) (*token.Token, bool, *errors.SyntaxError) {
	var tok token.TokenType
	var semi bool

	fromLine := l.Line
	fromCol := l.Column

	c := l.Chars[l.Position]
	l.Position++
	l.Column++

	switch c {
	case '(':
		tok = token.LeftParen
	case ')':
		tok = token.RightParen
		semi = true
	case '{':
		tok = token.LeftBrace
	case '}':
		tok = token.RightBrace
	case '[':
		tok = token.LeftBracket
	case ']':
		tok = token.RightBracket
		semi = true
	case ',':
		tok = token.Comma
	case '.':
		if l.isEnd() || isDigit(l.Chars[l.Position]) {
			num := l.number()
			return &num, true, nil
		} else {
			tok = token.Dot
		}
	case ';':
		tok = token.Semicolon
	case '+':
		if l.isNextChar('+') {
			tok = token.PlusPlus
		} else if l.isNextChar('=') {
			tok = token.PlusEqual
		} else {
			tok = token.Plus
		}
	case '-':
		if l.isNextChar('-') {
			tok = token.MinusMinus
		} else if l.isNextChar('=') {
			tok = token.MinusEqual
		} else {
			tok = token.Minus
		}
	case '*':
		if l.isNextChar('*') {
			if l.isNextChar('=') {
				tok = token.AsteriskAsteriskEqual
			} else {
				tok = token.AsteriskAsterisk
			}
		} else if l.isNextChar('=') {
			tok = token.AsteriskEqual
		} else {
			tok = token.Asterisk
		}
	case '/':
		if l.isNextChar('/') {
			l.waitForChar('\n')
		} else if l.isNextChar('=') {
			tok = token.SlashEqual
		} else {
			tok = token.Slash
		}
	case '%':
		if l.isNextChar('=') {
			tok = token.PercentEqual
		} else {
			tok = token.Percent
		}
	case '!':
		if l.isNextChar('=') {
			if l.isNextChar('=') {
				tok = token.BangEqualEqual
			} else {
				tok = token.BangEqual
			}
		} else {
			tok = token.Bang
		}

	case '=':
		if l.isNextChar('=') {
			if l.isNextChar('=') {
				tok = token.EqualEqualEqual
			} else {
				tok = token.EqualEqual
			}
		} else {
			tok = token.Equal
		}

	case '<':
		if l.isNextChar('=') {
			tok = token.LessEqual
		} else if l.isNextChar('<') {
			tok = token.LeftShift
		} else {
			tok = token.Less
		}
	case '>':
		if l.isNextChar('=') {
			tok = token.GreaterEqual
		} else if l.isNextChar('>') {
			if l.isNextChar('>') {
				tok = token.URightShift
			} else {
				tok = token.RightShift
			}
		} else {
			tok = token.Greater
		}

	case '&':
		if l.isNextChar('&') {
			tok = token.AmpersandAmpersand
		} else {
			tok = token.Ampersand
		}
	case '^':
		tok = token.Circumflex
	case '~':
		tok = token.Tilde
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		l.Line++
		l.Column = 0
		if lastSemi {
			tok = token.Semicolon
		}
	case '"':
		tok, err := l.string()
		return tok, true, err
	default:
		if isDigit(c) {
			num := l.number()
			return &num, true, nil
		} else if isAlpha(c) {
			id, semi := l.identifier()
			return &id, semi, nil
		} else {
			return nil, false, &errors.SyntaxError{Message: "Unexpected char", Line: l.Line, Column: l.Column - 1}
		}
	}
	return &token.Token{Type: tok, Line: fromLine, Column: fromCol, ToLine: l.Line, ToColumn: l.Column}, semi, nil
}

func (l *Lexer) isNextChar(char rune) bool {
	if l.isEnd() || l.Chars[l.Position] != char {
		return false
	}

	l.Position++
	l.Column++
	return true
}

func (l *Lexer) waitForChar(char rune) {
	for !l.isEnd() && l.Chars[l.Position] != char {
		l.Position++
		l.Column++
	}
}

func (l *Lexer) isEnd() bool {
	return l.Position >= len(l.Chars)
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (l *Lexer) string() (*token.Token, *errors.SyntaxError) {
	pos := l.Position

	for !l.isEnd() && l.Chars[l.Position] != '"' && l.Chars[l.Position] != '\n' {
		l.Position++
		l.Column++
	}

	if l.isEnd() {
		return nil, &errors.SyntaxError{Message: "File ended with open string", Line: l.Line, Column: l.Column}
	} else if l.Chars[l.Position] == '\n' {
		return nil, &errors.SyntaxError{Message: "Line ended with open string", Line: l.Line, Column: l.Column}
	}

	l.Position++
	l.Column++

	return &token.Token{Type: token.String, Literal: string(l.Chars[pos : l.Position-1]),
		Line: l.Line, Column: l.Column}, nil
}

func (l *Lexer) number() token.Token {
	l.Position--
	l.Column--
	col := l.Column
	pos := l.Position

	for !l.isEnd() && isDigit(l.Chars[l.Position]) {
		l.Position++
		l.Column++
	}

	if !l.isEnd() && l.Chars[l.Position] == '.' {
		l.Position++
		l.Column++

		for !l.isEnd() && isDigit(l.Chars[l.Position]) {
			l.Position++
			l.Column++
		}

		return token.Token{Type: token.Double, Literal: string(l.Chars[pos:l.Position]), Line: l.Line, Column: col, ToLine: l.Line, ToColumn: l.Column}
	}

	return token.Token{Type: token.Int, Literal: string(l.Chars[pos:l.Position]), Line: l.Line, Column: col, ToLine: l.Line, ToColumn: l.Column}
}

func (l *Lexer) identifier() (token.Token, bool) {
	pos := l.Position - 1
	col := l.Column - 1

	for !l.isEnd() && isAlpha(l.Chars[l.Position]) {
		l.Position++
		l.Column++
	}

	text := string(l.Chars[pos:l.Position])
	if dat, ok := token.Keywords[text]; ok {
		return token.Token{Type: dat.TokenType, Line: l.Line, Column: col, ToLine: l.Line, ToColumn: l.Column}, dat.Semi
	}
	return token.Token{Type: token.Identifier, Literal: text, Line: l.Line, Column: col, ToLine: l.Line, ToColumn: l.Column}, true
}
