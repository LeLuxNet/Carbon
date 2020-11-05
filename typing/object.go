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
	default:
		return true
	}
}
