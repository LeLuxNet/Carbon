package typing

import "fmt"

var _ Object = Instance{}

type Instance struct {
	class  Class
	fields map[string]Object
	file   *File
}

func (o Instance) ToString() string {
	if fun := o.getProp("toString"); fun != nil {
		res := fun.(Callable).Call(o, map[string]Object{}, []Object{}, map[string]Object{}, o.file)
		if res, ok := res.(Return); ok {
			return res.Data.(String).Value
		} else {
			panic(res.TData().ToString())
		}
	}
	return o.class.Name
}

func (o Instance) Class() Class {
	return o.class
}

func (o Instance) getProp(name string) Object {
	if val, ok := o.fields[name]; ok {
		return val
	}

	if val, ok := o.class.Properties[name]; ok {
		return val
	}

	return nil
}

func (o Instance) GetProperty(name string) (Object, Object) {
	res := o.getProp(name)

	if res == nil {
		return nil, AttributeError{Name: name}
	} else {
		return res, nil
	}
}

func (o Instance) SetProperty(name string, object Object) Object {
	if _, ok := o.class.Properties[name]; ok {
		o.fields[name] = object
		return nil
	}

	return Error{Message: fmt.Sprintf("Has no such field '%s'", name)}
}
