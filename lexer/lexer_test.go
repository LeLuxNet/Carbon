package lexer

import (
	"github.com/leluxnet/carbon/test"
	"github.com/leluxnet/carbon/token"
	"testing"
)

const Script = "var x = 12; if x > 10 { print(\"Hello\"); } // comment"

var Tokens = []token.Token{
	{Type: token.Var},
	{Type: token.Identifier, Literal: "x"},
	{Type: token.Equal},
	{Type: token.Int, Literal: "12"},
	{Type: token.Semicolon},
	{Type: token.If},
	{Type: token.Identifier, Literal: "x"},
	{Type: token.Greater},
	{Type: token.Int, Literal: "10"},
	{Type: token.LeftBrace},
	{Type: token.Identifier, Literal: "print"},
	{Type: token.LeftParen},
	{Type: token.String, Literal: "Hello"},
	{Type: token.RightParen},
	{Type: token.Semicolon},
	{Type: token.RightBrace},
}

func TestTokens(t *testing.T) {
	lexer := Lexer{Source: Script}
	tokens, errs := lexer.ScanTokens()
	if len(errs) != 0 {
		t.Fatal(errs)
	}

	test.AssertEq(t, len(Tokens), len(tokens))
	for i, tok := range tokens {
		etok := Tokens[i]

		test.AssertEq(t, etok.Type, tok.Type)
		test.AssertEq(t, etok.Literal, tok.Literal)
	}
}
