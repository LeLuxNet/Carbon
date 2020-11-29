package typing

type Getter interface {
	Get(this Object, file *File) (Object, Throwable)
}

type BGetter struct {
	Name string
	Cal  func(this Object, file *File) (Object, Throwable)
}

func (o BGetter) Get(this Object, file *File) (Object, Throwable) {
	return o.Cal(this, file)
}

func (o BGetter) ToString() string {
	panic("This should not be called! A getter is not a type")
}

func (o BGetter) Class() Class {
	panic("This should not be called! A getter is not a type")
}

var _ Object = GSetter{}

type GSetter struct {
	Getter
	Setter
}

func (o GSetter) ToString() string {
	panic("Should not be called!")
}

func (o GSetter) Class() Class {
	panic("Should not be called!")
}
