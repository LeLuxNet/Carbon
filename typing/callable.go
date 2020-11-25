package typing

type Callable interface {
	Data() ParamData
	Call(this Object, params map[string]Object, args []Object, kwArgs map[string]Object, file *File) Throwable
}

type Parameter struct {
	Name    string
	Type    Class
	Default Object
}

type ParamData struct {
	Params     []Parameter
	Args       string
	KwArgs     string
	ReturnType Class
}
