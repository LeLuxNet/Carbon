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

func Eq(a, b Object) bool {
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

func Add(a, b Object) Object {
	if a, ok := a.(Addable); ok {
		res := a.Add(b, true)
		if res != nil {
			return res
		}
	}
	if b, ok := b.(Addable); ok {
		res := b.Add(a, false)
		return res
	}
	return nil
}

func Sub(a, b Object) Object {
	if a, ok := a.(Subable); ok {
		res := a.Sub(b, true)
		if res != nil {
			return res
		}
	}
	if b, ok := b.(Subable); ok {
		res := b.Sub(a, false)
		return res
	}
	return nil
}

func Mult(a, b Object) Object {
	if a, ok := a.(Multable); ok {
		res := a.Mult(b, true)
		if res != nil {
			return res
		}
	}
	if b, ok := b.(Multable); ok {
		res := b.Mult(a, false)
		return res
	}
	return nil
}

func Div(a, b Object) Object {
	if a, ok := a.(Divable); ok {
		res := a.Div(b, true)
		if res != nil {
			return res
		}
	}
	if b, ok := b.(Divable); ok {
		res := b.Div(a, false)
		return res
	}
	return nil
}

func Pow(a, b Object) Object {
	if a, ok := a.(Powable); ok {
		res := a.Pow(b, true)
		if res != nil {
			return res
		}
	}
	if b, ok := b.(Powable); ok {
		res := b.Pow(a, false)
		return res
	}
	return nil
}
