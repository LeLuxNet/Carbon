package typing

type Setter interface {
	Set(this, val Object, file *File) Throwable
}

type BSetter struct {
	Name string
	Cal  func(this, val Object, file *File) Throwable
}

func (o BSetter) Set(this, val Object, file *File) Throwable {
	return o.Cal(this, val, file)
}

func (o BSetter) ToString() string {
	panic("This should not be called! A setter is not a type")
}

func (o BSetter) Class() Class {
	panic("This should not be called! A setter is not a type")
}
