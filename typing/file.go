package typing

type File struct {
	FPath string
	Path  string
	Props map[string]Object
}

func NewFile(fileName string, path string) *File {
	return &File{
		FPath: fileName,
		Path:  path,
		Props: make(map[string]Object),
	}
}
