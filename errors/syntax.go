package errors

import (
	"fmt"
	"github.com/leluxnet/carbon/token"
	"strings"
)

type SyntaxError struct {
	Message string

	Line   int
	Column int

	LineLen   int
	ColumnLen int
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("error: %s (%d:%d)\n", s.Message, s.Line+1, s.Column+1)
}

func (s SyntaxError) ToString(source string) string {
	line := strings.Split(source, "\n")[s.Line]

	colLen := s.ColumnLen
	if colLen == 0 {
		colLen = 1
	}

	return fmt.Sprintf("error: %s (%d:%d)\n%s\n%s%s\n",
		s.Message, s.Line+1, s.Column+1,
		line,
		strings.Repeat(" ", s.Column), strings.Repeat("^", colLen))
}

func NewSyntaxError(msg string, token token.Token) *SyntaxError {
	fmt.Println(token)
	return &SyntaxError{Message: msg, Line: token.Line, Column: token.Column,
		LineLen: token.ToLine - token.Line, ColumnLen: token.ToColumn - token.Column}
}
