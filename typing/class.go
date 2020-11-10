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

func (o Class) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(Class); ok {
		// TODO: Better comparison
		return Bool{o.Class().Name == other.Class().Name}, nil
	}
	return nil, nil
}

func (o Class) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(Class); ok {
		// TODO: Better comparison
		return Bool{o.Class().Name != other.Class().Name}, nil
	}
	return nil, nil
}

func (o Class) IsInstance(object Object) (Object, Throwable) {
	return Eq(object.Class(), o)
}
