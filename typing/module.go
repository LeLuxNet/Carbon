package typing

import (
	"fmt"
)

var _ Object = Module{}
var _ PropertyGettable = Module{}

type Module struct {
	Name  string
	Items map[string]Object
}

func (o Module) ToString() string {
	return fmt.Sprintf("module<%s>", o.Name)
}

func (o Module) Class() Class {
	return NewNativeClass("module", Properties{})
}

func (o Module) GetProperty(name string) (Object, Object) {
	if res, ok := o.Items[name]; ok {
		return res, nil
	} else {
		return nil, AttributeError{Name: name}
	}
}
