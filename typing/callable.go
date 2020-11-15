package typing

type Callable interface {
	Data() ParamData
	Call(this Object, args []Object) Throwable
}

type Parameter struct {
	Name    string
	Type    Class
	Default Object
}

type ParamData struct {
	Params []Parameter
	Args   string
	KwArgs string
}
