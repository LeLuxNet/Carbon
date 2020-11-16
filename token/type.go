package token

type TokenType int

const (
	Nothing = iota

	Identifier

	LeftParen
	RightParen

	LeftBrace
	LeftMBrace
	LeftSBrace

	RightBrace
	LeftBracket
	RightBracket
	Comma
	Dot
	Colon
	Semicolon

	Arrow

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
	Char

	True
	False

	Var
	Val
	If
	Else
	While
	Do
	For
	Class
	Fun

	Return
	Break
	Continue
	Import
	From
	Export
)
