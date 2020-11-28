package typing

var _ Object = Class{}

var toString = BFunction{Name: "toString", Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	return Return{String{this.ToString()}}
}}
var class = BFunction{Name: "Class", Cal: func(this Object, _ map[string]Object, _ []Object, _ map[string]Object, _ *File) Throwable {
	return Return{this.Class()}
}}

type Properties = map[string]Object
type Class struct {
	Name string
	Properties
	Static Properties
}

func NewNativeClass(name string, properties Properties) Class {
	// TODO: Set default functions
	properties["toString"] = toString
	properties["Class"] = class

	return Class{Name: name, Properties: properties, Static: Properties{}}
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

func (o Class) GetProperty(name string) (Object, Object) {
	if res, ok := o.Static[name]; ok {
		return res, nil
	}
	return nil, AttributeError{Name: name}
}

func (o Class) Allows(other Object) bool {
	return o.IsInstance(other)
}

func (o Class) TName() string {
	return o.Name
}
