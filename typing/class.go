package typing

var _ Object = Class{}

var toString = BFunction{Name: "toString", Cal: func(this Object, _ []Object) Throwable {
	return Return{String{this.ToString()}}
}}
var class = BFunction{Name: "Class", Cal: func(this Object, _ []Object) Throwable {
	return Return{this.Class()}
}}

type Properties = map[string]Object
type Class struct {
	Name string
	Properties
}

func NewNativeClass(name string, properties Properties) Class {
	// TODO: Set default functions
	properties["toString"] = toString
	properties["Class"] = class

	return Class{Name: name, Properties: properties}
}

func (o Class) ToString() string {
	return "class<" + o.Name + ">"
}

func (o Class) Class() Class {
	return NewNativeClass("class", Properties{})
}

func (o Class) Eq(other Object) (Object, Throwable) {
	if other, ok := other.(Class); ok {
		// TODO: Better comparison
		return Bool{o.Name == other.Class().Name}, nil
	}
	return nil, nil
}

func (o Class) NEq(other Object) (Object, Throwable) {
	if other, ok := other.(Class); ok {
		// TODO: Better comparison
		return Bool{o.Name != other.Class().Name}, nil
	}
	return nil, nil
}

func (o Class) IsInstance(other Object) bool {
	// TODO: Better comparison
	return o.Name == other.Class().Name
}
