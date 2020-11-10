package typing

import "fmt"

var _ Object = Slice{}

type Slice struct {
	From Int
	To   Int
	Step Int
}

func (o Slice) ToString() string {
	return fmt.Sprintf("slice(%s, %s, %s)", o.From.ToString(), o.To.ToString(), o.Step.ToString())
}

func (o Slice) Class() Class {
	return Class{"slice"}
}
