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

type Negable interface {
	Neg() Object
}
