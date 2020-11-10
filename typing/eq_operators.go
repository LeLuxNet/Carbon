package typing

type Eqable interface {
	Eq(Object) (Object, Throwable)
}

func Eq(a, b Object) (Object, Throwable) {
	if a, ok := a.(Eqable); ok {
		res, err := a.Eq(b)
		if err != nil || res != nil {
			return res, err
		}
	}
	if b, ok := b.(Eqable); ok {
		return b.Eq(a)
	}
	return Bool{false}, nil
}

type NEqable interface {
	NEq(Object) (Object, Throwable)
}

func NEq(a, b Object) (Object, Throwable) {
	if a, ok := a.(NEqable); ok {
		res, err := a.NEq(b)
		if err != nil || res != nil {
			return res, err
		}
	}
	if a, ok := a.(Eqable); ok {
		res, err := a.Eq(b)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return Bool{!Truthy(res)}, nil
		}
	}
	if b, ok := b.(NEqable); ok {
		res, err := b.NEq(a)

		if err != nil || res != nil {
			return res, err
		}
	}
	if b, ok := b.(Eqable); ok {
		res, err := b.Eq(a)
		if err != nil {
			return nil, err
		}
		if res != nil {
			return Bool{!Truthy(res)}, nil
		}
	}
	return Bool{true}, nil
}
