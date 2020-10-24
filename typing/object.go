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

func Eq(a Object, b Object) bool {
	switch a := a.(type) {
	case Null:
		// Other objects can not implement eq to null
		_, ok := b.(Null)
		return ok
	case Bool:
		if b, ok := b.(Bool); ok {
			return a.Value == b.Value
		}
	case Class:
		if b, ok := b.(Class); ok {
			// TODO: Better comparison
			return a.Class().Name == b.Class().Name
		}
	case Double:
		if b, ok := b.(Double); ok {
			return a.Value == b.Value
		}
		if b, ok := b.(Int); ok {
			return a.Value == float64(b.Value)
		}
	case Int:
		if b, ok := b.(Int); ok {
			return a.Value == b.Value
		}
		if b, ok := b.(Double); ok {
			return float64(a.Value) == b.Value
		}
		if b, ok := b.(Bool); ok {
			if b.Value {
				return a.Value == 1
			} else {
				return a.Value == 0
			}
		}
	case String:
		if b, ok := b.(String); ok {
			return a.Value == b.Value
		}
	}
	return false
}
