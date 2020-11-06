package token

type TokenType int

const (
	Nothing = iota
	EOF

	Identifier

	LeftParen
	RightParen
	LeftBrace
	RightBrace
	LeftBracket
	RightBracket
	Comma
	Dot
	Semicolon

	Plus
	PlusPlus
	PlusEqual

	Minus
	MinusMinus
	MinusEqual

	Asterisk
	AsteriskEqual

	Slash
	SlashEqual

	AsteriskAsterisk
	AsteriskAsteriskEqual

	Percent
	PercentEqual

	Equal
	EqualEqual
	EqualEqualEqual

	Bang
	BangEqual
	BangEqualEqual

	Greater
	GreaterEqual
	Less
	LessEqual

	Ampersand
	AmpersandAmpersand
	Pipe
	PipePipe

	Circumflex
	Tilde
	LeftShift
	RightShift
	URightShift

	Null
	Bool
	Int
	Double
	String

	True
	False

	Var
	Val
	Fun
	If
	Else
	Do
	While
	For
	Return
	Break
	Continue
)
