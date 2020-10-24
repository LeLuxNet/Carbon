package typing

type Callable interface {
	Call(args []Object) Object
}
