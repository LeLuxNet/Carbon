package typing

var _ Object = Module{}

type Module struct {
	Name string
	Properties
}

func (o Module) ToString() string {
	return o.Name
}

func (o Module) Class() Class {
	return NewNativeClass(o.Name, o.Properties)
}
