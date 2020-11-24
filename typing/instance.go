package typing

import "fmt"

var _ Object = Instance{}

type Instance struct {
	class  Class
	fields map[string]Object
}

func (o Instance) ToString() string {
	return o.class.Name
}

func (o Instance) Class() Class {
	return o.class
}

func (o Instance) GetProperty(name string) (Object, Object) {
	if val, ok := o.fields[name]; ok {
		return val, nil
	}

	if val, ok := o.class.Properties[name]; ok {
		return val, nil
	}

	return nil, AttributeError{Name: name}
}

func (o Instance) SetProperty(name string, object Object) Object {
	if _, ok := o.class.Properties[name]; ok {
		o.fields[name] = object
		return nil
	}

	return Error{Message: fmt.Sprintf("Has no such field '%s'", name)}
}