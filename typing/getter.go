package typing

type Getter interface {
	Call(object Object, file *File) (Object, Throwable)
}

type BGetter struct {
	Name string
	Cal  func(this Object, file *File) (Object, Throwable)
}

func (o BGetter) Call(this Object, file *File) (Object, Throwable) {
	return o.Cal(this, file)
}

func (o BGetter) ToString() string {
	panic("This should not be called! A getter is not a type")
}

func (o BGetter) Class() Class {
	panic("This should not be called! A getter is not a type")
}
