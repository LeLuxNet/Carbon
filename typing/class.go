package typing

var _ Object = Class{}

type Class struct {
	Name string
}

func (o Class) ToString() string {
	return "class<" + o.Name + ">"
}

func (o Class) Class() Class {
	return Class{"class"}
}

func (o Class) IsInstance(object Object) bool {
	return Eq(object.Class(), o)
}
