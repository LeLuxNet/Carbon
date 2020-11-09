package typing

type Object interface {
	ToString() string
	Class() Class
}

func Truthy(object Object) bool {
	switch object := object.(type) {
	case Bool:
		return object.Value
	case Null:
		return false
	case Int:
	case Double:
		return object.Value.Sign() != 0
	case String:
		return len(object.Value) != 0
	}
	return true
}
