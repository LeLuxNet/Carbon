package token

var Keywords = map[string]TokenType{
	"fun":   Fun,
	"var":   Var,
	"val":   Val,
	"if":    If,
	"else":  Else,
	"while": While,
	"do":    Do,
	"for":   For,

	"return":   Return,
	"break":    Break,
	"continue": Continue,

	"null":  Null,
	"true":  True,
	"false": False,
}
