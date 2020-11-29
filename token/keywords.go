package token

var Keywords = map[string]struct {
	TokenType
	Semi bool
}{
	"var":   {Var, false},
	"val":   {Val, false},
	"if":    {If, false},
	"else":  {Else, false},
	"while": {While, false},
	"do":    {Do, false},
	"for":   {For, false},
	"class": {Class, false},
	"fun":   {Fun, false},
	"get":   {Get, false},
	"set":   {Set, false},
	"con":   {Con, false},
	"new":   {New, false},

	"return":   {Return, true},
	"break":    {Break, true},
	"continue": {Continue, true},
	"export":   {Export, true},

	"null":  {Null, true},
	"true":  {True, true},
	"false": {False, true},
}
