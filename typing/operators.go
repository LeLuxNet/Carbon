package typing

type Addable interface {
	Add(Object, bool) Object
}

type Subable interface {
	Sub(Object, bool) Object
}

type Multable interface {
	Mult(Object, bool) Object
}

type Divable interface {
	Div(Object, bool) Object
}

type Powable interface {
	Pow(Object, bool) Object
}

type Modable interface {
	Mod(Object, bool) (Object, Object)
}

type Negable interface {
	Neg() Object
}
