package typing

var TypeClass = NewNativeClass("type", Properties{})

type Type interface {
	Allows(t Object) bool
	TName() string
}
