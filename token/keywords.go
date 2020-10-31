package token

var Keywords = map[string]struct {
	TokenType
	Semi bool
}{
	"fun":   {Fun, false},
	"var":   {Var, false},
	"val":   {Val, false},
	"if":    {If, false},
	"else":  {Else, false},
	"while": {While, false},
	"do":    {Do, false},
	"for":   {For, false},

	"return":   {Return, true},
	"break":    {Break, true},
	"continue": {Continue, true},

	"null":  {Null, true},
	"true":  {True, true},
	"false": {False, true},
}
