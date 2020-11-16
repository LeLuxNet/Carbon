package typing

var _ Object = Instance{}

type Instance struct {
	class Class
}

func (o Instance) ToString() string {
	return o.class.Name
}

func (o Instance) Class() Class {
	return o.class
}
