package token

type Token struct {
	Type    TokenType
	Literal string

	Line   int
	Column int

	ToLine   int
	ToColumn int
}
