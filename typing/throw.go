package typing

type Throwable interface {
	TData() Object
}

type Throw struct {
	Data Object
}
type Return struct {
	Data Object
}
type Break struct{}
type Continue struct{}

func (t Throw) TData() Object    { return t.Data }
func (t Return) TData() Object   { return t.Data }
func (t Break) TData() Object    { return Error{Message: "'break' outside of loop"} }
func (t Continue) TData() Object { return Error{Message: "'continue' outside of loop"} }

func NewError(msg string) Throw {
	return Throw{Data: Error{Message: msg}}
}
